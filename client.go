package mpu

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

// Config enable various options for the uploader
type Config struct {
	Gzip bool
}

// DefaultConfig defaults for the uploader
func DefaultConfig() *Config {
	return &Config{true}
}

// UploaderBuilder builds upload requests with a basic setup
type UploaderBuilder struct {
	http.Client
	config *Config
}

// Uploader make a new upload builder
func Uploader(config *Config) *UploaderBuilder {
	return &UploaderBuilder{http.Client{}, config}
}

// NewFileRequest build a new file upload http request with a multi part form consisting of the extra fields
// and the contents of the file loaded from the supplied path.
func (ub *UploaderBuilder) NewFileRequest(uri string, extraParams map[string]string, paramName, filePath string) (*http.Request, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)

	for key, val := range extraParams {
		_ = writer.WriteField(key, val)
	}

	part, err := writer.CreateFormFile(paramName, file.Name())
	if err != nil {
		return nil, err
	}

	written, err := io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	log.Printf("wrote %d", written)

	return http.NewRequest("POST", uri, body)
}
