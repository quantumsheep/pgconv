package pgconv

import (
	"github.com/jackc/pgx/v5/pgtype"
)

var EmptyText = pgtype.Text{}

func TextToString(v pgtype.Text) string {
	return v.String
}

func TextToStringPtr(v pgtype.Text) *string {
	if !v.Valid {
		return nil
	}
	return &v.String
}

func TextFromString(v string) pgtype.Text {
	return pgtype.Text{
		Valid:  true,
		String: v,
	}
}

func TextFromStringPtr(v *string) pgtype.Text {
	if v == nil {
		return pgtype.Text{}
	}

	return pgtype.Text{
		Valid:  true,
		String: *v,
	}
}
