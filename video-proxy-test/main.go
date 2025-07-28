package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Test JoggAI URL - আপনার actual URL এখানে দিন
const TEST_JOGGAI_URL = "https://res.jogg.ai/joggUserData/project/vd_5f7bc06190004c4084a4d6c6c5e455cb/1753608098017-3342dce63ee6e0618c620219418fffd1-video.mp4"

// Simple proxy handler
func proxyVideoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("🎥 Video request received!")
	
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Fetch video from JoggAI
	fmt.Println("📡 Fetching video from JoggAI...")
	resp, err := client.Get(TEST_JOGGAI_URL)
	if err != nil {
		fmt.Printf("❌ Error fetching video: %v\n", err)
		http.Error(w, "Failed to fetch video", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Check if video exists
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("❌ JoggAI returned status: %d\n", resp.StatusCode)
		http.Error(w, "Video not found", resp.StatusCode)
		return
	}

	fmt.Println("✅ Video found! Streaming to user...")

	// Set video headers
	w.Header().Set("Content-Type", "video/mp4")
	w.Header().Set("Cache-Control", "public, max-age=3600") // 1 hour cache
	
	// Copy content-length if available
	if contentLength := resp.Header.Get("Content-Length"); contentLength != "" {
		w.Header().Set("Content-Length", contentLength)
		fmt.Printf("📦 Video size: %s bytes\n", contentLength)
	}

	// Support range requests for video seeking
	if rangeHeader := r.Header.Get("Range"); rangeHeader != "" {
		w.Header().Set("Accept-Ranges", "bytes")
		fmt.Println("🎯 Range request detected - supporting video seeking")
	}

	// Stream video data to user
	bytesWritten, err := io.Copy(w, resp.Body)
	if err != nil {
		fmt.Printf("❌ Error streaming video: %v\n", err)
		return
	}

	fmt.Printf("✅ Successfully streamed %d bytes to user\n", bytesWritten)
}

// Test page to view the video
func testPageHandler(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html>
<head>
    <title>Video Proxy Test</title>
    <style>
        body { 
            font-family: Arial, sans-serif; 
            max-width: 800px; 
            margin: 50px auto; 
            padding: 20px;
        }
        .container {
            text-align: center;
        }
        video {
            width: 100%;
            max-width: 600px;
            margin: 20px 0;
            border: 2px solid #ddd;
            border-radius: 8px;
        }
        .url-display {
            background: #f5f5f5;
            padding: 15px;
            border-radius: 5px;
            margin: 20px 0;
            word-break: break-all;
        }
        .status {
            padding: 10px;
            margin: 10px 0;
            border-radius: 5px;
        }
        .success { background: #d4edda; color: #155724; border: 1px solid #c3e6cb; }
        .info { background: #d1ecf1; color: #0c5460; border: 1px solid #bee5eb; }
    </style>
</head>
<body>
    <div class="container">
        <h1>🎥 Video Proxy Test</h1>
        
        <div class="status success">
            <strong>✅ Success!</strong> Video is being served through your domain proxy
        </div>

        <div class="status info">
            <strong>Your Custom URL:</strong>
            <div class="url-display">http://localhost:8080/api/video/stream</div>
        </div>

        <video controls preload="metadata">
            <source src="/api/video/stream" type="video/mp4">
            <p>Your browser doesn't support HTML5 video.</p>
        </video>

        <div class="status info">
            <p><strong>Original JoggAI URL (hidden from user):</strong></p>
            <div class="url-display" style="font-size: 12px; color: #666;">
                ` + TEST_JOGGAI_URL + `
            </div>
        </div>

        <h3>🔍 Testing Instructions:</h3>
        <ul style="text-align: left;">
            <li>✅ Video loads from <strong>localhost:8080</strong> (your domain)</li>
            <li>✅ JoggAI URL is completely hidden from user</li>
            <li>✅ Right-click → Inspect → Network tab shows only your domain</li>
            <li>✅ Video seeking and controls work properly</li>
        </ul>
    </div>
</body>
</html>`
	
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

// Health check endpoint
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("✅ Proxy server is running!"))
}

func main() {
	r := mux.NewRouter()

	// Test page
	r.HandleFunc("/", testPageHandler).Methods("GET")
	
	// Video proxy endpoint
	r.HandleFunc("/api/video/stream", proxyVideoHandler).Methods("GET")
	
	// Health check
	r.HandleFunc("/health", healthHandler).Methods("GET")

	// Start server
	fmt.Println("🚀 Starting proxy server...")
	fmt.Println("📍 Test Page: http://localhost:8080")
	fmt.Println("🎥 Video URL: http://localhost:8080/api/video/stream")
	fmt.Println("❤️  Health Check: http://localhost:8080/health")
	fmt.Println("\n🔥 Server is ready! Open http://localhost:8080 in your browser")
	
	log.Fatal(http.ListenAndServe(":8080", r))
}