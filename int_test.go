package pgconv

import (
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInt2(t *testing.T) {
	t.Run("FromInt2", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v pgtype.Int2
			require.Equal(t, int16(0), FromInt2(v))
		})

		t.Run("value", func(t *testing.T) {
			v := pgtype.Int2{Int16: 123, Valid: true}
			require.Equal(t, v.Int16, FromInt2(v))
		})
	})

	t.Run("FromInt2Ptr", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v pgtype.Int2
			assert.Nil(t, FromInt2Ptr(v))
		})

		t.Run("value", func(t *testing.T) {
			v := pgtype.Int2{Int16: 123, Valid: true}
			require.Equal(t, v.Int16, *FromInt2Ptr(v))
		})
	})

	t.Run("ToInt2", func(t *testing.T) {
		t.Run("value", func(t *testing.T) {
			v := int16(123)
			require.Equal(t, pgtype.Int2{Int16: v, Valid: true}, ToInt2(v))
		})
	})

	t.Run("ToInt2Ptr", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v *int16
			require.Equal(t, pgtype.Int2{}, ToInt2Ptr(v))
		})

		t.Run("value", func(t *testing.T) {
			v := int16(123)
			require.Equal(t, pgtype.Int2{Int16: v, Valid: true}, ToInt2Ptr(&v))
		})
	})
}

func TestInt4(t *testing.T) {
	t.Run("FromInt4", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v pgtype.Int4
			require.Equal(t, int32(0), FromInt4(v))
		})

		t.Run("value", func(t *testing.T) {
			v := pgtype.Int4{Int32: 123, Valid: true}
			require.Equal(t, v.Int32, FromInt4(v))
		})
	})

	t.Run("FromInt4Ptr", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v pgtype.Int4
			assert.Nil(t, FromInt4Ptr(v))
		})

		t.Run("value", func(t *testing.T) {
			v := pgtype.Int4{Int32: 123, Valid: true}
			require.Equal(t, v.Int32, *FromInt4Ptr(v))
		})
	})

	t.Run("ToInt4", func(t *testing.T) {
		t.Run("value", func(t *testing.T) {
			v := int32(123)
			require.Equal(t, pgtype.Int4{Int32: v, Valid: true}, ToInt4(v))
		})
	})

	t.Run("ToInt4Ptr", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v *int32
			require.Equal(t, pgtype.Int4{}, ToInt4Ptr(v))
		})

		t.Run("value", func(t *testing.T) {
			v := int32(123)
			require.Equal(t, pgtype.Int4{Int32: v, Valid: true}, ToInt4Ptr(&v))
		})
	})
}

func TestInt8(t *testing.T) {
	t.Run("FromInt8", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v pgtype.Int8
			require.Equal(t, int64(0), FromInt8(v))
		})

		t.Run("value", func(t *testing.T) {
			v := pgtype.Int8{Int64: 123, Valid: true}
			require.Equal(t, v.Int64, FromInt8(v))
		})
	})

	t.Run("FromInt8Ptr", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v pgtype.Int8
			assert.Nil(t, FromInt8Ptr(v))
		})

		t.Run("value", func(t *testing.T) {
			v := pgtype.Int8{Int64: 123, Valid: true}
			require.Equal(t, v.Int64, *FromInt8Ptr(v))
		})
	})

	t.Run("ToInt8", func(t *testing.T) {
		t.Run("value", func(t *testing.T) {
			v := int64(123)
			require.Equal(t, pgtype.Int8{Int64: v, Valid: true}, ToInt8(v))
		})
	})

	t.Run("ToInt8Ptr", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var v *int64
			require.Equal(t, pgtype.Int8{}, ToInt8Ptr(v))
		})

		t.Run("value", func(t *testing.T) {
			v := int64(123)
			require.Equal(t, pgtype.Int8{Int64: v, Valid: true}, ToInt8Ptr(&v))
		})
	})
}
