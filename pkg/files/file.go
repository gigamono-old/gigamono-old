package files

import (
	"os"
	"path/filepath"
)

// OpenOrCreateFolder opens specified folder. Creates folder and its path if they do not exist.
func OpenOrCreateFolder(path string) error {
	// Had to pass 0777 here. https://stackoverflow.com/a/58403214/3984876
	return os.MkdirAll(path, 0777)
}

// OpenOrCreateFile opens specified file. Creates file and its path if they do not exist.
func OpenOrCreateFile(path string, isAppend bool) (*os.File, error) {
	var append int

	if isAppend {
		append = os.O_APPEND
	} else {
		append = 0
	}

	// Had to pass 0777 here. https://stackoverflow.com/a/58403214/3984876
	if err := os.MkdirAll(filepath.Dir(path), 0777); err != nil {
		return nil, err
	}

	// Sec: No point giving other non-grp users permissions.
	return os.OpenFile(path, os.O_CREATE|os.O_WRONLY|append, 0660)
}

// OpenFile opens specified file.
func OpenFile(path string, isAppend bool) (*os.File, error) {
	var append int

	if isAppend {
		append = os.O_APPEND
	} else {
		append = 0
	}

	// Sec: No point giving other non-grp users permissions.
	return os.OpenFile(path, os.O_CREATE|os.O_WRONLY|append, 0660)
}
