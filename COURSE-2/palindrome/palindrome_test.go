package palindrome

import (
	"testing"
)

var tests = map[string]struct {
	word string
	want bool
}{
	"commmon_odd":  {word: "civic", want: true},
	"commmon_even": {word: "civvic", want: true},
	"numbers":      {word: "22/2/22", want: true},
	"false":        {word: "golang", want: false},
	"empty":        {word: "", want: false},
	"with spaces":  {word: "Mr. Owl ate my metal worm", want: true},
}

func TestHalfComparison(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(tt *testing.T) {
			got := HalfComparison(test.word)
			if test.want != got {
				tt.Errorf("word %s expected: %v, got: %v", test.word, test.want, got)
			}
		})
	}
}



func TestReverse(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(tt *testing.T) {
			got := Reverse(test.word)
			if test.want != got {
				tt.Errorf("word %s expected: %v, got: %v", test.word, test.want, got)
			}
		})
	}
}

func TestRecursive(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(tt *testing.T) {
			got := Recursive(test.word)
			if test.want != got {
				tt.Errorf("word %s expected: %v, got: %v", test.word, test.want, got)
			}
		})
	}
}


func BenchmarkHalfComparison(b *testing.B) {
	for name, test := range tests {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++{
				HalfComparison(test.word)
			}
		})
	}
}

func BenchmarkReverse(b *testing.B) {
	
	for name, test := range tests {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++{
				Reverse(test.word)
			}
		})
	}
}

func BenchmarkRecursive(b *testing.B) {
	
	for name, test := range tests {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++{
				Recursive(test.word)
			}
		})
	}
}