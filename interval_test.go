package pgconv

import (
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInterval(t *testing.T) {
	t.Run("IntervalToDuration", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v pgtype.Interval
			require.Equal(t, time.Duration(0), IntervalToDuration(v))
		})

		t.Run("valid", func(t *testing.T) {
			v := pgtype.Interval{
				Valid:        true,
				Microseconds: 1,
				Days:         2,
				Months:       3,
			}

			duration := IntervalToDuration(v)
			require.Equal(t, time.Duration(1)*time.Microsecond+time.Duration(2)*time.Hour*24+time.Duration(3)*time.Hour*24*30, duration)
		})
	})

	t.Run("IntervalToDurationPtr", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v pgtype.Interval
			assert.Nil(t, IntervalToDurationPtr(v))
		})

		t.Run("valid", func(t *testing.T) {
			v := pgtype.Interval{
				Valid:        true,
				Microseconds: 1,
				Days:         2,
				Months:       3,
			}

			duration := IntervalToDurationPtr(v)
			require.NotNil(t, duration)
			require.Equal(t, time.Duration(1)*time.Microsecond+time.Duration(2)*time.Hour*24+time.Duration(3)*time.Hour*24*30, *duration)
		})
	})

	t.Run("IntervalFromDuration", func(t *testing.T) {
		t.Run("valid", func(t *testing.T) {
			v := time.Duration(1)*time.Microsecond + time.Duration(2)*time.Hour*24 + time.Duration(3)*time.Hour*24*30
			require.Equal(t, pgtype.Interval{
				Valid:        true,
				Microseconds: 1,
				Days:         2,
				Months:       3,
			}, IntervalFromDuration(v))
		})
	})

	t.Run("IntervalFromDurationPtr", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v *time.Duration
			assert.Equal(t, pgtype.Interval{}, IntervalFromDurationPtr(v))
		})

		t.Run("valid", func(t *testing.T) {
			v := time.Duration(1)*time.Microsecond + time.Duration(2)*time.Hour*24 + time.Duration(3)*time.Hour*24*30
			require.Equal(t, pgtype.Interval{
				Valid:        true,
				Microseconds: 1,
				Days:         2,
				Months:       3,
			}, IntervalFromDurationPtr(&v))
		})
	})
}
