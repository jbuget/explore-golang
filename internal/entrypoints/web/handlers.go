package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/jbuget.fr/explore-golang/internal/core/domain"
	"github.com/jbuget.fr/explore-golang/internal/core/ports"
	"golang.org/x/crypto/bcrypt"
)

type HTTPController struct {
	accountsService ports.AccountsService
	tokenAuth       *jwtauth.JWTAuth
}

func NewHTTPController(accountsService ports.AccountsService, tokenAuth *jwtauth.JWTAuth) *HTTPController {
	return &HTTPController{accountsService: accountsService, tokenAuth: tokenAuth}
}

func (controller *HTTPController) GetRoot(w http.ResponseWriter, r *http.Request) {
	render.PlainText(w, r, "It works!")
}

func (controller *HTTPController) GetAdmin(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	render.PlainText(w, r, fmt.Sprintf("protected area. hi %v", claims["user_id"]))
}

func (controller *HTTPController) CreateAccount(w http.ResponseWriter, r *http.Request) {
	data := &AccountCreationRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	account := domain.NewAccountWithEncryptedPassword(data.Name, data.Email, data.Password)
	id := controller.accountsService.InsertAccount(account)

	resp := &AccountCreationResponse{
		Id:    id,
		Name:  account.Account.Name,
		Email: account.Account.Email,
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, resp)

}

func (controller *HTTPController) GetAccessToken(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	account := controller.accountsService.GetActiveAccountByEmail(email)
	err := bcrypt.CompareHashAndPassword([]byte(account.EncryptedPassword), []byte(password))
	if err == nil {
		_, tokenString, _ := controller.tokenAuth.Encode(map[string]interface{}{"user_id": account.Account.Id})
		response := map[string]interface{}{
			"token":      tokenString,
			"token_type": "jwt",
		}
		render.JSON(w, r, response)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("Content-Type", "application/json")
	}

}

func (controller *HTTPController) GetAccount(w http.ResponseWriter, r *http.Request) {
	accountId, _ := strconv.Atoi(chi.URLParam(r, "accountId"))
	_, claims, _ := jwtauth.FromContext(r.Context())
	userId := int(claims["user_id"].(float64))
	if accountId != userId {
		render.Render(w, r, ErrForbidden)
		return
	}
	account := controller.accountsService.GetAccountById(userId)
	render.JSON(w, r, account)
}

func (controller *HTTPController) ListAccounts(w http.ResponseWriter, r *http.Request) {
	accounts := controller.accountsService.FindAccounts()
	render.JSON(w, r, accounts)
}

func (controller *HTTPController) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	accountId, _ := strconv.Atoi(chi.URLParam(r, "accountId"))

	_, claims, _ := jwtauth.FromContext(r.Context())
	userId := int(claims["user_id"].(float64))
	if accountId != userId {
		render.Render(w, r, ErrForbidden)
		return
	}

	controller.accountsService.DeleteAccount(accountId)
	render.NoContent(w, r)
}

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

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
