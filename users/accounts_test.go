package users

import "testing"

func TestFindAccounts(t *testing.T) {
	var want []Account = []Account{}
	if got := FindAccounts(); len(got) != len(want) {
		t.Errorf("FindAccounts() = %v, want %v", len(got), len(want))
	}
}
