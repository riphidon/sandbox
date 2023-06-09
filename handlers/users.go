package handlers

import (
	"context"
	"net/http"
	"sandbox-api/services"

	"github.com/gorilla/mux"
)

type users struct {
	us services.IUserService
}

// accessCheck implements IUsers.
func (u *users) accessCheck(ctx context.Context) bool {
	return u.us.UserCheck(ctx)
}

// authCheck implements IUsers.
func (u *users) authCheck(token string, ctx context.Context) (context.Context, error) {
	c, err := u.us.ByRemember(token, ctx)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// create implements IUsers.
func (u *users) Create(w http.ResponseWriter, r *http.Request) {
	new := u.us.NewUser()
	if err := parseForm(new, r); err != nil {
		logger.Debugf("[HANDLER] Body Parse Error: %v", err)
		http.Error(w, "Form Error : "+err.Error(), http.StatusBadRequest)
		return
	}

	id, err := u.us.Create(*new)
	if err != nil {
		logger.Debugf("[HANDLER] User Creation Error: %v", err)
		http.Error(w, internalErr, http.StatusInternalServerError)
		return
	}

	if err = httpResponse(w, r, id, http.StatusCreated); err != nil {
		logger.Debugf(" [HANDLER] Could not process request : %v", err)
		http.Error(w, internalErr, http.StatusInternalServerError)
	}

}

// login implements IUsers.
func (u *users) login() {
	panic("unimplemented")
}

// logout implements IUsers.
func (u *users) logout() {
	panic("unimplemented")
}

func (u *users) Register(r *mux.Router) {
	r.HandleFunc("/signin/", u.Create).Methods("POST")
}

func (u *users) AuthHandler(next http.HandlerFunc) {
	u.AuthFn(next.ServeHTTP)
}

func (u *users) AuthFn(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("remember_token")
		if err != nil {
			logger.Debugf("[AUTH STAUS]: access denied: %v", err)
			http.Error(w, "You have not the rights to access the Api", http.StatusBadRequest)
			return
		}

		ctx, err := u.authCheck(cookie.Value, r.Context())
		if err != nil {
			logger.Debugf("[AUTH STAUS]: access denied: %v", err)
			http.Error(w, "You have not the rights to access the Api", http.StatusBadRequest)
			return
		}

		r = r.WithContext(ctx)

		loggedIn := u.accessCheck(r.Context())
		if !loggedIn {
			http.Error(w, "You are not logged in", http.StatusBadRequest)
			return
		}
		next(w, r)
	})
}

type IUsers interface {
	AuthHandler(next http.HandlerFunc)
	authCheck(token string, ctx context.Context) (context.Context, error)
	accessCheck(ctx context.Context) bool
	Create(w http.ResponseWriter, r *http.Request)
	login()
	logout()
	Register(r *mux.Router)
}

func NewUsers(us services.IUserService) IUsers {
	return &users{
		us: us,
	}
}
