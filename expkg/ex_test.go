package expkg

import "testing"

func Test_hello(t *testing.T) {
	want := "Hello, world"
	if got := bello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func Test_bello2(t *testing.T) {
	want := "Hello, world"
	if got := bello2(); got != want {
		t.Errorf("Hello2() = %q, want %q", got, want)
	}
}
