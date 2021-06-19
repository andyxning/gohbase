// Copyright (C) 2020  The GoHBase Authors.  All rights reserved.
// This file is part of GoHBase.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

package hrpc

import (
	"context"

	"github.com/tsuna/gohbase/pb"
	"google.golang.org/protobuf/proto"
)

// QuotaStatesRequest represents a get quota states request
type QuotaStatesRequest struct {
	base
}

// NewQuotaStatesRequest creates a new QuotaStatesRequest with default fields
func NewQuotaStatesRequest() *QuotaStatesRequest {
	return &QuotaStatesRequest{
		base{
			ctx:      context.Background(),
			table:    []byte{},
			resultch: make(chan RPCResult, 1),
		},
	}
}

// Name returns the name of the rpc function
func (qs *QuotaStatesRequest) Name() string {
	return "GetQuotaStates"
}

// ToProto returns the Protobuf message to be sent
func (qs *QuotaStatesRequest) ToProto() proto.Message {
	return &pb.GetQuotaStatesRequest{}
}

// NewResponse returns the empty protobuf response
func (qs *QuotaStatesRequest) NewResponse() proto.Message {
	return &pb.GetQuotaStatesResponse{}
}
