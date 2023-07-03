package exercises

import (
	"testing"
)

func TestSetAlarm(t *testing.T) {
	tests := []struct {
		name     string
		employed bool
		vacation bool
		want     bool
	}{
		{name: "Employed and on vacation", employed: true, vacation: true, want: false},
		{name: "Employed and not on vacation", employed: true, vacation: false, want: true},
		{name: "Not employed and on vacation", employed: false, vacation: true, want: false},
		{name: "Not employed and not on vacation", employed: false, vacation: false, want: false},
	}

	for _, tc := range tests {
		got := SetAlarm(tc.employed, tc.vacation)
		if got != tc.want {
			t.Errorf("%s: expected %v, got %v", tc.name, tc.want, got)
		}
	}
}
