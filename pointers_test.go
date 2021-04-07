package uptr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointers(t *testing.T) {
	b := Bool(true)
	assert.Equal(t, true, *b)

	f32 := Float32(9.8)
	assert.Equal(t, float32(9.8), *f32)

	f64 := Float64(8.7653)
	assert.Equal(t, float64(8.7653), *f64)

	i := Int(81)
	assert.Equal(t, 81, *i)

	s := String("Hello")
	assert.Equal(t, "Hello", *s)
}
