package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jbuget.fr/explore-golang/database"
	"github.com/jbuget.fr/explore-golang/hello"
	"github.com/jbuget.fr/explore-golang/users"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

var accountRepository *users.AccountRepository

func main() {

	hello.Ciao()

	databaseUrl := os.Getenv("DATABASE_URL")
	db, err := database.Connect(databaseUrl)
	if err != nil {
		log.Printf("error: %v\n", err)
		os.Exit(1)
	}
	log.Println("Database connected")

	accountRepository = &users.AccountRepository{DB: db}

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.Compress(5))
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to my website!"))
	})

	r.Post("/decode", func(w http.ResponseWriter, r *http.Request) {
		var user User
		json.NewDecoder(r.Body).Decode(&user)

		fmt.Fprintf(w, "%s %s is %d years old!", user.Firstname, user.Lastname, user.Age)
	})

	r.Get("/encode", func(w http.ResponseWriter, r *http.Request) {
		peter := User{
			Firstname: "John",
			Lastname:  "Doe",
			Age:       25,
		}

		json.NewEncoder(w).Encode(peter)
	})

	r.Get("/pokemon/{id_or_name}", func(w http.ResponseWriter, r *http.Request) {
		pokemon_id_or_name := chi.URLParam(r, "id_or_name")

		url := "https://pokeapi.co/api/v2/pokemon/" + pokemon_id_or_name

		http_client := &http.Client{}

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
		}

		res, getErr := http_client.Do(req)
		if getErr != nil {
			log.Fatal(getErr)
		}

		if res.Body != nil {
			defer res.Body.Close()
		}

		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}

		fmt.Fprintf(w, "%s", body)
	})

	// curl -v -X POST http://localhost/accounts -d "name=tonton&email=tonton@example.org&password=Abcd1234"
	r.Post("/accounts", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		name := r.Form.Get("name")
		email := r.Form.Get("email")
		password := r.Form.Get("password")

		account := users.CreateAccount(name, email, password)
		id := accountRepository.InsertAccount(account)
		json.NewEncoder(w).Encode(id)
	})

	// curl -v -X POST http://localhost/accounts/authenticate -d "email=tonton@example.org&password=Abcd1234"
	r.Post("/accounts/authenticate", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		email := r.Form.Get("email")
		password := r.Form.Get("password")

		account := accountRepository.GetActiveAccountByEmail(email)
		err := bcrypt.CompareHashAndPassword([]byte(account.EncryptedPassword), []byte(password))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
		} else {
			json.NewEncoder(w).Encode(account)
		}
	})

	r.Get("/accounts", func(w http.ResponseWriter, r *http.Request) {
		accounts := accountRepository.FindAccounts()
		json.NewEncoder(w).Encode(accounts)
	})

	r.Get("/accounts/me", func(w http.ResponseWriter, r *http.Request) {
		account := accountRepository.GetAccount()
		json.NewEncoder(w).Encode(account)
	})

	r.Delete("/accounts/me", func(w http.ResponseWriter, r *http.Request) {
		log.Panicln("Not yet implemented `DELETE /accounts/me`")
	})

	log.Println("Server is up and listening on http://localhost…")

	http.ListenAndServe(":80", r)
}
