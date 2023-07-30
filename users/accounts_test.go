package users

import (
	"fmt"
	"testing"
	"time"

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

func TestInsertAccount(t *testing.T) {
	db := database.GetTestingDB()

	db.Client.Exec("DELETE FROM accounts WHERE name=$1", "jeremy")

	accountRepository := AccountRepository{DB: db}

	newAccount := CreateAccount("jeremy", "jeremy@example.org", "Abcd1234")
	id := accountRepository.InsertAccount(newAccount)

	if got := accountRepository.FindAccounts(); len(got) != 6 {
		t.Errorf("FindAccounts() = %v, want %v", len(got), 6)
	}

	type AccountRow struct {
		Id        int
		CreatedAt time.Time
		UpdatedAt time.Time
		Name      string
		Email     string
		Password  string
		Enabled   bool
	}

	accountInDB := AccountRow{}
	row := db.Client.QueryRow("SELECT id, created_at, name, email, password, enabled  FROM accounts WHERE id=$1", id)
	err := row.Scan(
		&accountInDB.Id,
		&accountInDB.CreatedAt,
		&accountInDB.Name,
		&accountInDB.Email,
		&accountInDB.Password,
		&accountInDB.Enabled,
	)
	if err != nil {
		fmt.Println("An error occurred")
	}

	if accountInDB.Id != id {
		t.Errorf("Id = %v, want %v", accountInDB.Id, id)
	}
	if accountInDB.Name != "jeremy" {
		t.Errorf("Name = %v, want %v", accountInDB.Name, "jeremy")
	}
	if accountInDB.Email != "jeremy@example.org" {
		t.Errorf("Email = %v, want %v", accountInDB.Email, "jeremy@example.org")
	}
	if accountInDB.Password != "Abcd1234" {
		t.Errorf("Password = %v, want %v", accountInDB.Password, "Abcd1234")
	}
	if accountInDB.Enabled != true {
		t.Errorf("Enabled = %v, want %v", accountInDB.Enabled, true)
	}

	db.Client.Exec("DELETE FROM accounts WHERE id=$1", id)
}
