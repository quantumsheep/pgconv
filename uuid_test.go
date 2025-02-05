package pgconv

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNullUUID(t *testing.T) {
	t.Run("FromNullUUID", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v uuid.NullUUID
			require.Equal(t, uuid.Nil, FromNullUUID(v))
		})

		t.Run("valid", func(t *testing.T) {
			v := uuid.NullUUID{UUID: uuid.New(), Valid: true}
			require.Equal(t, v.UUID, FromNullUUID(v))
		})
	})

	t.Run("ToNullUUID", func(t *testing.T) {
		t.Run("valid", func(t *testing.T) {
			v := uuid.New()
			require.Equal(t, uuid.NullUUID{UUID: v, Valid: true}, ToNullUUID(v))
		})
	})
}
