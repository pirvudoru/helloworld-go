package geometry

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRectanglePerimeter(t *testing.T) {
	r := Rectangle{10.0, 10.0}
	got := r.Perimeter()
	want := 40.0

	assert.Equal(t, want, got)
}

func TestRectangleArea(t *testing.T) {
	r := Rectangle{10.0, 5.0}
	got := r.Area()
	want := 50.0

	assert.Equal(t, want, got)
}
