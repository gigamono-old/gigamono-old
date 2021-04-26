package filestore

import "github.com/gigamono/gigamono/pkg/files"

// SetAvatarsLocation creates filestore location if they don't exist.
// SEC: Avatars filestore uses security by obscurity. Sensitive data should be stored in it.
func SetAvatarsLocation() error {
	return files.OpenOrCreateFolder("avatars")
}
