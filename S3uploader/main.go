package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// MinIOUploader handles MinIO operations
type MinIOUploader struct {
	client *minio.Client
}

// Config holds MinIO configuration
type Config struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
	BucketName      string
	ImageFile       string
}

// NewMinIOUploader creates a new MinIO uploader instance
func NewMinIOUploader(config Config) (*MinIOUploader, error) {
	// Initialize MinIO client
	minioClient, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKey, ""),
		Secure: config.UseSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize MinIO client: %v", err)
	}

	log.Printf("MinIO client initialized for endpoint: %s", config.Endpoint)
	return &MinIOUploader{client: minioClient}, nil
}

// EnsureBucketExists creates bucket if it doesn't exist
func (m *MinIOUploader) EnsureBucketExists(ctx context.Context, bucketName string) error {
	// Check if bucket exists
	exists, err := m.client.BucketExists(ctx, bucketName)
	if err != nil {
		return fmt.Errorf("failed to check if bucket exists: %v", err)
	}

	if !exists {
		// Create bucket
		err = m.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return fmt.Errorf("failed to create bucket: %v", err)
		}
		log.Printf("Bucket '%s' created successfully", bucketName)
	} else {
		log.Printf("Bucket '%s' already exists", bucketName)
	}

	return nil
}

// UploadImage uploads an image file to MinIO with UUID filename
func (m *MinIOUploader) UploadImage(ctx context.Context, localFilePath, bucketName string) (string, error) {
	// Check if file exists
	if _, err := os.Stat(localFilePath); os.IsNotExist(err) {
		return "", fmt.Errorf("file '%s' not found", localFilePath)
	}

	// Generate UUID filename
	fileExt := filepath.Ext(localFilePath)
	objectName := fmt.Sprintf("%s%s", uuid.New().String(), fileExt)

	// Upload file
	uploadInfo, err := m.client.FPutObject(ctx, bucketName, objectName, localFilePath, minio.PutObjectOptions{
		ContentType: "image/jpeg",
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %v", err)
	}

	log.Printf("File '%s' uploaded as '%s' to bucket '%s' (Size: %d bytes)",
		localFilePath, objectName, bucketName, uploadInfo.Size)

	return objectName, nil
}

// GetFileURL generates a presigned URL for the uploaded file
func (m *MinIOUploader) GetFileURL(ctx context.Context, bucketName, objectName string) (string, error) {
	// Generate presigned URL (valid for 7 days)
	reqParams := make(map[string][]string)
	url, err := m.client.PresignedGetObject(ctx, bucketName, objectName,
		7*24*time.Hour, // 7 days
		reqParams)
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL: %v", err)
	}

	return url.String(), nil
}

// FileExists checks if a file exists
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func main() {
	// Configuration
	config := Config{
		Endpoint:        "localhost:9000",
		AccessKeyID:     "minioadmin",
		SecretAccessKey: "minioadmin123",
		UseSSL:          false,
		BucketName:      "images",
		ImageFile:       "image.jpg",
	}

	// Check if image file exists
	if !FileExists(config.ImageFile) {
		log.Fatalf("Image file '%s' not found in current directory", config.ImageFile)
	}

	// Create context
	ctx := context.Background()

	// Initialize uploader
	uploader, err := NewMinIOUploader(config)
	if err != nil {
		log.Fatalf("Failed to initialize uploader: %v", err)
	}

	// Ensure bucket exists
	err = uploader.EnsureBucketExists(ctx, config.BucketName)
	if err != nil {
		log.Fatalf("Failed to ensure bucket exists: %v", err)
	}

	// Upload image
	log.Printf("Uploading '%s' with UUID filename...", config.ImageFile)
	objectName, err := uploader.UploadImage(ctx, config.ImageFile, config.BucketName)
	if err != nil {
		log.Fatalf("Upload failed: %v", err)
	}

	// Get file URL
	fileURL, err := uploader.GetFileURL(ctx, config.BucketName, objectName)
	if err != nil {
		log.Printf("Warning: Failed to generate file URL: %v", err)
		fileURL = "URL generation failed"
	}

	// Success message
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("UPLOAD SUCCESSFUL!")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("Original file: %s\n", config.ImageFile)
	fmt.Printf("Uploaded as: %s\n", objectName)
	fmt.Printf("Bucket: %s\n", config.BucketName)
	fmt.Printf("File URL: %s\n", fileURL)
	fmt.Println(strings.Repeat("=", 60))
}
