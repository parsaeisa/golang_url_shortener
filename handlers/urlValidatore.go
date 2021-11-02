package handler

import (
	"errors"
	"strings"
)

func ValidateUrl(url string) (string, error) {

	// check url validation
	if url == "" || !strings.Contains(url, ".") {
		return "", errors.New("please enter a valid URl")
	}

	// validate url
	if !strings.Contains(url, "https://") {
		return "https://" + url, nil
	}
	return url, nil
}
