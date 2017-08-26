package service

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func MakeGetCurrentTickEndpoint(svc MarketService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetCurrentTickRequest)
		tick, err := svc.GetCurrentTick(req.Eid, req.MCode)
		if err != nil {
			return GetCurrentTickResponse{tick, err.Error()}, err
		}
		return GetCurrentTickResponse{tick, "nil"}, nil
	}
}
func DecodeGetCurrentTickRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request GetCurrentTickRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func MakeGetCurrentTick2Endpoint(svc MarketService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetCurrentTick2Request)
		tick, err := svc.GetCurrentTick2(req.Eid, req.C1, req.C2)
		if err != nil {
			return GetCurrentTickResponse{tick, err.Error()}, err
		}
		return GetCurrentTickResponse{tick, "nil"}, nil
	}
}
func DecodeGetCurrentTick2Request(_ context.Context, r *http.Request) (interface{}, error) {
	var request GetCurrentTick2Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// Get Current Tick
type GetCurrentTickRequest struct {
	Eid   int    `json:"eid"`
	MCode string `json:"m_code"`
}
type GetCurrentTickResponse struct {
	Tick Ticker `json:"tick"`
	Err  string `json:"err,omitempty"`
}

// Get Current Tick
type GetCurrentTick2Request struct {
	Eid int    `json:"eid"`
	C1  string `json:"c1"`
	C2  string `json:"c2"`
}
