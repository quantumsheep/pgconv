package pgconv

import (
	"github.com/jackc/pgx/v5/pgtype"
)

var EmptyInt2 = pgtype.Int2{}

func FromInt2(v pgtype.Int2) int16 {
	return v.Int16
}

func FromInt2Ptr(v pgtype.Int2) *int16 {
	if !v.Valid {
		return nil
	}
	return &v.Int16
}

func ToInt2(v int16) pgtype.Int2 {
	return pgtype.Int2{
		Valid: true,
		Int16: v,
	}
}

func ToInt2Ptr(v *int16) pgtype.Int2 {
	if v == nil {
		return pgtype.Int2{}
	}

	return pgtype.Int2{
		Valid: true,
		Int16: *v,
	}
}

var EmptyInt4 = pgtype.Int4{}

func FromInt4(v pgtype.Int4) int32 {
	return v.Int32
}

func FromInt4Ptr(v pgtype.Int4) *int32 {
	if !v.Valid {
		return nil
	}
	return &v.Int32
}

func ToInt4(v int32) pgtype.Int4 {
	return pgtype.Int4{
		Valid: true,
		Int32: v,
	}
}

func ToInt4Ptr(v *int32) pgtype.Int4 {
	if v == nil {
		return pgtype.Int4{}
	}

	return pgtype.Int4{
		Valid: true,
		Int32: *v,
	}
}

var EmptyInt8 = pgtype.Int8{}

func FromInt8(v pgtype.Int8) int64 {
	return v.Int64
}

func FromInt8Ptr(v pgtype.Int8) *int64 {
	if !v.Valid {
		return nil
	}
	return &v.Int64
}

func ToInt8(v int64) pgtype.Int8 {
	return pgtype.Int8{
		Valid: true,
		Int64: v,
	}
}

func ToInt8Ptr(v *int64) pgtype.Int8 {
	if v == nil {
		return pgtype.Int8{}
	}

	return pgtype.Int8{
		Valid: true,
		Int64: *v,
	}
}
