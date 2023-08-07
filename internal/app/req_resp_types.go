package app

import "github.com/bersennaidoo/Loggeez/internal/domain/model"

type ProduceRequest struct {
	Record model.Record `json:"record"`
}

type ProduceResponse struct {
	Offset uint64 `json:"offset"`
}

type ConsumeRequest struct {
	Offset uint64 `json:"offset"`
}

type ConsumeResponse struct {
	Record model.Record `json:"record"`
}
