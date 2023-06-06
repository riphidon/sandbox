package handlers

import (
	"net/http"
	"sandbox-api/services"
	"strconv"

	"github.com/gorilla/mux"
)

const internalErr string = "Something went wrong. Please try again later."
const notEmpty string = "Field must not be empty."

type IMakers interface {
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Register(*mux.Router)
}
type Makers struct {
	ms services.IMakerService
}

func NewMakers(s services.IMakerService) IMakers {
	return &Makers{
		ms: s,
	}
}

func (m *Makers) Register(r *mux.Router) {
	r.HandleFunc("/makers", m.Get).Methods("GET")
	r.HandleFunc("/makers", m.Create).Methods("POST")
	r.HandleFunc("/makers/{id}/delete", m.Delete).Methods("POST")
	r.HandleFunc("/makers/{id}/update", m.Update).Methods("POST")
}

func (m *Makers) Get(w http.ResponseWriter, r *http.Request) {
	makers, err := m.ms.Get()
	if err != nil {
		http.Error(w, "Something went wrong, please try again later", http.StatusInternalServerError)
		return
	}
	if err = httpResponse(w, r, makers, http.StatusOK); err != nil {
		logger.Debugf("Could not process request : %v", err)
		http.Error(w, internalErr, http.StatusInternalServerError)
	}
}

func (m *Makers) Create(w http.ResponseWriter, r *http.Request) {
	new := m.ms.NewMaker()

	if err := parseForm(new, r); err != nil {
		logger.Debugf("[HANDLER] Body Parse Error: %v", err)
		http.Error(w, "Form Error : "+err.Error(), http.StatusBadRequest)
		return
	}

	if new.Name == "" {
		logger.Debugf("[HANDLER] Form 'name' : %v", notEmpty)
		http.Error(w, "Name field must not be empty", http.StatusBadRequest)
		return
	}
	id, err := m.ms.Create(*new)
	if err != nil {
		http.Error(w, internalErr, http.StatusInternalServerError)
		return
	}

	if err = httpResponse(w, r, id, http.StatusCreated); err != nil {
		logger.Debugf("Could not process request : %v", err)
		http.Error(w, internalErr, http.StatusInternalServerError)
	}
}

func (m *Makers) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.Debugf("[HANDLER] ID Parsing Error : %v", err)
		http.Error(w, "Error : Parameter 'id' is invalid.", http.StatusBadRequest)
		return
	}
	if err = m.ms.Delete(id); err != nil {
		http.Error(w, internalErr, http.StatusInternalServerError)
		return
	}

	if err = httpResponse(w, r, nil, http.StatusOK); err != nil {
		logger.Debugf("Could not process request : %v", err)
		http.Error(w, internalErr, http.StatusInternalServerError)
	}
}

func (m *Makers) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.Debugf("[HANDLER] ID Parsing Error : %v", err)
		http.Error(w, "Error : Parameter 'id' is invalid.", http.StatusBadRequest)
		return
	}
	update := m.ms.NewMaker()

	if err := parseForm(update, r); err != nil {
		logger.Debugf("[HANDLER] Body Parse Error: %v", err)
		http.Error(w, "Form Error : "+err.Error(), http.StatusBadRequest)
		return
	}

	if update.Name == "" {
		logger.Debugf("[HANDLER] Form 'name' : %v", notEmpty)
		http.Error(w, "Name field must not be empty", http.StatusBadRequest)
		return
	}
	err = m.ms.Update(*update, id)
	if err != nil {
		http.Error(w, internalErr, http.StatusInternalServerError)
		return
	}

	if err = httpResponse(w, r, id, http.StatusCreated); err != nil {
		logger.Debugf("Could not process request : %v", err)
		http.Error(w, internalErr, http.StatusInternalServerError)
	}
}
