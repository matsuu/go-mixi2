package mixi2

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"connectrpc.com/connect"

	"github.com/matsuu/go-mixi2/gen/com/mixi/mercury/api/apiconnect"
)

const (
	baseUri = `https://mixi.social/api/connect`

	cookieHeader        = "Cookie"
	authTokenCookieName = "auth_token"
	authKeyHeader       = "X-Auth-Key"
	userAgentHeader     = "X-Mercury-User-Agent"
)

func WithAuth(authKey, authToken, userAgent string) connect.ClientOption {
	return connect.WithInterceptors(func() connect.UnaryInterceptorFunc {
		interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
			return connect.UnaryFunc(func(
				ctx context.Context,
				req connect.AnyRequest,
			) (connect.AnyResponse, error) {
				if req.Spec().IsClient {
					req.Header().Set(authKeyHeader, authKey)
					req.Header().Set(userAgentHeader, userAgent)
					req.Header().Set(cookieHeader, fmt.Sprintf("%s=%s", authTokenCookieName, authToken))
				} else {
					var missings []string
					if req.Header().Get(authKeyHeader) == "" {
						missings = append(missings, authKeyHeader)
					}
					if req.Header().Get(userAgentHeader) == "" {
						missings = append(missings, userAgentHeader)
					}
					if req.Header().Get(cookieHeader) == "" {
						missings = append(missings, cookieHeader)
					} else {
						cookies, err := http.ParseCookie(req.Header().Get(cookieHeader))
						if err != nil {
							return nil, connect.NewError(
								connect.CodeInternal,
								fmt.Errorf("Failed to parse cookie: %w", err),
							)
						}
						authToken := ""
						for _, cookie := range cookies {
							if cookie.Name == authTokenCookieName {
								authToken = cookie.Value
								break
							}
						}
						if authToken == "" {
							missings = append(missings, authTokenCookieName)
						}
					}
					if len(missings) > 0 {
						return nil, connect.NewError(
							connect.CodeUnauthenticated,
							fmt.Errorf("no %s provided", strings.Join(missings, ", ")),
						)
					}
				}
				return next(ctx, req)
			})
		}
		return connect.UnaryInterceptorFunc(interceptor)
	}())
}

func NewClient(opts ...connect.ClientOption) apiconnect.MercuryServiceClient {
	client := apiconnect.NewMercuryServiceClient(
		http.DefaultClient,
		baseUri,
		opts...,
	)
	return client
}
