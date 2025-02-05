package pgconv

import (
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestText(t *testing.T) {
	t.Run("FromText", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v pgtype.Text
			require.Equal(t, "", FromText(v))
		})

		t.Run("valid", func(t *testing.T) {
			v := pgtype.Text{String: "test", Valid: true}
			require.Equal(t, v.String, FromText(v))
		})
	})

	t.Run("FromTextPtr", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v pgtype.Text
			assert.Nil(t, FromTextPtr(v))
		})

		t.Run("valid", func(t *testing.T) {
			v := pgtype.Text{String: "test", Valid: true}
			require.Equal(t, v.String, *FromTextPtr(v))
		})
	})

	t.Run("ToText", func(t *testing.T) {
		t.Run("valid", func(t *testing.T) {
			v := "test"
			require.Equal(t, pgtype.Text{String: v, Valid: true}, ToText(v))
		})
	})

	t.Run("ToTextPtr", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v *string
			assert.Equal(t, pgtype.Text{}, ToTextPtr(v))
		})

		t.Run("valid", func(t *testing.T) {
			v := "test"
			require.Equal(t, pgtype.Text{String: v, Valid: true}, ToTextPtr(&v))
		})
	})
}
