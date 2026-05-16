package helpers

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

const (
	MaxUploadSize = 5 << 20 // 5MB
	UploadDir     = "uploads/avatars"
)

var AllowedImageTypes = map[string]bool{
	"image/jpeg": true,
	"image/jpg":  true,
	"image/png":  true,
	"image/webp": true,
}

type UploadResult struct {
	Filename string `json:"filename"`
	URL      string `json:"url"`
	Size     int64  `json:"size"`
}

func ValidateAndSaveImage(file *multipart.FileHeader) (*UploadResult, error) {
	if file.Size > MaxUploadSize {
		return nil, errors.New("file size exceeds maximum limit of 5MB")
	}

	src, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer src.Close()

	buf := make([]byte, 512)
	_, err = src.Read(buf)
	if err != nil {
		return nil, fmt.Errorf("failed to read file header: %w", err)
	}

	contentType := http.DetectContentType(buf)
	if !AllowedImageTypes[contentType] {
		return nil, fmt.Errorf("invalid file type: %s (allowed: jpeg, png, webp)", contentType)
	}

	ext := filepath.Ext(file.Filename)
	uniqueName := uuid.New().String() + ext
	uploadPath := filepath.Join(UploadDir, uniqueName)

	if err := os.MkdirAll(UploadDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create upload directory: %w", err)
	}

	_, err = src.Seek(0, io.SeekStart)
	if err != nil {
		return nil, fmt.Errorf("failed to seek file: %w", err)
	}

	dst, err := os.Create(uploadPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return nil, fmt.Errorf("failed to save file: %w", err)
	}

	return &UploadResult{
		Filename: uniqueName,
		URL:      "/uploads/avatars/" + uniqueName,
		Size:     file.Size,
	}, nil
}

func DeleteAvatar(filename string) error {
	if filename == "" {
		return nil
	}
	path := filepath.Join(UploadDir, filename)
	return os.Remove(path)
}