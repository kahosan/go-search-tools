package util

import (
	"os"
	"path/filepath"
	"tools/pkg/errors"
)

func GetAllFileName(path string) ([]string, *errors.IError) {
	filesNames := make([]string, 0)

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			filesNames = append(filesNames, path)
		}

		return nil
	})

	if err != nil {
		return nil, errors.NewError("GetAllFileNameError", err)
	}

	return filesNames, nil
}
