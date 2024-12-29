package namaste

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := SayNamaste()
	want := "Namaste"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
