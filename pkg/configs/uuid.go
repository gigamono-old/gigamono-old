package configs

import (
	"github.com/gofrs/uuid"
)

// UUID aliases gofrs/UUID type to allow custom unmarhsalling.
type UUID uuid.UUID
