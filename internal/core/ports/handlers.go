package ports

import "net/http"

type HTTPHandler interface {
	GetRoot(w http.ResponseWriter, r *http.Request)
	GetAdmin(w http.ResponseWriter, r *http.Request)
	CreateAccount(w http.ResponseWriter, r *http.Request)
	GetAccessToken(w http.ResponseWriter, r *http.Request)
	GetAccount(w http.ResponseWriter, r *http.Request)
	ListAccounts(w http.ResponseWriter, r *http.Request)
	DeleteAccount(w http.ResponseWriter, r *http.Request)
}
