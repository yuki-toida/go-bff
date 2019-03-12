package controller_email

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
	"go-bff/bff/application/usecase/usecase_email"
	"net/http"
	"strconv"
)

func MakeFirstEndpoint(u usecase_email.UseCase) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(string)
		emailID, _ := strconv.Atoi(req)
		email, err := u.First(uint64(emailID))
		if err != nil {
			return nil, err
		}
		return email, nil
	}
}

func DecodeFirstRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	emailID, ok := vars["id"]
	if !ok {
		return nil, errors.New("id parse error")
	}
	return emailID, nil
}

func MakeUpdateEndpoint(u usecase_email.UseCase) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(updateRequest)
		emailID, _ := strconv.Atoi(req.ID)
		email, err := u.Update(uint64(emailID), req.Email)
		if err != nil {
			return nil, err
		}
		return email, nil
	}
}

func DecodeUpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request updateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

type updateRequest struct {
	ID    string
	Email string
}
