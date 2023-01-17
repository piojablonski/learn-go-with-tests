package main

import "testing"

func TestPerimeter(t *testing.T) {
	want := 60.0
	rectangle := Rectangle{10.0, 20.0}
	got := rectangle.Perimeter()

	if got != want {
		t.Errorf("got: %.2f, want: %.2f", got, want)
	}
}

type Shape interface {
	Area() float64
}

func TestArea(t *testing.T) {

	areaTest := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{10.0, 20.0}, want: 20.0},
		{shape: Circle{10.0}, want: 314.1592653589793},
		{shape: Triangle{12, 6}, want: 36.0},
	}

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("%#v got: %.2f, want: %.2f", shape, got, want)
		}
	}

	for _, tt := range areaTest {
		checkArea(t, tt.shape, tt.want)
	}
}
