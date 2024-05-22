package sdk

import (
	"context"
	"net/http"
)

type FingerprintApiService struct{}

func (f *FingerprintApiService) DeleteVisitorData(ctx context.Context, visitorId string) (*http.Response, error) {
	// Implement your method here
	return nil, nil
}

func (f *FingerprintApiService) GetEvent(ctx context.Context, requestId string) (EventResponse, *http.Response, error) {
	// Implement your method here
	return EventResponse{}, nil, nil
}

func (f *FingerprintApiService) GetVisits(ctx context.Context, visitorId string, opts *FingerprintApiGetVisitsOpts) (Response, *http.Response, error) {
	// Implement your method here
	return Response{}, nil, nil
}
