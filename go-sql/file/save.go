package file

import (
	"os"
	"path/filepath"

	"learn/go-sql/exceptions"
)

func Save(filename string, content []byte) {
	path := filepath.Join("jsons", filename)
	err := os.WriteFile(path, content, 0o644)
	exceptions.HandleBasicException(err)
}
