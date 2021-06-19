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

// SetQuotaRequest allows to set quota.
type SetQuotaRequest struct {
	base
	req *pb.SetQuotaRequest
}

// SetNamespaceSpaceQuotaRequest allows to set namespace space quota
type SetNamespaceSpaceQuotaRequest SetQuotaRequest

// NewSetNamespaceSpaceQuotaRequest creates a namespace space quota request.
func NewSetNamespaceSpaceQuotaRequest(ctx context.Context, namespace string, quota uint64) (*SetNamespaceSpaceQuotaRequest, error) {
	noInsertsViolationPolicy := pb.SpaceViolationPolicy_NO_INSERTS
	remove := false

	return &SetNamespaceSpaceQuotaRequest{
		base: base{
			ctx:      ctx,
			resultch: make(chan RPCResult, 1),
		},
		req: &pb.SetQuotaRequest{
			Namespace: &namespace,
			SpaceLimit: &pb.SpaceLimitRequest{
				Quota: &pb.SpaceQuota{
					SoftLimit:       &quota,
					ViolationPolicy: &noInsertsViolationPolicy,
					Remove:          &remove,
				},
			},
		},
	}, nil
}

// Name returns the name of this RPC call.
func (nsqr *SetNamespaceSpaceQuotaRequest) Name() string {
	return "SetQuotaRequest"
}

// ToProto converts the RPC into a protobuf message.
func (nsqr *SetNamespaceSpaceQuotaRequest) ToProto() proto.Message {
	return nsqr.req
}

// NewResponse creates an empty protobuf message to read the response of this RPC.
func (nsqr *SetNamespaceSpaceQuotaRequest) NewResponse() proto.Message {
	return &pb.SetQuotaResponse{}
}
