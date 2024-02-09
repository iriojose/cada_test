package utils

import (
	"bytes"
	"io"
	"net/http"
)

// make an http call
func HttpGET(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	byteValue, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return byteValue, nil
}

func HttpPOST(url string, contentType string, body []byte) ([]byte, error) {
	response, err := http.Post(url, contentType, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	byteValue, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return byteValue, nil
}
