package hrpc

import (
	"context"
	"github.com/tsuna/gohbase/pb"
	"google.golang.org/protobuf/proto"
)

// ListNamespaceDescriptors models a ListNamespaceDescriptors pb call
type ListNamespaceDescriptors struct {
	base
}

// NewListNamespaceDescriptors creates a new ListNamespaceDescriptors request that will list namespaces in hbase.
func NewListNamespaceDescriptors(ctx context.Context, opts ...func(Call) error) (*ListNamespaceDescriptors, error) {
	tn := &ListNamespaceDescriptors{
		base: base{
			ctx:      ctx,
			resultch: make(chan RPCResult, 1),
		},
	}
	if err := applyOptions(tn, opts...); err != nil {
		return nil, err
	}
	return tn, nil
}

// Name returns the name of this RPC call.
func (tn *ListNamespaceDescriptors) Name() string {
	return "ListNamespaceDescriptors"
}

// ToProto converts the RPC into a protobuf message.
func (tn *ListNamespaceDescriptors) ToProto() proto.Message {
	return &pb.ListNamespaceDescriptorsRequest{
	}
}

// NewResponse creates an empty protobuf message to read the response of this
// RPC.
func (tn *ListNamespaceDescriptors) NewResponse() proto.Message {
	return &pb.ListNamespaceDescriptorsResponse{}
}
