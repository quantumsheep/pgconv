package pgconv

import (
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInt2(t *testing.T) {
	t.Run("FromInt16", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v pgtype.Int2
			require.Equal(t, int16(0), FromInt16(v))
		})

		t.Run("value", func(t *testing.T) {
			v := pgtype.Int2{Int16: 123, Valid: true}
			require.Equal(t, v.Int16, FromInt16(v))
		})
	})

	t.Run("FromInt16Ptr", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v pgtype.Int2
			assert.Nil(t, FromInt16Ptr(v))
		})

		t.Run("value", func(t *testing.T) {
			v := pgtype.Int2{Int16: 123, Valid: true}
			require.Equal(t, v.Int16, *FromInt16Ptr(v))
		})
	})

	t.Run("ToInt16", func(t *testing.T) {
		t.Run("value", func(t *testing.T) {
			v := int16(123)
			require.Equal(t, pgtype.Int2{Int16: v, Valid: true}, ToInt16(v))
		})
	})

	t.Run("ToInt16Ptr", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v *int16
			require.Equal(t, pgtype.Int2{}, ToInt16Ptr(v))
		})

		t.Run("value", func(t *testing.T) {
			v := int16(123)
			require.Equal(t, pgtype.Int2{Int16: v, Valid: true}, ToInt16Ptr(&v))
		})
	})
}

func TestInt4(t *testing.T) {
	t.Run("FromInt32", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v pgtype.Int4
			require.Equal(t, int32(0), FromInt32(v))
		})

		t.Run("value", func(t *testing.T) {
			v := pgtype.Int4{Int32: 123, Valid: true}
			require.Equal(t, v.Int32, FromInt32(v))
		})
	})

	t.Run("FromInt32Ptr", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v pgtype.Int4
			assert.Nil(t, FromInt32Ptr(v))
		})

		t.Run("value", func(t *testing.T) {
			v := pgtype.Int4{Int32: 123, Valid: true}
			require.Equal(t, v.Int32, *FromInt32Ptr(v))
		})
	})

	t.Run("ToInt32", func(t *testing.T) {
		t.Run("value", func(t *testing.T) {
			v := int32(123)
			require.Equal(t, pgtype.Int4{Int32: v, Valid: true}, ToInt32(v))
		})
	})

	t.Run("ToInt32Ptr", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v *int32
			require.Equal(t, pgtype.Int4{}, ToInt32Ptr(v))
		})

		t.Run("value", func(t *testing.T) {
			v := int32(123)
			require.Equal(t, pgtype.Int4{Int32: v, Valid: true}, ToInt32Ptr(&v))
		})
	})
}

func TestInt8(t *testing.T) {
	t.Run("FromInt64", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v pgtype.Int8
			require.Equal(t, int64(0), FromInt64(v))
		})

		t.Run("value", func(t *testing.T) {
			v := pgtype.Int8{Int64: 123, Valid: true}
			require.Equal(t, v.Int64, FromInt64(v))
		})
	})

	t.Run("FromInt64Ptr", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v pgtype.Int8
			assert.Nil(t, FromInt64Ptr(v))
		})

		t.Run("value", func(t *testing.T) {
			v := pgtype.Int8{Int64: 123, Valid: true}
			require.Equal(t, v.Int64, *FromInt64Ptr(v))
		})
	})

	t.Run("ToInt64", func(t *testing.T) {
		t.Run("value", func(t *testing.T) {
			v := int64(123)
			require.Equal(t, pgtype.Int8{Int64: v, Valid: true}, ToInt64(v))
		})
	})

	t.Run("ToInt64Ptr", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v *int64
			require.Equal(t, pgtype.Int8{}, ToInt64Ptr(v))
		})

		t.Run("value", func(t *testing.T) {
			v := int64(123)
			require.Equal(t, pgtype.Int8{Int64: v, Valid: true}, ToInt64Ptr(&v))
		})
	})
}
