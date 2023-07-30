package users

import (
	"log"
	"os"
	"testing"

	"github.com/jbuget.fr/explore-golang/database"
	_ "github.com/lib/pq"
)

func TestFindAccounts(t *testing.T) {
	databaseUrl := "postgres://test:test@localhost:15434/test?sslmode=disable"
	db, err := database.Connect(databaseUrl)
	if err != nil {
		log.Printf("error: %v\n", err)
		os.Exit(1)
	}
	log.Println("Database connected")

	accountRepository := AccountRepository{DB: db}

	// Must find same accounts number
	if got := accountRepository.FindAccounts(); len(got) != 5 {
		t.Errorf("FindAccounts() = %v, want %v", len(got), 5)
	}
}
