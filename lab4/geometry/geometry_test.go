package geometry_test

import (
	"geometry"
	"testing"
)

func TestRectangle(t *testing.T) {
	r := geometry.Rectangle{Width: 3, Height: 4}
	area := r.Area()
	expectedArea := 12.0
	if area != expectedArea {
		t.Errorf("Expected area to be %f, but got %f", expectedArea, area)
	}

	perimeter := r.Perimeter()
	expectedPerimeter := 14.0
	if perimeter != expectedPerimeter {
		t.Errorf("Expected perimeter to be %f, but got %f", expectedPerimeter, perimeter)
	}
}

func TestCircle(t *testing.T) {
	c := geometry.Circle{Radius: 5}
	area := c.Area()
	expectedArea := 78.53981633974483
	if area != expectedArea {
		t.Errorf("Expected area to be %f, but got %f", expectedArea, area)
	}

	perimeter := c.Perimeter()
	expectedPerimeter := 31.41592653589793
	if perimeter != expectedPerimeter {
		t.Errorf("Expected perimeter to be %f, but got %f", expectedPerimeter, perimeter)
	}
}
