package users

type Account struct {
	Id int
    Email string
    Name  string
    Password string
	Enabled bool
}

func FindAccounts() []Account {
	alice := Account{
		Id: 1,
		Email: "alice@example.com",
		Name: "Alice",
		Password: "abcd1234",
		Enabled: true,
	}
	bob := Account{
		Id: 2,
		Email: "bob@example.com",
		Name: "Bob",
		Password: "abcd1234",
		Enabled: true,
	}

	accounts := []Account{alice, bob}
	return accounts
}
