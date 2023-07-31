// Copyright 2020-2023 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: buf/alpha/registry/v1alpha1/repository_branch.proto

package registryv1alpha1connect

import (
	context "context"
	errors "errors"
	v1alpha1 "github.com/bufbuild/buf/private/gen/proto/go/buf/alpha/registry/v1alpha1"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion1_7_0

const (
	// RepositoryBranchServiceName is the fully-qualified name of the RepositoryBranchService service.
	RepositoryBranchServiceName = "buf.alpha.registry.v1alpha1.RepositoryBranchService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// RepositoryBranchServiceListRepositoryBranchesProcedure is the fully-qualified name of the
	// RepositoryBranchService's ListRepositoryBranches RPC.
	RepositoryBranchServiceListRepositoryBranchesProcedure = "/buf.alpha.registry.v1alpha1.RepositoryBranchService/ListRepositoryBranches"
	// RepositoryBranchServiceGetCurrentDefaultBranchProcedure is the fully-qualified name of the
	// RepositoryBranchService's GetCurrentDefaultBranch RPC.
	RepositoryBranchServiceGetCurrentDefaultBranchProcedure = "/buf.alpha.registry.v1alpha1.RepositoryBranchService/GetCurrentDefaultBranch"
)

// RepositoryBranchServiceClient is a client for the
// buf.alpha.registry.v1alpha1.RepositoryBranchService service.
type RepositoryBranchServiceClient interface {
	// ListRepositoryBranchs lists the repository branches associated with a Repository.
	ListRepositoryBranches(context.Context, *connect_go.Request[v1alpha1.ListRepositoryBranchesRequest]) (*connect_go.Response[v1alpha1.ListRepositoryBranchesResponse], error)
	// GetCurrentDefaultBranch returns the branch name that is mapped to the main/BSR_HEAD. This might
	// not be the same value in the repository's `default_branch` field, since that value can be
	// changed at will by repository's owners/admins for syncing git repositories. This RPC retrieves
	// the branch from the latest commit labeled as BSR_HEAD, even if that value differs from the one
	// stored in the `default_branch` field.
	//
	// TODO: Rename this RPC to something more appropriate like "GetLatestHEADCommit".
	GetCurrentDefaultBranch(context.Context, *connect_go.Request[v1alpha1.GetCurrentDefaultBranchRequest]) (*connect_go.Response[v1alpha1.GetCurrentDefaultBranchResponse], error)
}

// NewRepositoryBranchServiceClient constructs a client for the
// buf.alpha.registry.v1alpha1.RepositoryBranchService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewRepositoryBranchServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) RepositoryBranchServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &repositoryBranchServiceClient{
		listRepositoryBranches: connect_go.NewClient[v1alpha1.ListRepositoryBranchesRequest, v1alpha1.ListRepositoryBranchesResponse](
			httpClient,
			baseURL+RepositoryBranchServiceListRepositoryBranchesProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		getCurrentDefaultBranch: connect_go.NewClient[v1alpha1.GetCurrentDefaultBranchRequest, v1alpha1.GetCurrentDefaultBranchResponse](
			httpClient,
			baseURL+RepositoryBranchServiceGetCurrentDefaultBranchProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
	}
}

// repositoryBranchServiceClient implements RepositoryBranchServiceClient.
type repositoryBranchServiceClient struct {
	listRepositoryBranches  *connect_go.Client[v1alpha1.ListRepositoryBranchesRequest, v1alpha1.ListRepositoryBranchesResponse]
	getCurrentDefaultBranch *connect_go.Client[v1alpha1.GetCurrentDefaultBranchRequest, v1alpha1.GetCurrentDefaultBranchResponse]
}

// ListRepositoryBranches calls
// buf.alpha.registry.v1alpha1.RepositoryBranchService.ListRepositoryBranches.
func (c *repositoryBranchServiceClient) ListRepositoryBranches(ctx context.Context, req *connect_go.Request[v1alpha1.ListRepositoryBranchesRequest]) (*connect_go.Response[v1alpha1.ListRepositoryBranchesResponse], error) {
	return c.listRepositoryBranches.CallUnary(ctx, req)
}

// GetCurrentDefaultBranch calls
// buf.alpha.registry.v1alpha1.RepositoryBranchService.GetCurrentDefaultBranch.
func (c *repositoryBranchServiceClient) GetCurrentDefaultBranch(ctx context.Context, req *connect_go.Request[v1alpha1.GetCurrentDefaultBranchRequest]) (*connect_go.Response[v1alpha1.GetCurrentDefaultBranchResponse], error) {
	return c.getCurrentDefaultBranch.CallUnary(ctx, req)
}

// RepositoryBranchServiceHandler is an implementation of the
// buf.alpha.registry.v1alpha1.RepositoryBranchService service.
type RepositoryBranchServiceHandler interface {
	// ListRepositoryBranchs lists the repository branches associated with a Repository.
	ListRepositoryBranches(context.Context, *connect_go.Request[v1alpha1.ListRepositoryBranchesRequest]) (*connect_go.Response[v1alpha1.ListRepositoryBranchesResponse], error)
	// GetCurrentDefaultBranch returns the branch name that is mapped to the main/BSR_HEAD. This might
	// not be the same value in the repository's `default_branch` field, since that value can be
	// changed at will by repository's owners/admins for syncing git repositories. This RPC retrieves
	// the branch from the latest commit labeled as BSR_HEAD, even if that value differs from the one
	// stored in the `default_branch` field.
	//
	// TODO: Rename this RPC to something more appropriate like "GetLatestHEADCommit".
	GetCurrentDefaultBranch(context.Context, *connect_go.Request[v1alpha1.GetCurrentDefaultBranchRequest]) (*connect_go.Response[v1alpha1.GetCurrentDefaultBranchResponse], error)
}

// NewRepositoryBranchServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewRepositoryBranchServiceHandler(svc RepositoryBranchServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	repositoryBranchServiceListRepositoryBranchesHandler := connect_go.NewUnaryHandler(
		RepositoryBranchServiceListRepositoryBranchesProcedure,
		svc.ListRepositoryBranches,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	)
	repositoryBranchServiceGetCurrentDefaultBranchHandler := connect_go.NewUnaryHandler(
		RepositoryBranchServiceGetCurrentDefaultBranchProcedure,
		svc.GetCurrentDefaultBranch,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	)
	return "/buf.alpha.registry.v1alpha1.RepositoryBranchService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case RepositoryBranchServiceListRepositoryBranchesProcedure:
			repositoryBranchServiceListRepositoryBranchesHandler.ServeHTTP(w, r)
		case RepositoryBranchServiceGetCurrentDefaultBranchProcedure:
			repositoryBranchServiceGetCurrentDefaultBranchHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedRepositoryBranchServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedRepositoryBranchServiceHandler struct{}

func (UnimplementedRepositoryBranchServiceHandler) ListRepositoryBranches(context.Context, *connect_go.Request[v1alpha1.ListRepositoryBranchesRequest]) (*connect_go.Response[v1alpha1.ListRepositoryBranchesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.alpha.registry.v1alpha1.RepositoryBranchService.ListRepositoryBranches is not implemented"))
}

func (UnimplementedRepositoryBranchServiceHandler) GetCurrentDefaultBranch(context.Context, *connect_go.Request[v1alpha1.GetCurrentDefaultBranchRequest]) (*connect_go.Response[v1alpha1.GetCurrentDefaultBranchResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.alpha.registry.v1alpha1.RepositoryBranchService.GetCurrentDefaultBranch is not implemented"))
}
