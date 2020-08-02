package client

import (
	"context"
	"github.com/go-masonry/mortar/constructors/partial"
	"github.com/go-masonry/mortar/interfaces/cfg"
	"github.com/go-masonry/mortar/mortar"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strings"
)

func CopyGRPCHeadersClientInterceptorOption() fx.Option {
	return fx.Provide(
		fx.Annotated{
			Group:  partial.FxGroupGRPCUnaryClientInterceptors,
			Target: CopyGRPCHeadersClientInterceptor,
		})
}

type copyHeadersDeps struct {
	fx.In

	Config cfg.Config
}

func CopyGRPCHeadersClientInterceptor(deps copyHeadersDeps) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			headerPrefixes := deps.Config.Get(mortar.MiddlewareServerGRPCCopyHeadersPrefixes).StringSlice()
			for _, headerPrefix := range headerPrefixes {
				for k, vs := range md {
					if strings.HasPrefix(strings.ToLower(k), headerPrefix) {
						for _, v := range vs {
							ctx = metadata.AppendToOutgoingContext(ctx, k, v)
						}
					}
				}
			}
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
