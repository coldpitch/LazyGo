package lg

import (
	"errors"
	"strings"
)

func TokenAt(str, sep string, idx int) (string, error) {
	if str == "" {
		return "", errors.New("string is empty")
	}
	if sep == "" {
		return "", errors.New("separator is empty")
	}
	if idx < 0 {
		return "", errors.New("index cannot be negative")
	}

	parts := strings.Split(str, sep)
	if idx >= len(parts) {
		return "", errors.New("index out of range")
	}

	return parts[idx], nil
}
