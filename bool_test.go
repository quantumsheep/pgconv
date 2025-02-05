package pgconv

import (
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBool(t *testing.T) {
	t.Run("ToBool", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v pgtype.Bool
			require.Equal(t, false, ToBool(v))
		})

		t.Run("true", func(t *testing.T) {
			v := pgtype.Bool{Bool: true, Valid: true}
			require.Equal(t, v.Bool, ToBool(v))
		})

		t.Run("false", func(t *testing.T) {
			v := pgtype.Bool{Bool: false, Valid: true}
			require.Equal(t, v.Bool, ToBool(v))
		})
	})

	t.Run("ToBoolPtr", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v pgtype.Bool
			assert.Nil(t, ToBoolPtr(v))
		})

		t.Run("true", func(t *testing.T) {
			v := pgtype.Bool{Bool: true, Valid: true}
			require.Equal(t, v.Bool, *ToBoolPtr(v))
		})

		t.Run("false", func(t *testing.T) {
			v := pgtype.Bool{Bool: false, Valid: true}
			require.Equal(t, v.Bool, *ToBoolPtr(v))
		})
	})

	t.Run("FromBool", func(t *testing.T) {
		t.Run("true", func(t *testing.T) {
			v := true
			require.Equal(t, pgtype.Bool{Bool: v, Valid: true}, FromBool(v))
		})

		t.Run("false", func(t *testing.T) {
			v := false
			require.Equal(t, pgtype.Bool{Bool: v, Valid: true}, FromBool(v))
		})
	})

	t.Run("FromBoolPtr", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v *bool
			require.Equal(t, pgtype.Bool{}, FromBoolPtr(v))
		})

		t.Run("true", func(t *testing.T) {
			v := true
			require.Equal(t, pgtype.Bool{Bool: v, Valid: true}, FromBoolPtr(&v))
		})

		t.Run("false", func(t *testing.T) {
			v := false
			require.Equal(t, pgtype.Bool{Bool: v, Valid: true}, FromBoolPtr(&v))
		})
	})
}
