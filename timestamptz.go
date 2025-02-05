package pgconv

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

var EmptyTimestamptz = pgtype.Timestamptz{}

func FromTimestamptz(v pgtype.Timestamptz) time.Time {
	return v.Time
}

func FromTimestamptzPtr(v pgtype.Timestamptz) *time.Time {
	if !v.Valid {
		return nil
	}
	return &v.Time
}

func ToTimestamptz(v time.Time) pgtype.Timestamptz {
	return pgtype.Timestamptz{
		Valid: true,
		Time:  v,
	}
}

func ToTimestamptzPtr(v *time.Time) pgtype.Timestamptz {
	if v == nil {
		return pgtype.Timestamptz{}
	}

	return pgtype.Timestamptz{
		Valid: true,
		Time:  *v,
	}
}
