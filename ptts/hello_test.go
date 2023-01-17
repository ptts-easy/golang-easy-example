package ptts

import "testing"

func TestHello_0(t *testing.T) {
	want := "Hello, world."
	if got := Hello_0(); got != want {
		t.Errorf("Hello_0() = %q, want %q", got, want)
	}
}

func TestHello_1(t *testing.T) {
	want := "Hello, world."
	if got := Hello_1(); got != want {
		t.Errorf("Hello_1() = %q, want %q", got, want)
	}
}
