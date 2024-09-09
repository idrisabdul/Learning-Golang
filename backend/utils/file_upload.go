package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type ResponseFile struct {
	Msg   Link    `json:"msg"`
	Info  *string `json:"info"`
	Error bool    `json:"error"`
}

type Link struct {
	Host     string
	Path     string
	RawQuery string
}

func UploadFile(c *fiber.Ctx, fh *multipart.FileHeader, alter_filename string) (string, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	fileData, err := fh.Open()
	if err != nil {
		return "", c.Status(fiber.StatusInternalServerError).SendString("error")
	}
	defer fileData.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("fileUpload", alter_filename)
	if err != nil {
		return "", fmt.Errorf("failed to create form file: %v", err)
	}
	if _, err := io.Copy(part, fileData); err != nil {
		return "", fmt.Errorf("failed to copy file data: %v", err)
	}
	if err := writer.Close(); err != nil {
		return "", fmt.Errorf("failed to close multipart writer: %v", err)
	}

	req, err := http.NewRequest("POST", os.Getenv("URL_MINIO")+"/upload", body)
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %v", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to perform HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return GetFileURL(c, alter_filename)
}

func GetFileURL(c *fiber.Ctx, fh string) (string, error) {
	// s.log.Print("run geturl")
	fh = strings.Replace(fh, "%20", " ", -1)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	if err := writer.WriteField("fileUpload", fh); err != nil {
		return "", fmt.Errorf("failed to write field to multipart request: %v", err)
	}
	if err := writer.Close(); err != nil {
		return "", fmt.Errorf("failed to close multipart writer: %v", err)
	}

	req, err := http.NewRequest("POST", os.Getenv("URL_MINIO")+"/download_file_url", body)
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %v", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to perform HTTP request: %v", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	var response *ResponseFile
	if err := json.Unmarshal([]byte(responseBody), &response); err != nil {
		return "", fmt.Errorf("failed to unmarshal JSON response: %v", err)
	}

	url := formatURL(response.Msg.Host, response.Msg.Path, response.Msg.RawQuery)
	return fetchGetBase64("http://" + url)
	// Get File Base64
}

func formatURL(host, path, rawQuery string) string {
	return fmt.Sprintf("%v%v?%v", host, path, rawQuery)
}

func GetFileExtension(filename string) (string, string) {
	parts := strings.Split(filename, ".")
	if len(parts) > 1 {
		return parts[0], parts[1]
	}
	return "", ""
}

func fetchGetBase64(url string) (string, error) {
	// Make an HTTP GET request to fetch the file
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch file: %w", err)
	}
	defer resp.Body.Close()

	// Check for a successful response
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch file: status code %d", resp.StatusCode)
	}

	// Read the response body
	fileBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	// Encode the file bytes to Base64
	encodedString := base64.StdEncoding.EncodeToString(fileBytes)

	return encodedString, nil
}
