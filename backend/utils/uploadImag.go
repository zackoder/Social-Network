package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func UploadImage(r *http.Request) (string, error) {
	log.Println("image 1111111111111111111111111")
	file, handler, err := r.FormFile("avatar")
	if err != nil {
		if file == nil {
			return "", nil
		}
		return "", err
	}
	defer file.Close()

	os.MkdirAll("uploads", os.ModePerm)
	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), handler.Filename)
	filePath := "./uploads/" + fileName

	// Create the output file
	out, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// TeeReader will write to out and also let us read the bytes
	var fileBytes bytes.Buffer
	tee := io.TeeReader(file, &fileBytes)

	// Read from tee to consume the content for validation
	data, err := io.ReadAll(tee)
	if err != nil {
		return "", err
	}

	// Optional: write a debug file
	// os.WriteFile("test.txt", data, 0o644)

	// Validate file content
	if !CheckExtension(string(data)) {
		RemoveIMG(out.Name())
		return "", fmt.Errorf("invalid file type")
	}
	// Now, write the buffered content to disk
	_, err = io.Copy(out, &fileBytes)
	if err != nil {
		return "", err
	}
	return filePath[1:], nil
}

// if err := os.WriteFile("test.txt", parts[0], 0o644); err != nil {
// 	fmt.Println("writing file error ", err)
// 	return message, fmt.Errorf("internal sercer error")
// }

func UploadMsgImg(payload []byte) (Message, error) {
	if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
		log.Println("creating uploads folder", err)
	}

	delimiter := []byte("::")
	parts := bytes.SplitN(payload, delimiter, 2)

	var message Message
	if len(parts) != 2 {
		return message, fmt.Errorf("send valid data")
	}

	metaPart := parts[0]
	filePart := parts[1]

	err := json.Unmarshal(metaPart, &message)
	if err != nil {
		fmt.Println("invalid metadata:", err)
		return message, fmt.Errorf("check your data")
	}

	// Validate file content type
	contentType := http.DetectContentType(filePart)
	if !strings.HasPrefix(contentType, "image/") {
		return message, fmt.Errorf("invalid file type, only images allowed")
	}

	// Ensure filename

	// Save file
	message.Filename = fmt.Sprintf("uploads/%d_%s", time.Now().Unix(), message.Filename)
	if err := os.WriteFile(message.Filename, filePart, 0o644); err != nil {
		fmt.Println("writing file error", err)
		return message, fmt.Errorf("internal server error")
	}
	if message.Filename != "" {
		message.Filename = "/" + message.Filename
	}
	return message, nil
}

// func checkFileType(extention string, textimage []byte) bool {
// 	if strings.Contains(string(textimage), strings.ToUpper(extention)) {
// 		return true
// 	}
// 	return false
// }
