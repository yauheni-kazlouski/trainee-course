package hello

import(
	"testing"
)

func TestHello(t *testing.T) {
	expected := "Hello, world"
	got := Hello()

	if expected != got {
		t.Errorf("expected: %s, got: %s", expected, got)
	}
}