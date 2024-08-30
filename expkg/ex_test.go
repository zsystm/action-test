package expkg

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world"
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestHello2(t *testing.T) {
	want := "Hello, world"
	if got := Hello2(); got != want {
		t.Errorf("Hello2() = %q, want %q", got, want)
	}
}
