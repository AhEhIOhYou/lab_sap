package triangle_test

import (
	"math"
	"testing"
	"triangle"
)

func TestPointInTriangle(t *testing.T) {
	v1 := triangle.Point{-10, -10}
	v2 := triangle.Point{0, 10}
	v3 := triangle.Point{10, 0}
	tri := triangle.Triangle{v1, v2, v3}

	pt1 := triangle.Point{2, 2}
	if !triangle.PointInTriangle(pt1, tri.V1, tri.V2, tri.V3) {
		t.Errorf("Point %v should be inside triangle %v", pt1, tri)
	}

	pt2 := triangle.Point{20, 20}
	if triangle.PointInTriangle(pt2, tri.V1, tri.V2, tri.V3) {
		t.Errorf("Point %v should be outside triangle %v", pt2, tri)
	}
}

func TestTrianglePerimeter(t *testing.T) {
	tri := triangle.Triangle{
		V1: triangle.Point{X: 0, Y: 0},
		V2: triangle.Point{X: 0, Y: 10},
		V3: triangle.Point{X: 10, Y: 0},
	}

	expectedPerimeter := 34.0
	actualPerimeter := math.Round(tri.Perimeter())

	if actualPerimeter != expectedPerimeter {
		t.Errorf("Expected perimeter %f, but got %f", expectedPerimeter, actualPerimeter)
	}
}

func TestTriangleArea(t *testing.T) {
	tri := triangle.Triangle{
		V1: triangle.Point{X: 0, Y: 0},
		V2: triangle.Point{X: 0, Y: 10},
		V3: triangle.Point{X: 10, Y: 0},
	}

	expectedArea := 50.0
	actualArea := tri.Area()

	if actualArea != expectedArea {
		t.Errorf("Expected area %f, but got %f", expectedArea, actualArea)
	}
}
