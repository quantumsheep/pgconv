package pgconv

import (
	"github.com/jackc/pgx/v5/pgtype"
)

var EmptyBool = pgtype.Bool{}

func ToBool(v pgtype.Bool) bool {
	return v.Bool
}

func ToBoolPtr(v pgtype.Bool) *bool {
	if !v.Valid {
		return nil
	}
	return &v.Bool
}

func FromBool(v bool) pgtype.Bool {
	return pgtype.Bool{
		Valid: true,
		Bool:  v,
	}
}

func FromBoolPtr(v *bool) pgtype.Bool {
	if v == nil {
		return pgtype.Bool{}
	}

	return pgtype.Bool{
		Valid: true,
		Bool:  *v,
	}
}
