package structs

import "testing"

var rectangle = Rectangle{10.0, 10.0}
var circle = Circle{10}

func TestPerimeter(t *testing.T) {
	t.Run("Rectangle perimeter", func(t *testing.T) {

		expected := 40.0
		result := Perimeter(rectangle)
		if expected != result {
			t.Errorf("Expected %.2f but got %.2f", expected, result)
		}
	})
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, expected float64) {
		t.Helper()
		result := shape.Area()
		if expected != result {
			t.Errorf("Expected %g but got %g", expected, result)
		}
	}
	t.Run("Rectangle area", func(t *testing.T) {
		expected := 100.0
		checkArea(t, rectangle, expected)
	})

	t.Run("Circle area", func(t *testing.T) {
		expected := 314.1592653589793
		checkArea(t, circle, expected)
	})
}

// the same as previous test but using table driven! tests
func TestAreaAnonymousStruct(t *testing.T) {
	areaTest := []struct { //anonymous struct
		shape    Shape
		expected float64
	}{
		{shape: Rectangle{Width: 10, Height: 10}, expected: 100},
		{shape: Circle{Radius: 10}, expected: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, expected: 36.0},
	}
	for _, tt := range areaTest {
		result := tt.shape.Area()
		if result != tt.expected {
			t.Errorf("for %#v result %g want %g", tt.shape, result, tt.expected)
		}
	}
}
