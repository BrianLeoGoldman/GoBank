package main

import "net/http"

// Here is the API server and handlers

// APIServer represents an API server and contains information and functionality related to that server
type APIServer struct {
	// Stores the address at which the server will listen for incoming requests (ex: 127.0.0.1:8080)
	listenAddress string
	// TODO: add PostgreSQL database
}

// NewAPIServer returns a pointer to an APIServer.
// Convention in Go for functions that create and return pointers to structures is to use New followed by the name
// of the structure (in this case, New + APIServer)
func NewAPIServer(listenAddress string) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
	}
}

// handleAccount is a method, which means it is a function associated with a specific type, in this case, APIServer
func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
