package exercises

import (
	"reflect"
	"testing"
)

func TestDoubleArray(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{name: "Double int[]{1, 2, 3}", input: []int{1, 2, 3}, want: []int{2, 4, 6}},
		{name: "Double []int{10, 20, 30}", input: []int{10, 20, 30}, want: []int{20, 40, 60}},
		{name: "Double []int{50, 100, 200}", input: []int{50, 100, 200}, want: []int{100, 200, 400}},
	}

	for _, tc := range tests {
		got := DoubleArray(tc.input)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("%s: expected %v, got %v", tc.name, tc.want, got)
		}
	}
}

func TestAlternateDoubleArray(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{name: "Double int[]{1, 2, 3}", input: []int{1, 2, 3}, want: []int{2, 4, 6}},
		{name: "Double []int{10, 20, 30}", input: []int{10, 20, 30}, want: []int{20, 40, 60}},
		{name: "Double []int{50, 100, 200}", input: []int{50, 100, 200}, want: []int{100, 200, 400}},
	}

	for _, tc := range tests {
		got := AlternateDoubleArray(tc.input)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("%s: expected %v, got %v", tc.name, tc.want, got)
		}
	}
}
