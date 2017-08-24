package marketservice

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func MakeRegisterUserEndpoint(svc DataService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RegisterUserRequest)
		err := svc.RegisterUser(req.Id, req.Username)
		if err != nil {
			return RegisterUserResponse{err.Error()}, nil
		}
		return RegisterUserResponse{"nil"}, nil
	}
}

func DecodeRegisterUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request RegisterUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func MakeGetMarketSummaryEndpoint(svc DataService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(GetMarketSummaryRequest)
		fmt.Println("sebelum masuk=" + req.Code)
		ms, err := svc.GetMarketSummary(req.Code)
		if err != nil {
			return GetMarketSummaryResponse{nil, err.Error()}, nil
		}
		return GetMarketSummaryResponse{&ms, "nil"}, nil
	}
}

func DecodeGetMarketSummaryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request GetMarketSummaryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// Register User
type RegisterUserRequest struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}
type RegisterUserResponse struct {
	Err string `json:"err,omitempty"`
}

// Get Market Summary
type GetMarketSummaryRequest struct {
	Code string `json:"code"`
}

type GetMarketSummaryResponse struct {
	Summary *MarketSummary `json:"summary,omitempty"`
	Err     string         `json:"err,omitempty"`
}
