package gu

import (
	"fmt"
	"strings"
)

type ExampleA struct {
	ID int32
	S  []string
}

func (a ExampleA) Transform() ExampleB {
	return ExampleB{
		ID: int64(a.ID),
		S:  strings.Join(a.S, ", "),
	}
}

type ExampleB struct {
	ID int64
	S  string
}

func ExampleTransform() {
	in := []ExampleA{
		{
			ID: 1,
			S:  []string{"Hello", "World!"},
		},
		{
			ID: 2,
			S:  []string{"foo", "bar"},
		},
		{
			ID: 3,
			S:  []string{"spanac"},
		},
	}

	out := Transform[ExampleB](in)

	fmt.Printf("out is of type %T and contains %v", out, out)
	// Output: out is of type []gu.ExampleB and contains [{1 Hello, World!} {2 foo, bar} {3 spanac}]
}
