package hello

import "testing"

func TestCiao(t *testing.T) {
	want := "Ciao, tutti."
	if got := Ciao(); got != want {
		t.Errorf("Ciao() = %q, want %q", got, want)
	}
}
