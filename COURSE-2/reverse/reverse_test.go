package reverse

import (
	"testing"
)

var tests = map[string]struct {
	input string
	want  string
}{
	"common":       {input: "reverse this string", want: "esrever siht gnirts"},
	"one word":     {input: "oneLongWord", want: "droWgnoLeno"},
	"empty":        {input: "", want: ""},
	"only spaces":  {input: "        ", want: "        "},
	"extra spaces": {input: "  reverse  this string    ", want: "esrever siht gnirts"},
}

func TestClearReverse(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := ClearReverse(test.input)
			if got != test.want {
				t.Fatalf("input \"%s\" expected: \"%v\", got: \"%v\"", test.input, test.want, got)
			}
		})
	}
}

func TestPoweredReverse(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := PoweredReverse(test.input)
			if got != test.want {
				t.Fatalf("input \"%s\" expected: \"%v\", got: \"%v\"", test.input, test.want, got)
			}
		})
	}
}
