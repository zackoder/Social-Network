package utils

import (
	"fmt"
	"io"
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
		fmt.Println(err)
		return "", err
	}
	os.MkdirAll("uploads", os.ModePerm)
	defer file.Close()
	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), handler.Filename)
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
	return filePath, nil
}
