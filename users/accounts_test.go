package users

import (
	"testing"

	"github.com/jbuget.fr/explore-golang/database"
	_ "github.com/lib/pq"
)

func TestFindAccounts(t *testing.T) {
	db := database.GetTestingDB()
	accountRepository := AccountRepository{DB: db}

	// Must find same accounts number
	if got := accountRepository.FindAccounts(); len(got) != 5 {
		t.Errorf("FindAccounts() = %v, want %v", len(got), 5)
	}
}
