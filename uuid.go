package pgconv

import (
	"github.com/google/uuid"
)

func FromNullUUID(v uuid.NullUUID) uuid.UUID {
	return v.UUID
}

func ToNullUUID(v uuid.UUID) uuid.NullUUID {
	return uuid.NullUUID{
		Valid: true,
		UUID:  v,
	}
}
