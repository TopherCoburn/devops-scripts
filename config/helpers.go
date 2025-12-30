package helpers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// GenerateRandomString generates a random string of length 16.
func GenerateRandomString() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// GetFileName returns the file name from the given full path.
func GetFileName(fullPath string) string {
	return filepath.Base(fullPath)
}

// GetFileNameWithoutExtension returns the file name without extension from the given full path.
func GetFileNameWithoutExtension(fullPath string) string {
	return strings.TrimSuffix(GetFileName(fullPath), filepath.Ext(fullPath))
}

// GetDateInYYYYMMDD returns the current date in the format YYYYMMDD.
func GetDateInYYYYMMDD() string {
	return time.Now().Format("20060102")
}

// LogError logs the error message and exits the program.
func LogError(err error) {
	log.Fatal(err)
}

// IsFileExists checks if the given file exists.
func IsFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

// IsDirectoryExists checks if the given directory exists.
func IsDirectoryExists(dirPath string) bool {
	_, err := os.Stat(dirPath)
	return err == nil && !IsFileExists(dirPath)
}

// CreateDirectory creates a new directory with the given path.
func CreateDirectory(dirPath string) error {
	return os.MkdirAll(dirPath, 0755)
}

// GetEnvironmentVariable returns the value of the given environment variable.
func GetEnvironmentVariable(varName string) string {
	return os.Getenv(varName)
}

// PrintRed prints the message in red color.
func PrintRed(msg string) {
	fmt.Printf("\033[91m%s\033[0m\n", msg)
}

// PrintGreen prints the message in green color.
func PrintGreen(msg string) {
	fmt.Printf("\033[92m%s\033[0m\n", msg)
}

// PrintYellow prints the message in yellow color.
func PrintYellow(msg string) {
	fmt.Printf("\033[93m%s\033[0m\n", msg)
}