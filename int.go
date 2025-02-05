package pgconv

import (
	"github.com/jackc/pgx/v5/pgtype"
)

var EmptyInt2 = pgtype.Int2{}

func Int2ToInt16(v pgtype.Int2) int16 {
	return v.Int16
}

func Int2ToInt16Ptr(v pgtype.Int2) *int16 {
	if !v.Valid {
		return nil
	}
	return &v.Int16
}

func Int2FromInt16(v int16) pgtype.Int2 {
	return pgtype.Int2{
		Valid: true,
		Int16: v,
	}
}

func Int2FromInt16Ptr(v *int16) pgtype.Int2 {
	if v == nil {
		return pgtype.Int2{}
	}

	return pgtype.Int2{
		Valid: true,
		Int16: *v,
	}
}

var EmptyInt4 = pgtype.Int4{}

func Int4ToInt32(v pgtype.Int4) int32 {
	return v.Int32
}

func Int4ToInt32Ptr(v pgtype.Int4) *int32 {
	if !v.Valid {
		return nil
	}
	return &v.Int32
}

func Int4FromInt32(v int32) pgtype.Int4 {
	return pgtype.Int4{
		Valid: true,
		Int32: v,
	}
}

func Int4FromInt32Ptr(v *int32) pgtype.Int4 {
	if v == nil {
		return pgtype.Int4{}
	}

	return pgtype.Int4{
		Valid: true,
		Int32: *v,
	}
}

var EmptyInt8 = pgtype.Int8{}

func Int8ToInt64(v pgtype.Int8) int64 {
	return v.Int64
}

func Int8ToInt64Ptr(v pgtype.Int8) *int64 {
	if !v.Valid {
		return nil
	}
	return &v.Int64
}

func Int8FromInt64(v int64) pgtype.Int8 {
	return pgtype.Int8{
		Valid: true,
		Int64: v,
	}
}

func Int8FromInt64Ptr(v *int64) pgtype.Int8 {
	if v == nil {
		return pgtype.Int8{}
	}

	return pgtype.Int8{
		Valid: true,
		Int64: *v,
	}
}
