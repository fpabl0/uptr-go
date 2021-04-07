package uptr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValueWithFallbacks(t *testing.T) {
	t.Run("bool", func(t *testing.T) {
		r := BoolValueWithFallback(nil, false)
		assert.Equal(t, false, r)
		v := true
		r = BoolValueWithFallback(&v, false)
		assert.Equal(t, true, r)
	})
	t.Run("int", func(t *testing.T) {
		r := IntValueWithFallback(nil, 5)
		assert.Equal(t, 5, r)
		v := 241
		r = IntValueWithFallback(&v, 8)
		assert.Equal(t, 241, r)
	})
	t.Run("float32", func(t *testing.T) {
		r := Float32ValueWithFallback(nil, 7.8)
		assert.Equal(t, float32(7.8), r)
		v := float32(89.31)
		r = Float32ValueWithFallback(&v, 8.5)
		assert.Equal(t, float32(89.31), r)
	})
	t.Run("float64", func(t *testing.T) {
		r := Float64ValueWithFallback(nil, 7.8)
		assert.Equal(t, float64(7.8), r)
		v := float64(89.31)
		r = Float64ValueWithFallback(&v, 8.5)
		assert.Equal(t, float64(89.31), r)
	})
	t.Run("string", func(t *testing.T) {
		r := StringValueWithFallback(nil, "unknown")
		assert.Equal(t, "unknown", r)
		v := "pokemon"
		r = StringValueWithFallback(&v, "not a pokemon")
		assert.Equal(t, "pokemon", r)
	})
}

func TestZeroFallbacks(t *testing.T) {
	assert.Equal(t, false, BoolValue(nil))
	assert.Equal(t, int(0), IntValue(nil))
	assert.Equal(t, float32(0), Float32Value(nil))
	assert.Equal(t, float64(0), Float64Value(nil))
	assert.Equal(t, "", StringValue(nil))

	b := true
	assert.Equal(t, true, BoolValue(&b))
	i := 62
	assert.Equal(t, int(62), IntValue(&i))
	f32 := float32(78.91)
	assert.Equal(t, float32(78.91), Float32Value(&f32))
	f64 := float64(145.82)
	assert.Equal(t, float64(145.82), Float64Value(&f64))
	s := "Hello"
	assert.Equal(t, "Hello", StringValue(&s))
}
