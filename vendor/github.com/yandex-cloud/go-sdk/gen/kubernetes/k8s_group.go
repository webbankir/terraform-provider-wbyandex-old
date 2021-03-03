// Code generated by sdkgen. DO NOT EDIT.

package k8s

import (
	"context"

	"google.golang.org/grpc"
)

// Kubernetes provides access to "k8s" component of Yandex.Cloud
type Kubernetes struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// NewKubernetes creates instance of Kubernetes
func NewKubernetes(g func(ctx context.Context) (*grpc.ClientConn, error)) *Kubernetes {
	return &Kubernetes{g}
}

// Cluster gets ClusterService client
func (k *Kubernetes) Cluster() *ClusterServiceClient {
	return &ClusterServiceClient{getConn: k.getConn}
}

// NodeGroup gets NodeGroupService client
func (k *Kubernetes) NodeGroup() *NodeGroupServiceClient {
	return &NodeGroupServiceClient{getConn: k.getConn}
}

// Version gets VersionService client
func (k *Kubernetes) Version() *VersionServiceClient {
	return &VersionServiceClient{getConn: k.getConn}
}
