package pgconv

import (
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTimestamptz(t *testing.T) {
	t.Run("ToTimestamptz", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v pgtype.Timestamptz
			require.Equal(t, time.Time{}, TimestamptzToTime(v))
		})

		t.Run("valid", func(t *testing.T) {
			v := pgtype.Timestamptz{Time: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), Valid: true}
			require.Equal(t, v.Time, TimestamptzToTime(v))
		})
	})

	t.Run("ToTimestamptzPtr", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v pgtype.Timestamptz
			assert.Nil(t, TimestamptzToTimePtr(v))
		})

		t.Run("valid", func(t *testing.T) {
			v := pgtype.Timestamptz{Time: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), Valid: true}
			require.Equal(t, v.Time, *TimestamptzToTimePtr(v))
		})
	})

	t.Run("FromTimestamptz", func(t *testing.T) {
		t.Run("valid", func(t *testing.T) {
			v := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
			require.Equal(t, pgtype.Timestamptz{Time: v, Valid: true}, TimestamptzFromTime(v))
		})
	})

	t.Run("FromTimestamptzPtr", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v *time.Time
			assert.Equal(t, pgtype.Timestamptz{}, TimestamptzFromTimePtr(v))
		})

		t.Run("valid", func(t *testing.T) {
			v := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
			require.Equal(t, pgtype.Timestamptz{Time: v, Valid: true}, TimestamptzFromTimePtr(&v))
		})
	})
}
