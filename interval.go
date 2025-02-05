package pgconv

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const (
	hoursInDay   = 24
	daysInMonth  = 30 // pgtype.Interval considers a month to be 30 days
	hoursInMonth = hoursInDay * daysInMonth

	hourInMicroseconds = 1000000 * 60 * 60
)

var EmptyInterval = pgtype.Interval{}

func IntervalToDuration(v pgtype.Interval) time.Duration {
	if !v.Valid {
		return time.Duration(0)
	}

	duration := time.Duration(0)
	duration += time.Duration(v.Microseconds) * time.Microsecond
	duration += time.Duration(v.Days) * (time.Hour * hoursInDay)
	duration += time.Duration(v.Months) * (time.Hour * hoursInMonth)
	return duration
}

func IntervalToDurationPtr(v pgtype.Interval) *time.Duration {
	if !v.Valid {
		return nil
	}

	duration := IntervalToDuration(v)
	return &duration
}

func IntervalFromDuration(v time.Duration) pgtype.Interval {
	microseconds := v.Microseconds() % hourInMicroseconds
	hours := int64(v.Hours())

	microseconds += (hours % hoursInDay) * hourInMicroseconds

	days := int32(hours/hoursInDay) % daysInMonth
	months := int32(hours / hoursInMonth)

	return pgtype.Interval{
		Valid:        true,
		Microseconds: microseconds,
		Days:         days,
		Months:       months,
	}
}

func IntervalFromDurationPtr(v *time.Duration) pgtype.Interval {
	if v == nil {
		return pgtype.Interval{}
	}

	return IntervalFromDuration(*v)
}
