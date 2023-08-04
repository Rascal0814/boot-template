package service

import (
	"context"
	v1 "hml/api/v1"
)

type EchoService struct {
	v1.UnimplementedEchoServiceServer
}

func (s EchoService) Say(ctx context.Context, resp *v1.SayRequest) (*v1.SayResponse, error) {
	return &v1.SayResponse{Text: "hello world " + resp.Text}, nil
}

func NewEchoService() *EchoService {
	return &EchoService{}
}
