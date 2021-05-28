package files

import (
	"fmt"

	"github.com/gigamono/gigamono/pkg/encodings"
	"github.com/gofrs/uuid"
)

// GenerateObfuscatedFilePath generates an obfuscated filepath using the following scheme.
//
// [workspace-short-uuid]/[system-generated-resource-prefix].[system-generated-short-uuid]
// or
// [workspace-short-uuid]/[system-generated-resource-prefix].[resource-short-uuid].[system-generated-short-uuid]
func GenerateObfuscatedFilePath(
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

	resourcePostfix := encodings.ShortenUUID(randomID)

	return fmt.Sprintf("%v/%v.%v", encodings.ShortenUUID(workspaceID), resourcePrefix, resourcePostfix), nil
}
