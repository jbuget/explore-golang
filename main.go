package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/jbuget.fr/explore-golang/internal"
	"github.com/jbuget.fr/explore-golang/internal/core/services"
	"github.com/jbuget.fr/explore-golang/internal/entrypoints/web"
	"github.com/jbuget.fr/explore-golang/internal/infrastructure/repositories/accounts"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {

	// gracefully exit on keyboard interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	databaseUrl := os.Getenv("DATABASE_URL")
	db, err := internal.Connect(databaseUrl)
	if err != nil {
		log.Printf("error: %v\n", err)
		os.Exit(1)
	}
	log.Println("Database connected")

	accountsRepository := accounts.NewAccountsRepositoryPostgres(db)
	accountsService := services.NewAccountsService(accountsRepository)
	tokenAuth := jwtauth.New("HS256", []byte("SecretYouShouldHide"), nil)
	handlers := web.NewHTTPController(accountsService, tokenAuth)

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.Compress(5))
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

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

	// Protected routes
	r.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(tokenAuth))

		// Handle valid / invalid tokens. In this example, we use
		// the provided authenticator middleware, but you can write your
		// own very easily, look at the Authenticator method in jwtauth.go
		// and tweak it, its not scary.
		r.Use(jwtauth.Authenticator)

		// curl http://localhost/accounts -H "Authorization: Bearer {token}"
		r.Get("/accounts", handlers.ListAccounts)

		// curl http://localhost/accounts/{id} -H "Authorization: Bearer {token}"
		r.Get("/accounts/{accountId}", handlers.GetAccount)

		// curl -X PATCH http://localhost/accounts/6 -H "Authorization: Bearer $TOKEN" -d '{"name": "Another name"}'
		r.Patch("/accounts/{accountId}", handlers.UpdateAccount)

		// curl -X DELETE http://localhost/accounts/{id} -H "Authorization: Bearer {token}"
		r.Delete("/accounts/{accountId}", handlers.DeleteAccount)

		// curl http://localhost/admin -H "Authorization: Bearer {token}"
		r.Get("/admin", handlers.GetAdmin)
	})

	// Public routes
	r.Group(func(r chi.Router) {
		r.Get("/", handlers.GetRoot)

		// curl -v -X POST http://localhost/accounts -d '{"name":"Loulou","email":"loulou@example.org","password":"Abcd1234"}'
		r.Post("/accounts", handlers.CreateAccount)

		// curl -v -X POST http://localhost/token -d "email=tonton@example.org&password=Abcd1234"
		r.Post("/token", handlers.GetAccessToken)
	})

	if os.Getenv("PANIC") == "true" {
		panic("this is crashing")
	}

	port := "80"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	go func() {
		if err := http.ListenAndServe(":"+port, r); err != nil {
			log.Fatal("failed to start server", err)
			os.Exit(1)
		}
	}()

	log.Println("ready to serve requests on " + "0.0.0.0:" + port)
	<-c
	log.Println("gracefully shutting down")
	os.Exit(0)
}
