package pgconv

import (
	"github.com/google/uuid"
)

func ToNullUUID(v uuid.UUID) uuid.NullUUID {
	return uuid.NullUUID{
		Valid: true,
		UUID:  v,
	}
}
