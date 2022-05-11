package geometry

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, item := range areaTests {
		t.Run("when running for: "+reflect.TypeOf(item.shape).Name(), func(t *testing.T) {
			got := item.shape.Area()
			assert.Equal(t, item.want, got)
		})
	}
}
