package pgconv

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNullUUID(t *testing.T) {
	t.Run("ToNullUUID", func(t *testing.T) {
		t.Run("valid", func(t *testing.T) {
			v := uuid.New()
			require.Equal(t, uuid.NullUUID{UUID: v, Valid: true}, ToNullUUID(v))
		})
	})
}
