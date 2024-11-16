package utils

import (
	"errors"
	"path/filepath"
)

func JsonChecking(path string) error {

	if !isFileExtensionValid(path, ".json") {
		return errors.New("the file extension must be a json")
	}
	return nil
}

func isFileExtensionValid(path, ext string) bool {
	return filepath.Ext(path) == ext
}
