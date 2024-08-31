//nolint:testpackage // Ignore this linter rule because we are not using the _test suffix for the package name.
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
