// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: com/mixi/mercury/api/mixi2.proto

package apiconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	api "github.com/matsuu/go-mixi2/gen/com/mixi/mercury/api"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// MercuryServiceName is the fully-qualified name of the MercuryService service.
	MercuryServiceName = "com.mixi.mercury.api.MercuryService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// MercuryServiceCreatePostProcedure is the fully-qualified name of the MercuryService's CreatePost
	// RPC.
	MercuryServiceCreatePostProcedure = "/com.mixi.mercury.api.MercuryService/CreatePost"
	// MercuryServiceGetPostsProcedure is the fully-qualified name of the MercuryService's GetPosts RPC.
	MercuryServiceGetPostsProcedure = "/com.mixi.mercury.api.MercuryService/GetPosts"
	// MercuryServiceSwitchPersonaProcedure is the fully-qualified name of the MercuryService's
	// SwitchPersona RPC.
	MercuryServiceSwitchPersonaProcedure = "/com.mixi.mercury.api.MercuryService/SwitchPersona"
	// MercuryServiceGetPersonasProcedure is the fully-qualified name of the MercuryService's
	// GetPersonas RPC.
	MercuryServiceGetPersonasProcedure = "/com.mixi.mercury.api.MercuryService/GetPersonas"
)

// MercuryServiceClient is a client for the com.mixi.mercury.api.MercuryService service.
type MercuryServiceClient interface {
	CreatePost(context.Context, *connect.Request[api.CreatePostRequest]) (*connect.Response[api.CreatePostResponse], error)
	GetPosts(context.Context, *connect.Request[api.GetPostsRequest]) (*connect.Response[api.GetPostsResponse], error)
	SwitchPersona(context.Context, *connect.Request[api.SwitchPersonaRequest]) (*connect.Response[api.SwitchPersonaResponse], error)
	GetPersonas(context.Context, *connect.Request[api.GetPersonasRequest]) (*connect.Response[api.GetPersonasResponse], error)
}

// NewMercuryServiceClient constructs a client for the com.mixi.mercury.api.MercuryService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewMercuryServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) MercuryServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	mercuryServiceMethods := api.File_com_mixi_mercury_api_mixi2_proto.Services().ByName("MercuryService").Methods()
	return &mercuryServiceClient{
		createPost: connect.NewClient[api.CreatePostRequest, api.CreatePostResponse](
			httpClient,
			baseURL+MercuryServiceCreatePostProcedure,
			connect.WithSchema(mercuryServiceMethods.ByName("CreatePost")),
			connect.WithClientOptions(opts...),
		),
		getPosts: connect.NewClient[api.GetPostsRequest, api.GetPostsResponse](
			httpClient,
			baseURL+MercuryServiceGetPostsProcedure,
			connect.WithSchema(mercuryServiceMethods.ByName("GetPosts")),
			connect.WithClientOptions(opts...),
		),
		switchPersona: connect.NewClient[api.SwitchPersonaRequest, api.SwitchPersonaResponse](
			httpClient,
			baseURL+MercuryServiceSwitchPersonaProcedure,
			connect.WithSchema(mercuryServiceMethods.ByName("SwitchPersona")),
			connect.WithClientOptions(opts...),
		),
		getPersonas: connect.NewClient[api.GetPersonasRequest, api.GetPersonasResponse](
			httpClient,
			baseURL+MercuryServiceGetPersonasProcedure,
			connect.WithSchema(mercuryServiceMethods.ByName("GetPersonas")),
			connect.WithClientOptions(opts...),
		),
	}
}

// mercuryServiceClient implements MercuryServiceClient.
type mercuryServiceClient struct {
	createPost    *connect.Client[api.CreatePostRequest, api.CreatePostResponse]
	getPosts      *connect.Client[api.GetPostsRequest, api.GetPostsResponse]
	switchPersona *connect.Client[api.SwitchPersonaRequest, api.SwitchPersonaResponse]
	getPersonas   *connect.Client[api.GetPersonasRequest, api.GetPersonasResponse]
}

// CreatePost calls com.mixi.mercury.api.MercuryService.CreatePost.
func (c *mercuryServiceClient) CreatePost(ctx context.Context, req *connect.Request[api.CreatePostRequest]) (*connect.Response[api.CreatePostResponse], error) {
	return c.createPost.CallUnary(ctx, req)
}

// GetPosts calls com.mixi.mercury.api.MercuryService.GetPosts.
func (c *mercuryServiceClient) GetPosts(ctx context.Context, req *connect.Request[api.GetPostsRequest]) (*connect.Response[api.GetPostsResponse], error) {
	return c.getPosts.CallUnary(ctx, req)
}

// SwitchPersona calls com.mixi.mercury.api.MercuryService.SwitchPersona.
func (c *mercuryServiceClient) SwitchPersona(ctx context.Context, req *connect.Request[api.SwitchPersonaRequest]) (*connect.Response[api.SwitchPersonaResponse], error) {
	return c.switchPersona.CallUnary(ctx, req)
}

// GetPersonas calls com.mixi.mercury.api.MercuryService.GetPersonas.
func (c *mercuryServiceClient) GetPersonas(ctx context.Context, req *connect.Request[api.GetPersonasRequest]) (*connect.Response[api.GetPersonasResponse], error) {
	return c.getPersonas.CallUnary(ctx, req)
}

// MercuryServiceHandler is an implementation of the com.mixi.mercury.api.MercuryService service.
type MercuryServiceHandler interface {
	CreatePost(context.Context, *connect.Request[api.CreatePostRequest]) (*connect.Response[api.CreatePostResponse], error)
	GetPosts(context.Context, *connect.Request[api.GetPostsRequest]) (*connect.Response[api.GetPostsResponse], error)
	SwitchPersona(context.Context, *connect.Request[api.SwitchPersonaRequest]) (*connect.Response[api.SwitchPersonaResponse], error)
	GetPersonas(context.Context, *connect.Request[api.GetPersonasRequest]) (*connect.Response[api.GetPersonasResponse], error)
}

// NewMercuryServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewMercuryServiceHandler(svc MercuryServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	mercuryServiceMethods := api.File_com_mixi_mercury_api_mixi2_proto.Services().ByName("MercuryService").Methods()
	mercuryServiceCreatePostHandler := connect.NewUnaryHandler(
		MercuryServiceCreatePostProcedure,
		svc.CreatePost,
		connect.WithSchema(mercuryServiceMethods.ByName("CreatePost")),
		connect.WithHandlerOptions(opts...),
	)
	mercuryServiceGetPostsHandler := connect.NewUnaryHandler(
		MercuryServiceGetPostsProcedure,
		svc.GetPosts,
		connect.WithSchema(mercuryServiceMethods.ByName("GetPosts")),
		connect.WithHandlerOptions(opts...),
	)
	mercuryServiceSwitchPersonaHandler := connect.NewUnaryHandler(
		MercuryServiceSwitchPersonaProcedure,
		svc.SwitchPersona,
		connect.WithSchema(mercuryServiceMethods.ByName("SwitchPersona")),
		connect.WithHandlerOptions(opts...),
	)
	mercuryServiceGetPersonasHandler := connect.NewUnaryHandler(
		MercuryServiceGetPersonasProcedure,
		svc.GetPersonas,
		connect.WithSchema(mercuryServiceMethods.ByName("GetPersonas")),
		connect.WithHandlerOptions(opts...),
	)
	return "/com.mixi.mercury.api.MercuryService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case MercuryServiceCreatePostProcedure:
			mercuryServiceCreatePostHandler.ServeHTTP(w, r)
		case MercuryServiceGetPostsProcedure:
			mercuryServiceGetPostsHandler.ServeHTTP(w, r)
		case MercuryServiceSwitchPersonaProcedure:
			mercuryServiceSwitchPersonaHandler.ServeHTTP(w, r)
		case MercuryServiceGetPersonasProcedure:
			mercuryServiceGetPersonasHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedMercuryServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedMercuryServiceHandler struct{}

func (UnimplementedMercuryServiceHandler) CreatePost(context.Context, *connect.Request[api.CreatePostRequest]) (*connect.Response[api.CreatePostResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("com.mixi.mercury.api.MercuryService.CreatePost is not implemented"))
}

func (UnimplementedMercuryServiceHandler) GetPosts(context.Context, *connect.Request[api.GetPostsRequest]) (*connect.Response[api.GetPostsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("com.mixi.mercury.api.MercuryService.GetPosts is not implemented"))
}

func (UnimplementedMercuryServiceHandler) SwitchPersona(context.Context, *connect.Request[api.SwitchPersonaRequest]) (*connect.Response[api.SwitchPersonaResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("com.mixi.mercury.api.MercuryService.SwitchPersona is not implemented"))
}

func (UnimplementedMercuryServiceHandler) GetPersonas(context.Context, *connect.Request[api.GetPersonasRequest]) (*connect.Response[api.GetPersonasResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("com.mixi.mercury.api.MercuryService.GetPersonas is not implemented"))
}
