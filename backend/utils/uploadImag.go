package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func UploadImage(r *http.Request) (string, error) {
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
	// os.WriteFile("test.txt", data[:10], 0o644)

	// Validate file content
	if !CheckExtension(string(data[:10])) {
		return "", fmt.Errorf("invalid file type")
	}
	// Now, write the buffered content to disk
	_, err = io.Copy(out, &fileBytes)
	if err != nil {
		return "", err
	}
	return filePath[1:], nil
}

func UploadMsgImg(pyload []byte) (Message, error) {
	delimiter := []byte("::")
	parts := bytes.SplitN(pyload, delimiter, 2)
	var message Message
	if len(parts) != 2 {
		return message, fmt.Errorf("Send a valid data")
	}
	// if err := os.WriteFile("test.txt", pyload, 0o644); err != nil {
	// 	fmt.Println("writing file error ", err)
	// 	return message, fmt.Errorf("internal sercer error")
	// } else {
	// 	log.Println("file was created")
	// }

	metaPart := parts[0]
	filePart := parts[1]
	log.Println("metta data:", string(metaPart))
	err := json.Unmarshal(metaPart, &message)
	if err != nil {
		fmt.Println("invalid meta data", err)
		return message, fmt.Errorf("Check your data")
	}
	// extention := strings.Split(message.Mime, "/")[1]

	if !CheckExtension(string(filePart[:10])) {
		return message, fmt.Errorf("invalid file type you can only send images")
	}
	os.MkdirAll("uploads", os.ModePerm)

	message.Filename = fmt.Sprintf("uploads/%d_%s", time.Now().Unix(), message.Filename)

	if err := os.WriteFile(message.Filename, filePart, 0o644); err != nil {
		fmt.Println("writing file error ", err)
		return message, fmt.Errorf("internal sercer error")
	}
	message.Filename = "/" + message.Filename
	return message, nil
}

// func checkFileType(extention string, textimage []byte) bool {
// 	if strings.Contains(string(textimage), strings.ToUpper(extention)) {
// 		return true
// 	}
// 	return false
// }
