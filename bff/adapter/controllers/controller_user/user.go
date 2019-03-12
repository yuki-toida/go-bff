package controller_user

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
	"go-bff/bff/application/usecase/usecase_user"
)

func MakeFindEndpoint(u usecase_user.UseCase) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		users, err := u.Find()
		if err != nil {
			return nil, err
		}
		return users, nil
	}
}

func MakeCreateEndpoint(u usecase_user.UseCase) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(createRequest)
		user, err := u.Create(req.Name)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
}

func DecodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request createRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

type createRequest struct {
	Name string
}

func MakeFirstEndpoint(u usecase_user.UseCase) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(string)
		uid, _ := strconv.Atoi(req)
		user, err := u.First(uint64(uid))
		if err != nil {
			return nil, err
		}
		return user, nil
	}
}

func DecodeFirstRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	uid, ok := vars["uid"]
	if !ok {
		return nil, errors.New("uid parse error")
	}
	return uid, nil
}

func MakeDeleteEndpoint(u usecase_user.UseCase) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(string)
		uid, _ := strconv.Atoi(req)
		if err := u.Delete(uint64(uid)); err != nil {
			return nil, err
		}
		return "deleted", nil
	}
}

func DecodeDeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	uid, ok := vars["uid"]
	if !ok {
		return nil, errors.New("uid parse error")
	}
	return uid, nil
}

func MakeCreateEmailEndpoint(u usecase_user.UseCase) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(createEmailRequest)
		uid, _ := strconv.Atoi(req.Uid)
		email, err := u.CreateEmail(uint64(uid), req.Email)
		if err != nil {
			return nil, err
		}
		return email, nil
	}
}

func DecodeCreateEmailRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	uid, ok := vars["uid"]
	if !ok {
		return nil, errors.New("uid parse error")
	}
	var request createEmailRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	request.Uid = uid
	return request, nil
}

type createEmailRequest struct {
	Uid   string
	Email string
}
