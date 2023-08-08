package main

import (
	"encoding/json"
	"fmt"
	"io"
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
	"github.com/jbuget.fr/explore-golang/database"
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

var accountRepository users.AccountRepository

var tokenAuth *jwtauth.JWTAuth

type AccountCreationRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *AccountCreationRequest) Bind(r *http.Request) error {
	return nil
}

type AccountCreationResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (acr *AccountCreationResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type BearerTokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (b *BearerTokenRequest) Bind(r *http.Request) error {
	return nil
}

type BearerTokenResponse struct {
	Token     string `json:"token"`
	TokenType string `json:"token_type"`
}

func main() {

	// gracefully exit on keyboard interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	databaseUrl := os.Getenv("DATABASE_URL")
	db, err := database.Connect(databaseUrl)
	if err != nil {
		log.Printf("error: %v\n", err)
		os.Exit(1)
	}
	log.Println("Database connected")

	accountRepository = users.AccountRepository{DB: db}

	tokenAuth = jwtauth.New("HS256", []byte("SecretYouShouldHide"), nil)

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
		r.Get("/accounts", func(w http.ResponseWriter, r *http.Request) {
			accounts := accountRepository.FindAccounts()
			json.NewEncoder(w).Encode(accounts)
		})

		r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context())
			w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["user_id"])))
		})
	})

	// Public routes
	r.Group(func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("welcome anonymous"))
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

			body, readErr := io.ReadAll(res.Body)
			if readErr != nil {
				log.Fatal(readErr)
			}

			fmt.Fprintf(w, "%s", body)
		})

		// curl -v -X POST http://localhost/accounts -d '{"name":"Loulou","email":"loulou@example.org","password":"Abcd1234"}'
		r.Post("/accounts", func(w http.ResponseWriter, r *http.Request) {
			data := &AccountCreationRequest{}
			if err := render.Bind(r, data); err != nil {
				render.Render(w, r, ErrInvalidRequest(err))
				return
			}

			account := users.CreateAccount(data.Name, data.Email, data.Password)
			id := accountRepository.InsertAccount(account)

			resp := &AccountCreationResponse{
				Id:    id,
				Name:  account.Account.Name,
				Email: account.Account.Email,
			}

			render.Status(r, http.StatusCreated)
			render.Render(w, r, resp)
		})

		// curl -v -X POST http://localhost/token -d "email=tonton@example.org&password=Abcd1234"
		r.Post("/token", func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			email := r.Form.Get("email")
			password := r.Form.Get("password")

			account := accountRepository.GetActiveAccountByEmail(email)
			err := bcrypt.CompareHashAndPassword([]byte(account.EncryptedPassword), []byte(password))
			if err == nil {

				_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": account.Account.Id})
				response := map[string]interface{}{
					"token":      tokenString,
					"token_type": "jwt",
				}
				json.NewEncoder(w).Encode(response)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				w.Header().Set("Content-Type", "application/json")
			}
		})
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

//--
// Error response payloads & renderers
//--

// ErrResponse renderer type for handling all sorts of errors.
//
// In the best case scenario, the excellent github.com/pkg/errors package
// helps reveal information on the error, setting it on Err, and in the Render()
// method, using it to set the application-specific error code in AppCode.
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}
var ErrForbidden = &ErrResponse{HTTPStatusCode: 403, StatusText: "Forbidden."}
