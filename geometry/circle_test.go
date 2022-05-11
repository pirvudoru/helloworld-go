package geometry

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCircleArea(t *testing.T) {
	r := Circle{10.0}
	got := r.Area()
	want := 314.1592653589793

	assert.Equal(t, want, got)
}
