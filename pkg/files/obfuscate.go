package files

import (
	"fmt"

	"github.com/gigamono/gigamono/pkg/encodings"
	"github.com/gofrs/uuid"
)

// GenerateObfuscatedFilePath generates an obfuscated filepath using the following schemes:
//
// 		[workspace-id-shortened]/[resource-name].[random-id-shortened].[file-extension]
//
// 		[workspace-id-shortened]/[resource-name].[resource-id-shortened].[random-id-shortened].[file-extension]
func GenerateObfuscatedFilePath(
	fileExtension string,
	workspaceID uuid.UUID,
	resourceName string,
	resourceID *uuid.UUID,
) (string, error) {
	resourcePrefix := resourceName
	if resourceID != nil {
		resourcePrefix = resourcePrefix + "." + encodings.ShortenUUID(*resourceID)
	}

	randomID, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"%v/%v.%v.%v",
		encodings.ShortenUUID(workspaceID),
		resourcePrefix,
		encodings.ShortenUUID(randomID),
		fileExtension,
	), nil
}
