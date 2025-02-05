package pgconv

import (
	"github.com/jackc/pgx/v5/pgtype"
)

var EmptyText = pgtype.Timestamptz{}

func FromText(v pgtype.Text) string {
	return v.String
}

func FromTextPtr(v pgtype.Text) *string {
	if !v.Valid {
		return nil
	}
	return &v.String
}

func ToText(v string) pgtype.Text {
	return pgtype.Text{
		Valid:  true,
		String: v,
	}
}

func ToTextPtr(v *string) pgtype.Text {
	if v == nil {
		return pgtype.Text{}
	}

	return pgtype.Text{
		Valid:  true,
		String: *v,
	}
}
