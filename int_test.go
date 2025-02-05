package pgconv

import (
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInt2(t *testing.T) {
	t.Run("Int2ToInt16", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v pgtype.Int2
			require.Equal(t, int16(0), Int2ToInt16(v))
		})

		t.Run("value", func(t *testing.T) {
			v := pgtype.Int2{Int16: 123, Valid: true}
			require.Equal(t, v.Int16, Int2ToInt16(v))
		})
	})

	t.Run("Int2ToInt16Ptr", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v pgtype.Int2
			assert.Nil(t, Int2ToInt16Ptr(v))
		})

		t.Run("value", func(t *testing.T) {
			v := pgtype.Int2{Int16: 123, Valid: true}
			require.Equal(t, v.Int16, *Int2ToInt16Ptr(v))
		})
	})

	t.Run("Int2FromInt16", func(t *testing.T) {
		t.Run("value", func(t *testing.T) {
			v := int16(123)
			require.Equal(t, pgtype.Int2{Int16: v, Valid: true}, Int2FromInt16(v))
		})
	})

	t.Run("Int2FromInt16Ptr", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v *int16
			require.Equal(t, pgtype.Int2{}, Int2FromInt16Ptr(v))
		})

		t.Run("value", func(t *testing.T) {
			v := int16(123)
			require.Equal(t, pgtype.Int2{Int16: v, Valid: true}, Int2FromInt16Ptr(&v))
		})
	})
}

func TestInt4(t *testing.T) {
	t.Run("Int4ToInt32", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v pgtype.Int4
			require.Equal(t, int32(0), Int4ToInt32(v))
		})

		t.Run("value", func(t *testing.T) {
			v := pgtype.Int4{Int32: 123, Valid: true}
			require.Equal(t, v.Int32, Int4ToInt32(v))
		})
	})

	t.Run("Int4ToInt32Ptr", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v pgtype.Int4
			assert.Nil(t, Int4ToInt32Ptr(v))
		})

		t.Run("value", func(t *testing.T) {
			v := pgtype.Int4{Int32: 123, Valid: true}
			require.Equal(t, v.Int32, *Int4ToInt32Ptr(v))
		})
	})

	t.Run("Int4FromInt32", func(t *testing.T) {
		t.Run("value", func(t *testing.T) {
			v := int32(123)
			require.Equal(t, pgtype.Int4{Int32: v, Valid: true}, Int4FromInt32(v))
		})
	})

	t.Run("Int4FromInt32Ptr", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v *int32
			require.Equal(t, pgtype.Int4{}, Int4FromInt32Ptr(v))
		})

		t.Run("value", func(t *testing.T) {
			v := int32(123)
			require.Equal(t, pgtype.Int4{Int32: v, Valid: true}, Int4FromInt32Ptr(&v))
		})
	})
}

func TestInt8(t *testing.T) {
	t.Run("Int8ToInt64", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v pgtype.Int8
			require.Equal(t, int64(0), Int8ToInt64(v))
		})

		t.Run("value", func(t *testing.T) {
			v := pgtype.Int8{Int64: 123, Valid: true}
			require.Equal(t, v.Int64, Int8ToInt64(v))
		})
	})

	t.Run("Int8ToInt64Ptr", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v pgtype.Int8
			assert.Nil(t, Int8ToInt64Ptr(v))
		})

		t.Run("value", func(t *testing.T) {
			v := pgtype.Int8{Int64: 123, Valid: true}
			require.Equal(t, v.Int64, *Int8ToInt64Ptr(v))
		})
	})

	t.Run("Int8FromInt64", func(t *testing.T) {
		t.Run("value", func(t *testing.T) {
			v := int64(123)
			require.Equal(t, pgtype.Int8{Int64: v, Valid: true}, Int8FromInt64(v))
		})
	})

	t.Run("Int8FromInt64Ptr", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v *int64
			require.Equal(t, pgtype.Int8{}, Int8FromInt64Ptr(v))
		})

		t.Run("value", func(t *testing.T) {
			v := int64(123)
			require.Equal(t, pgtype.Int8{Int64: v, Valid: true}, Int8FromInt64Ptr(&v))
		})
	})
}
