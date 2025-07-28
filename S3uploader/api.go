package main

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"mime/multipart"
// 	"net/http"
// 	"path/filepath"
// 	"strings"

// 	"github.com/google/uuid"
// 	"github.com/minio/minio-go/v7"
// )

// // UploadResponse represents the response for image uploads
// // type UploadResponse struct {
// // 	Success bool       `json:"success"`
// // 	Message string     `json:"message"`
// // 	Files   []FileInfo `json:"files,omitempty"`
// // 	Error   string     `json:"error,omitempty"`
// // }

// // FileInfo represents information about an uploaded file
// type FileInfo struct {
// 	OriginalName string `json:"original_name"`
// 	FileName     string `json:"file_name"`
// 	FilePath     string `json:"file_path"`
// 	FileURL      string `json:"file_url,omitempty"`
// 	Size         int64  `json:"size"`
// }

// // handleUploadImages handles multiple image uploads
// func (m *MinIOUploader) handleUploadImages(w http.ResponseWriter, r *http.Request) {
// 	// Set CORS headers
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

// 	// Handle preflight requests
// 	if r.Method == "OPTIONS" {
// 		w.WriteHeader(http.StatusOK)
// 		return
// 	}

// 	// Only allow POST method
// 	if r.Method != "POST" {
// 		respondWithError(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	// Parse multipart form (max 32MB)
// 	err := r.ParseMultipartForm(32 << 20)
// 	if err != nil {
// 		respondWithError(w, "Failed to parse form: "+err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	// Get uploaded files
// 	files := r.MultipartForm.File["images"]
// 	if len(files) == 0 {
// 		respondWithError(w, "No images provided. Use 'images' as the form field name.", http.StatusBadRequest)
// 		return
// 	}

// 	// Validate file count (max 10 images)
// 	if len(files) > 10 {
// 		respondWithError(w, "Maximum 10 images allowed per request", http.StatusBadRequest)
// 		return
// 	}

// 	var uploadedFiles []FileInfo
// 	ctx := context.Background()

// 	// Process each uploaded file
// 	for _, fileHeader := range files {
// 		fileInfo, err := m.processUploadedFile(ctx, fileHeader)
// 		if err != nil {
// 			log.Printf("Failed to process file %s: %v", fileHeader.Filename, err)
// 			continue
// 		}
// 		uploadedFiles = append(uploadedFiles, fileInfo)
// 	}

// 	// Check if any files were successfully uploaded
// 	if len(uploadedFiles) == 0 {
// 		respondWithError(w, "No files were successfully uploaded", http.StatusInternalServerError)
// 		return
// 	}

// 	// Return success response
// 	response := UploadResponse{
// 		Success: true,
// 		Message: fmt.Sprintf("Successfully uploaded %d image(s)", len(uploadedFiles)),
// 		Files:   uploadedFiles,
// 	}

// 	respondWithJSON(w, response, http.StatusOK)
// }

// // processUploadedFile processes a single uploaded file
// func (m *MinIOUploader) processUploadedFile(ctx context.Context, fileHeader *multipart.FileHeader) (FileInfo, error) {
// 	// Validate file type
// 	if !isValidImageFile(fileHeader.Filename) {
// 		return FileInfo{}, fmt.Errorf("invalid image file type: %s", fileHeader.Filename)
// 	}

// 	// Open the uploaded file
// 	file, err := fileHeader.Open()
// 	if err != nil {
// 		return FileInfo{}, fmt.Errorf("failed to open uploaded file: %v", err)
// 	}
// 	defer file.Close()

// 	// Generate UUID filename
// 	fileExt := filepath.Ext(fileHeader.Filename)
// 	objectName := fmt.Sprintf("%s%s", uuid.New().String(), fileExt)

// 	// Upload to MinIO
// 	uploadInfo, err := m.client.PutObject(ctx, m.config.BucketName, objectName, file, fileHeader.Size, minio.PutObjectOptions{
// 		ContentType: getContentType(fileExt),
// 	})
// 	if err != nil {
// 		return FileInfo{}, fmt.Errorf("failed to upload to MinIO: %v", err)
// 	}

// 	// Generate file URL
// 	fileURL, err := m.GetFileURL(ctx, m.config.BucketName, objectName)
// 	if err != nil {
// 		log.Printf("Warning: Failed to generate URL for %s: %v", objectName, err)
// 		fileURL = ""
// 	}

// 	fileInfo := FileInfo{
// 		OriginalName: fileHeader.Filename,
// 		FileName:     objectName,
// 		FilePath:     fmt.Sprintf("%s/%s", m.config.BucketName, objectName),
// 		FileURL:      fileURL,
// 		Size:         uploadInfo.Size,
// 	}

// 	log.Printf("Successfully uploaded: %s -> %s (Size: %d bytes)",
// 		fileHeader.Filename, objectName, uploadInfo.Size)

// 	return fileInfo, nil
// }

// // isValidImageFile checks if the file is a valid image
// func isValidImageFile(filename string) bool {
// 	ext := strings.ToLower(filepath.Ext(filename))
// 	validExtensions := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp"}

// 	for _, validExt := range validExtensions {
// 		if ext == validExt {
// 			return true
// 		}
// 	}
// 	return false
// }

// // getContentType returns the MIME type based on file extension
// func getContentType(ext string) string {
// 	ext = strings.ToLower(ext)
// 	switch ext {
// 	case ".jpg", ".jpeg":
// 		return "image/jpeg"
// 	case ".png":
// 		return "image/png"
// 	case ".gif":
// 		return "image/gif"
// 	case ".bmp":
// 		return "image/bmp"
// 	case ".webp":
// 		return "image/webp"
// 	default:
// 		return "application/octet-stream"
// 	}
// }

// // respondWithJSON sends a JSON response
// func respondWithJSON(w http.ResponseWriter, data interface{}, statusCode int) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(statusCode)
// 	json.NewEncoder(w).Encode(data)
// }

// // respondWithError sends an error response
// func respondWithError(w http.ResponseWriter, message string, statusCode int) {
// 	response := UploadResponse{
// 		Success: false,
// 		Error:   message,
// 	}
// 	respondWithJSON(w, response, statusCode)
// }

// // handleHealthCheck provides a health check endpoint
// func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
// }
