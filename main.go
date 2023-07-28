package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jbuget.fr/explore-golang/hello"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {

	hello.Ciao()

	databaseUrl := os.Getenv("DATABASE_URL")
	log.Println(databaseUrl)

	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT COUNT(*) FROM accounts")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	rows.Next()
	var count string
	rows.Scan(&count)
	log.Println(count)

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

	r.Post("/accounts/signup", func(w http.ResponseWriter, r *http.Request) {
		log.Panicln("Not yet implemented `POST /accounts/signup`")
	})

	r.Post("/accounts/signin", func(w http.ResponseWriter, r *http.Request) {
		log.Panicln("Not yet implemented `POST /accounts/signin`")
	})

	r.Get("/accounts/me", func(w http.ResponseWriter, r *http.Request) {
		log.Panicln("Not yet implemented `GET /accounts/me`")
	})

	r.Delete("/accounts/me", func(w http.ResponseWriter, r *http.Request) {
		log.Panicln("Not yet implemented `DELETE /accounts/me`")
	})

	log.Println("Server is up and listening on http://localhost…")

	http.ListenAndServe(":80", r)
}
