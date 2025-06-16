package logic

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// FileType represents the type of file
type FileType string

const (
	TypeDirectory FileType = "dir"
	TypeFile      FileType = "file"
	TypeImage     FileType = "image"
	TypeVideo     FileType = "video"
)

// PathEntry represents a file or directory entry
type PathEntry struct {
	Name      string
	EntryType FileType
	Size      int64
	IsHidden  bool
}

// PathHandler handles file system operations
type PathHandler struct {
	logger *log.Logger
}

// NewPathHandler creates a new PathHandler instance
func NewPathHandler(logger *log.Logger) *PathHandler {
	return &PathHandler{
		logger: logger,
	}
}

// isImageFile checks if the file is an image based on its extension
func isImageFile(name string) bool {
	ext := strings.ToLower(filepath.Ext(name))
	imageExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".bmp":  true,
		".webp": true,
		".svg":  true,
	}
	return imageExts[ext]
}

// isVideoFile checks if the file is a video based on its extension
func isVideoFile(name string) bool {
	ext := strings.ToLower(filepath.Ext(name))
	videoExts := map[string]bool{
		".mp4":  true,
		".avi":  true,
		".mkv":  true,
		".mov":  true,
		".wmv":  true,
		".flv":  true,
		".webm": true,
	}
	return videoExts[ext]
}

// getFileType determines the type of file based on its name and directory status
func getFileType(d fs.DirEntry, name string) FileType {
	if d.IsDir() {
		return TypeDirectory
	}
	if isImageFile(name) {
		return TypeImage
	}
	if isVideoFile(name) {
		return TypeVideo
	}
	return TypeFile
}

// GetEntries returns a slice of PathEntry for the given directory
func (p *PathHandler) GetEntries(dir string, showHidden bool) ([]PathEntry, error) {
	var entries []PathEntry

	// Read the directory contents
	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	// If directory is empty, return empty slice
	if len(dirEntries) == 0 {
		return entries, nil
	}

	// Process each entry
	for _, d := range dirEntries {
		name := d.Name()
		isHidden := strings.HasPrefix(name, ".")

		// Skip hidden files if not requested
		if !showHidden && isHidden {
			continue
		}

		// Get file info
		info, err := d.Info()
		if err != nil {
			p.logger.Printf("Error getting file info for %s: %v", name, err)
			continue
		}

		entry := PathEntry{
			Name:      name,
			EntryType: getFileType(d, name),
			Size:      info.Size(),
			IsHidden:  isHidden,
		}

		entries = append(entries, entry)
	}

	return entries, nil
} 