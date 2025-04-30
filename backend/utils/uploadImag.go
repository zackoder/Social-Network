package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func UploadImage(r *http.Request) (string, error) {
	file, handler, err := r.FormFile("avatar")
	if err != nil {
		if file == nil {
			return "", nil
		}
		fmt.Println(err)
		return "", err
	}
	os.MkdirAll("uploads", os.ModePerm)
	defer file.Close()
	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), handler.Filename)
	if !CheckExtension(fileName) {
		return "", fmt.Errorf("invalid file type")
	}
	fmt.Println(handler)
	filePath := "./uploads/" + fileName

	out, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
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

	metaPart := parts[0]
	filePart := parts[1]

	err := json.Unmarshal(metaPart, &message)
	if err != nil {
		fmt.Println("invalid meta data", err)
		return message, fmt.Errorf("Check your data")
	}

	if !strings.Contains(message.Mime, "image/") {
		return message, fmt.Errorf("invalid file type you can only send images")
	}
	os.MkdirAll("uploads", os.ModePerm)

	message.Filename = fmt.Sprintf("uploads/%d_%s", time.Now().Unix(), message.Filename)

	if err := os.WriteFile(message.Filename, filePart, 0644); err != nil {
		fmt.Println("writing file error ", err)
		return message, fmt.Errorf("internal sercer error")
	}
	message.Filename = "/" + message.Filename
	return message, nil
}
