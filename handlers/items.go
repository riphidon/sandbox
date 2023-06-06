package handlers

import (
	"net/http"
	"sandbox-api/services"
	"strconv"

	"github.com/gorilla/mux"
)

type IItems interface {
	Get(w http.ResponseWriter, r *http.Request)
	ById(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Register(*mux.Router)
}

type items struct {
	is services.IItemService
}

func newItems(s services.IItemService) IItems {
	return &items{
		is: s,
	}
}

func (i *items) Register(r *mux.Router) {
	r.HandleFunc("/items", i.Get).Methods("GET")
	r.HandleFunc("/items", i.Create).Methods("POST")
	r.HandleFunc("/items/{id}", i.ById).Methods("GET")
	r.HandleFunc("/items/{id}/update", i.Update).Methods("POST")
	r.HandleFunc("/items/{id}/delete", i.Delete).Methods("POST")
}

// Get GET "api/items"
func (i *items) Get(w http.ResponseWriter, r *http.Request) {
	items, err := i.is.Get()
	if err != nil {
		http.Error(w, "Something went wrong, please try again later", http.StatusInternalServerError)
		return
	}
	err = httpResponse(w, r, items, http.StatusOK)
	if err != nil {
		logger.Debugf("Could not process request : %v", err)
		http.Error(w, internalErr, http.StatusInternalServerError)
	}
}

// ById GET "api/items/${id}"
func (i *items) ById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.Debugf("[HANDLER] ID Parsing Error : %v", err)
		http.Error(w, "Error : Parameter 'id' is invalid.", http.StatusBadRequest)
		return
	}
	item, err := i.is.ById(id)
	if err != nil {
		http.Error(w, "Something went wrong, please try again later", http.StatusInternalServerError)
		return
	}
	if err = httpResponse(w, r, item, http.StatusOK); err != nil {
		logger.Debugf("Could not process request : %v", err)
		http.Error(w, internalErr, http.StatusInternalServerError)
	}
}

// Create POST "api/items"
func (i *items) Create(w http.ResponseWriter, r *http.Request) {
	new := i.is.NewItem()
	if err := parseForm(new, r); err != nil {
		logger.Debugf("[HANDLER] Body Parse Error: %v", err)
		http.Error(w, "Form Error : "+err.Error(), http.StatusBadRequest)
		return
	}

	id, err := i.is.Create(*new)
	if err != nil {
		http.Error(w, internalErr, http.StatusInternalServerError)
		return
	}

	if err = httpResponse(w, r, id, http.StatusCreated); err != nil {
		logger.Debugf(" [HANDLER] Could not process request : %v", err)
		http.Error(w, internalErr, http.StatusInternalServerError)
	}
}

// Update POST "api/items/${id}/update"
func (i *items) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.Debugf("[HANDLER] ID Parsing Error : %v", err)
		http.Error(w, "Error : Parameter 'id' is invalid.", http.StatusBadRequest)
		return
	}
	fields := i.is.NewItem()

	if err = parseForm(fields, r); err != nil {
		logger.Debugf("[HANDLER] Body Parse Error: %v", err)
		http.Error(w, "Form Error : "+err.Error(), http.StatusBadRequest)
		return
	}

	if err = i.is.Update(fields, id); err != nil {
		http.Error(w, internalErr, http.StatusInternalServerError)
		return
	}

	err = httpResponse(w, r, nil, http.StatusOK)
	if err != nil {
		logger.Debugf("Could not process request : %v", err)
		http.Error(w, internalErr, http.StatusInternalServerError)
	}

}

// Delete POST "api/items/${id}/delete"
func (i *items) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.Debugf("[HANDLER] ID Parsing Error : %v", err)
		http.Error(w, "Error : Parameter 'id' is invalid.", http.StatusBadRequest)
		return
	}
	if err := i.is.Delete(id); err != nil {
		http.Error(w, internalErr, http.StatusInternalServerError)
	}

	if err := httpResponse(w, r, nil, http.StatusOK); err != nil {
		logger.Debugf("Could not process request : %v", err)
		http.Error(w, internalErr, http.StatusInternalServerError)
	}

}
