package mpu

import (
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

type Config struct {
	Gzip bool
}

func DefaultConfig() *Config {
	return &Config{true}
}

type UploaderBuilder struct {
	config *Config
}

func Uploader(config *Config) *UploaderBuilder {
	return &UploaderBuilder{config}
}

func (ub *UploaderBuilder) NewFileRequest(uri string, extraParams map[string]string, paramName, filePath string) (*http.Request, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}

	file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile(paramName, fi.Name())
	if err != nil {
		return nil, err
	}

	part.Write(fileContents)

	for key, val := range extraParams {
		_ = writer.WriteField(key, val)
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	return http.NewRequest("POST", uri, body)
}
