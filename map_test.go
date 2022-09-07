package gu

import (
	"testing"
)

func TestMapEqual(t *testing.T) {
	tests := []struct {
		name string
		a    map[int]int
		b    map[int]int
		want bool
	}{
		{
			"both nil",
			nil,
			nil,
			true,
		},
		{
			"a nil",
			nil,
			map[int]int{
				1: 2,
				3: 4,
			},
			false,
		},
		{
			"b nil",
			map[int]int{
				1: 2,
				3: 4,
			},
			nil,
			false,
		},
		{
			"different length",
			map[int]int{
				1: 2,
				3: 4,
			},
			map[int]int{
				1: 2,
				3: 4,
				5: 6,
			},
			false,
		},
		{
			"different content",
			map[int]int{
				1: 2,
				3: 4,
			},
			map[int]int{
				3: 4,
				5: 6,
			},
			false,
		},
		{
			"equal",
			map[int]int{
				5: 6,
				1: 2,
				3: 4,
			},
			map[int]int{
				1: 2,
				3: 4,
				5: 6,
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapEqual(tt.a, tt.b); got != tt.want {
				t.Errorf("MapEqual = %v, want %v", got, tt.want)
			}
		})
	}
}

