package service

import (
	"context"
	"fmt"

	"github.com/elin-croft/go-utils/gen-go/ew_toolbox"
)

type PredictHandler struct{}

func NewProcessor() *ew_toolbox.PredictProcessor {
	processor := ew_toolbox.NewPredictProcessor(&PredictHandler{})
	return processor
}

func (p *PredictHandler) Get(ctx context.Context, req *ew_toolbox.Request) (*ew_toolbox.Response, error) {
	fmt.Printf("request: %v\n", req)
	return &ew_toolbox.Response{ID: req.ID, Name: "Hello, World!"}, nil
}
