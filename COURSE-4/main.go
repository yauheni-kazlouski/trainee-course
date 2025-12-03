package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

// Application error definitions
var (
	ErrInvalidURL       = errors.New("invalid URL")
	ErrConnectionFailed = errors.New("connection failed")
	ErrDownloadFailed   = errors.New("download failed")
	ErrFileNotFound     = errors.New("file not found")
)

// InvalidProtocolError - custom error for unsupported protocols
type InvalidProtocolError struct {
	Protocol string
}

func (e InvalidProtocolError) Error() string {
	return fmt.Sprintf("invalid protocol: %s", e.Protocol)
}

// downloadFile downloads a file from the specified URL and saves it locally
// Returns an error if any step of the process fails
func downloadFile(fileURL, fileName string) error {
	// Validate and parse the URL
	parsedURL, err := url.ParseRequestURI(fileURL)
	if err != nil {
		return ErrInvalidURL
	}

	// Check if the protocol is supported (only http/https)
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return InvalidProtocolError{Protocol: parsedURL.Scheme}
	}

	// Create HTTP client with timeout to prevent hanging requests
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Perform HTTP GET request
	resp, err := client.Get(fileURL)
	if err != nil {
		return ErrConnectionFailed
	}
	defer resp.Body.Close() 

	// Handle specific HTTP status codes
	if resp.StatusCode == http.StatusNotFound {
		return ErrFileNotFound
	}

	// Handle other error status codes (4xx, 5xx)
	if resp.StatusCode >= 400 {
		return ErrDownloadFailed
	}

	// Create output file for saving downloaded content
	resultFile, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer resultFile.Close() // Ensure file is closed

	// Copy response body to file efficiently
	_, err = io.Copy(resultFile, resp.Body)
	if err != nil {
		return fmt.Errorf("error copying data: %w", err)
	}

	return nil
}

func main() {
	var fileURL, fileName string

	flag.StringVar(&fileURL, "url", "", "url must be filled")
	flag.StringVar(&fileName, "output", "", "file name must be filled")

	flag.Parse()

	if fileURL == "" || fileName == "" {
		fmt.Println("Error: both --url and --output flags are required")
		flag.Usage()
		os.Exit(1)
	}

	err := downloadFile(fileURL, fileName)
	if err != nil {
		// Handle specific error types with user-friendly messages
		switch {
		case errors.As(err, &InvalidProtocolError{}):
			var protErr InvalidProtocolError
			if errors.As(err, &protErr) {
				fmt.Printf("Error: %s protocol is not supported. Please use http or https.\n", protErr.Protocol)
			}
		case errors.Is(err, ErrInvalidURL):
			fmt.Printf("Error: %s. Please enter a valid URL.\n", err.Error())
		case errors.Is(err, ErrConnectionFailed):
			fmt.Printf("Error: %s. Please check your connection and try again.\n", err.Error())
		case errors.Is(err, ErrDownloadFailed):
			fmt.Printf("Error: %s. Please ensure the file is available for download.\n", err.Error())
		case errors.Is(err, ErrFileNotFound):
			fmt.Printf("Error: %s. Please check the URL and ensure the file exists.\n", err.Error())
		default:
			// Handle any other unexpected errors
			fmt.Printf("Error: %s\n", err.Error())
		}
		os.Exit(1)
	}
}
