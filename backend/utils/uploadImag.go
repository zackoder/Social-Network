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
	if _, exists := r.Form["avatar"]; !exists {
		return "", fmt.Errorf("nothing")
	}
	file, handler, err := r.FormFile("avatar")
	fmt.Println(handler.Filename)
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

func UploadMsgImg(pyload []byte) string {
	delimiter := []byte("::")
	parts := bytes.SplitN(pyload, delimiter, 2)
	if len(parts) != 2 {
		fmt.Println("wtf")
	}
	metaPart := parts[0]
	filePart := parts[1]

	var meta map[string]string
	err := json.Unmarshal(metaPart, &meta)
	if err != nil {
		fmt.Println("invalid meta data")
	}
	if !strings.Contains(meta["mime"], "image/") {
		return ""
	}
	os.MkdirAll("uploads", os.ModePerm)

	fileName := fmt.Sprintf("uploads/%d_%s", time.Now().Unix(), meta["name"])

	if err := os.WriteFile(fileName, filePart, 0644); err != nil {
		fmt.Println("writing file error ")
	}

	return fileName
}
