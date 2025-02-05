package pgconv

import (
	"github.com/jackc/pgx/v5/pgtype"
)

var EmptyInt2 = pgtype.Int2{}

func FromInt16(v pgtype.Int2) int16 {
	return v.Int16
}

func FromInt16Ptr(v pgtype.Int2) *int16 {
	if !v.Valid {
		return nil
	}
	return &v.Int16
}

func ToInt16(v int16) pgtype.Int2 {
	return pgtype.Int2{
		Valid: true,
		Int16: v,
	}
}

func ToInt16Ptr(v *int16) pgtype.Int2 {
	if v == nil {
		return pgtype.Int2{}
	}

	return pgtype.Int2{
		Valid: true,
		Int16: *v,
	}
}

var EmptyInt4 = pgtype.Int4{}

func FromInt32(v pgtype.Int4) int32 {
	return v.Int32
}

func FromInt32Ptr(v pgtype.Int4) *int32 {
	if !v.Valid {
		return nil
	}
	return &v.Int32
}

func ToInt32(v int32) pgtype.Int4 {
	return pgtype.Int4{
		Valid: true,
		Int32: v,
	}
}

func ToInt32Ptr(v *int32) pgtype.Int4 {
	if v == nil {
		return pgtype.Int4{}
	}

	return pgtype.Int4{
		Valid: true,
		Int32: *v,
	}
}

var EmptyInt8 = pgtype.Int8{}

func FromInt64(v pgtype.Int8) int64 {
	return v.Int64
}

func FromInt64Ptr(v pgtype.Int8) *int64 {
	if !v.Valid {
		return nil
	}
	return &v.Int64
}

func ToInt64(v int64) pgtype.Int8 {
	return pgtype.Int8{
		Valid: true,
		Int64: v,
	}
}

func ToInt64Ptr(v *int64) pgtype.Int8 {
	if v == nil {
		return pgtype.Int8{}
	}

	return pgtype.Int8{
		Valid: true,
		Int64: *v,
	}
}
