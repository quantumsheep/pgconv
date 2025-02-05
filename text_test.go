package pgconv

import (
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestText(t *testing.T) {
	t.Run("TextToString", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v pgtype.Text
			require.Equal(t, "", TextToString(v))
		})

		t.Run("valid", func(t *testing.T) {
			v := pgtype.Text{String: "test", Valid: true}
			require.Equal(t, v.String, TextToString(v))
		})
	})

	t.Run("TextToStringPtr", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v pgtype.Text
			assert.Nil(t, TextToStringPtr(v))
		})

		t.Run("valid", func(t *testing.T) {
			v := pgtype.Text{String: "test", Valid: true}
			require.Equal(t, v.String, *TextToStringPtr(v))
		})
	})

	t.Run("TextFromString", func(t *testing.T) {
		t.Run("valid", func(t *testing.T) {
			v := "test"
			require.Equal(t, pgtype.Text{String: v, Valid: true}, TextFromString(v))
		})
	})

	t.Run("TextFromStringPtr", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v *string
			assert.Equal(t, pgtype.Text{}, TextFromStringPtr(v))
		})

		t.Run("valid", func(t *testing.T) {
			v := "test"
			require.Equal(t, pgtype.Text{String: v, Valid: true}, TextFromStringPtr(&v))
		})
	})
}
