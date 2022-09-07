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

func TestMapCopy(t *testing.T) {
	tests := []struct {
		name string
		src  map[int]int
	}{
		{
			"nil",
			nil,
		},
		{
			"empty",
			map[int]int{},
		},
		{
			"values",
			map[int]int{
				1: 2,
				3: 4,
				5: 6,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if dst := MapCopy(tt.src); !MapEqual(dst, tt.src) {
				t.Errorf("MapCopy =\n%v\nwant\n%v", dst, tt.src)
			}
		})
	}
}

func TestMapCopyKeys(t *testing.T) {
	tests := []struct {
		name string
		src  map[int]int
		keys []int
		want map[int]int
	}{
		{
			"nil",
			nil,
			nil,
			nil,
		},
		{
			"all empty",
			map[int]int{},
			[]int{},
			map[int]int{},
		},
		{
			"keys nil",
			map[int]int{
				1: 2,
				3: 4,
				5: 6,
			},
			nil,
			map[int]int{},
		},
		{
			"keys empty",
			map[int]int{
				1: 2,
				3: 4,
				5: 6,
			},
			[]int{},
			map[int]int{},
		},
		{
			"subset",
			map[int]int{
				1: 2,
				3: 4,
				5: 6,
			},
			[]int{1, 5},
			map[int]int{
				1: 2,
				5: 6,
			},
		},
		{
			"superset",
			map[int]int{
				1: 2,
				3: 4,
				5: 6,
			},
			[]int{1, 5, 7, 9},
			map[int]int{
				1: 2,
				5: 6,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapCopyKeys(tt.src, tt.keys...); !MapEqual(got, tt.want) {
				t.Errorf("MapCopyKeys =\n%v\nwant\n%v", got, tt.want)
			}
		})
	}
}
