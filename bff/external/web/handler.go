package web

import (
	"context"
	"encoding/json"
	transport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"go-bff/bff/adapter/transports/transport_email"
	"go-bff/bff/adapter/transports/transport_user"
	"go-bff/bff/registry"
	"net/http"
)

func Handle(r *registry.Registry) http.Handler {
	router := mux.NewRouter()

	router.Methods("GET").Path("/users").Handler(transport.NewServer(
		transport_user.MakeFindEndpoint(r.UseCases.User),
		defaultDecodeRequest,
		encodeResponse,
	))
	router.Methods("POST").Path("/users").Handler(transport.NewServer(
		transport_user.MakeCreateEndpoint(r.UseCases.User),
		transport_user.DecodeCreateRequest,
		encodeResponse,
	))
	router.Methods("GET").Path("/users/{uid}").Handler(transport.NewServer(
		transport_user.MakeFirstEndpoint(r.UseCases.User),
		transport_user.DecodeFirstRequest,
		encodeResponse,
	))
	router.Methods("DELETE").Path("/users/{uid}").Handler(transport.NewServer(
		transport_user.MakeDeleteEndpoint(r.UseCases.User),
		transport_user.DecodeDeleteRequest,
		encodeResponse,
	))
	router.Methods("POST").Path("/users/{uid}/emails").Handler(transport.NewServer(
		transport_user.MakeCreateEmailEndpoint(r.UseCases.User),
		transport_user.DecodeCreateEmailRequest,
		encodeResponse,
	))

	router.Methods("GET").Path("/emails/{id}").Handler(transport.NewServer(
		transport_email.MakeFirstEndpoint(r.UseCases.Email, r.MicroServices.Email, r.Context),
		transport_email.DecodeFirstRequest,
		encodeResponse,
	))
	router.Methods("PATCH").Path("/emails").Handler(transport.NewServer(
		transport_email.MakeUpdateEndpoint(r.UseCases.Email, r.MicroServices.Email, r.Context),
		transport_email.DecodeUpdateRequest,
		encodeResponse,
	))

	return router
}

func defaultDecodeRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return r, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
