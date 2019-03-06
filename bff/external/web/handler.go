package web

import (
	"context"
	"encoding/json"
	transport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"go-bff/bff/adapter/transports/transport_email"
	"go-bff/bff/adapter/transports/transport_user"
	"go-bff/bff/registry/interfaces"
	"net/http"
)

func Handle(u interfaces.UseCases, m interfaces.MicroServices, ctx context.Context) http.Handler {
	router := mux.NewRouter()

	uu := u.NewUserUseCase()
	router.Methods("GET").Path("/users").Handler(transport.NewServer(
		transport_user.MakeFindEndpoint(uu),
		defaultDecodeRequest,
		encodeResponse,
	))
	router.Methods("POST").Path("/users").Handler(transport.NewServer(
		transport_user.MakeCreateEndpoint(uu),
		transport_user.DecodeCreateRequest,
		encodeResponse,
	))
	router.Methods("GET").Path("/users/{uid}").Handler(transport.NewServer(
		transport_user.MakeFirstEndpoint(uu),
		transport_user.DecodeFirstRequest,
		encodeResponse,
	))
	router.Methods("DELETE").Path("/users/{uid}").Handler(transport.NewServer(
		transport_user.MakeDeleteEndpoint(uu),
		transport_user.DecodeDeleteRequest,
		encodeResponse,
	))
	router.Methods("POST").Path("/users/{uid}/emails").Handler(transport.NewServer(
		transport_user.MakeCreateEmailEndpoint(uu),
		transport_user.DecodeCreateEmailRequest,
		encodeResponse,
	))

	eu := u.NewEmailUseCase()
	es := m.GetEmailService()
	router.Methods("GET").Path("/emails/{id}").Handler(transport.NewServer(
		transport_email.MakeFirstEndpoint(eu, es, ctx),
		transport_email.DecodeFirstRequest,
		encodeResponse,
	))
	router.Methods("PATCH").Path("/emails").Handler(transport.NewServer(
		transport_email.MakeUpdateEndpoint(eu, es, ctx),
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
