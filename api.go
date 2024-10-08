package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// Here is the API server and handlers

// WriteJSON sends JSON responses to an HTTP client, commonly used in web applications to respond with structured data
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// apiFunc is the function signature of the function we will be using
// we define a function type that accepts two parameters and returns a value of type error
type apiFunc func(w http.ResponseWriter, r *http.Request) error

type APIError struct {
	Error string
}

// makeHTTPHandleFunc decorates the APIFunc into a http.HandlerFunc
func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			// Handle the error
			WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

// APIServer represents an API server and contains information and functionality related to that server
type APIServer struct {
	listenAddress string
	storage       Storage
}

// NewAPIServer returns a pointer to an APIServer.
func NewAPIServer(listenAddress string, storage Storage) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
		storage:       storage,
	}
}

// Run starts the HTTP server, configures routes, and starts listening on the specified address
func (s *APIServer) Run() {
	// returns a new router that will be used to define and handle API routes
	router := mux.NewRouter()
	// defines an /account route on the router which is handled using the APIServer's handleAccount function
	// handleAccount is converted to a http.HandlerFunc using the makeHTTPHandleFunc function
	// this allows the handleAccount method to handle requests to the /account route
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleAccountByID))
	log.Println("JSON API server running on port ", s.listenAddress)
	http.ListenAndServe(s.listenAddress, router)
}

// handleAccount is a method, which means it is a function associated with a specific type, in this case, APIServer
func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	// With the mux router we cannot specify if the request is GET, POST, PUT or DELETE
	if r.Method == "GET" {
		return s.handleGetAccounts(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}
	return fmt.Errorf("method not allowed: %s", r.Method)
}

func (s *APIServer) handleAccountByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccountByID(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}
	return fmt.Errorf("method not allowed: %s", r.Method)
}

func (s *APIServer) handleGetAccounts(w http.ResponseWriter, r *http.Request) error {
	accounts, err := s.storage.GetAccounts()
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, accounts)
}

func (s *APIServer) handleGetAccountByID(w http.ResponseWriter, r *http.Request) error {
	idString := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println("ID argument is not valid:", idString)
		return err
	}
	fmt.Printf("\nGetting account with id %v", id)
	account, err := s.storage.GetAccountByID(id)
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	createAccountReq := new(CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(createAccountReq); err != nil {
		return err
	}
	account := NewAccount(createAccountReq.Firstname, createAccountReq.Lastname)
	savedAccount, err := s.storage.CreateAccount(account)
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusCreated, savedAccount)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	idString := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println("ID argument is not valid:", idString)
		return err
	}
	err = s.storage.DeleteAccount(id)
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, "Deleted account with id "+idString)
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
