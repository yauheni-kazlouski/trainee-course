package set

import (
	"slices"
	"testing"
)

func TestUniqueStrings(t *testing.T) {
	var tests = map[string]struct{
		input	[]string
		want 	[]string
	} {
		"common": {input: []string{"str-1", "str-2", "str-3", "str-3", "str-2", "str-4"}, want:[]string{"str-1", "str-2", "str-3", "str-4"}},
		"no repeats": {input: []string{"str-1", "str-2", "str-3"}, want:[]string{"str-1", "str-2", "str-3"}},
		"all repeats": {input: []string{"str", "str", "str"}, want:[]string{"str"}},
		"empty": {input: []string{}, want: []string{}},
	}

	for name, test := range(tests) {
		t.Run(name, func(t *testing.T){
			got := UniqueStrings(test.input)
			if !slices.Equal(got, test.want) {
				t.Fatalf("input %#v, expected: %#v, got: %#v", test.input, test.want, got)
			}
		})
	}
}