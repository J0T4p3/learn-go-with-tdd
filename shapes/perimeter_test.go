package shapes

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("Positive perimeter calculation", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		want := 200.0

		if got != want {
			t.Errorf("Got %.2f want %.2f", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{
			name:  "Rectangle()",
			shape: Rectangle{12, 6},
			want:  72.0,
		},
		{
			name:  "Triangle()",
			shape: Triangle{12, 6},
			want:  36.0,
		},
		{
			name:  "Circle()",
			shape: Circle{10},
			want:  314.1592653589793,
		},
	}

	for _, tt := range areaTests {

		// Use the name to properly identify the tests
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			// Test with tolerance when using floats
			if math.Abs(got-tt.want) > 0.0001 {
				t.Errorf(
					"%#v: got %g want %g",
					tt.shape,
					got,
					tt.want,
				)
			}
		})
	}
}
