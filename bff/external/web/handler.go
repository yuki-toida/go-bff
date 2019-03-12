package web

import (
	"context"
	"encoding/json"
	transport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"go-bff/bff/adapter/controllers/controller_email"
	"go-bff/bff/adapter/controllers/controller_user"
	"go-bff/bff/registry"
	"net/http"
)

func Handle(r *registry.Registry) http.Handler {
	router := mux.NewRouter()

	router.Methods("GET").Path("/users").Handler(transport.NewServer(
		controller_user.MakeFindEndpoint(r.UserUseCase),
		defaultDecodeRequest,
		encodeResponse,
	))
	router.Methods("POST").Path("/users").Handler(transport.NewServer(
		controller_user.MakeCreateEndpoint(r.UserUseCase),
		controller_user.DecodeCreateRequest,
		encodeResponse,
	))
	router.Methods("GET").Path("/users/{uid}").Handler(transport.NewServer(
		controller_user.MakeFirstEndpoint(r.UserUseCase),
		controller_user.DecodeFirstRequest,
		encodeResponse,
	))
	router.Methods("DELETE").Path("/users/{uid}").Handler(transport.NewServer(
		controller_user.MakeDeleteEndpoint(r.UserUseCase),
		controller_user.DecodeDeleteRequest,
		encodeResponse,
	))
	router.Methods("POST").Path("/users/{uid}/emails").Handler(transport.NewServer(
		controller_user.MakeCreateEmailEndpoint(r.UserUseCase),
		controller_user.DecodeCreateEmailRequest,
		encodeResponse,
	))

	router.Methods("GET").Path("/emails/{id}").Handler(transport.NewServer(
		controller_email.MakeFirstEndpoint(r.EmailUseCase),
		controller_email.DecodeFirstRequest,
		encodeResponse,
	))
	router.Methods("PATCH").Path("/emails").Handler(transport.NewServer(
		controller_email.MakeUpdateEndpoint(r.EmailUseCase),
		controller_email.DecodeUpdateRequest,
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
