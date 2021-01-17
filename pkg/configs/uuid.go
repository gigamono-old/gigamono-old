package configs

import (
	"github.com/gofrs/uuid"
)

// UUID aliases Google's UUID type to allow custom unmarhsalling.
type UUID uuid.UUID
