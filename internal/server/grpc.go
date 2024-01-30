package server

import (
	v1 "github.com/Allen-Career-Institute/go-kratos-sample/api/user/v1"
	"github.com/Allen-Career-Institute/go-kratos-sample/internal/conf"
	"github.com/Allen-Career-Institute/go-kratos-sample/internal/request"
	"github.com/Allen-Career-Institute/go-kratos-sample/internal/service"
	"go.opentelemetry.io/otel"

	"github.com/Allen-Career-Institute/go-kratos-commons/otel/middleware/metric"
	"github.com/Allen-Career-Institute/go-kratos-commons/otel/middleware/trace"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
// No Lint is just used for sample code. We should avoid it
// nolint: revive
func NewGRPCServer(c *conf.Server, user *service.UserService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			validate.Validator(),
			metric.Meter(otel.Meter("meter-go-kratos-sample")),
			trace.Trace(otel.Tracer("tracer-go-kratos-sample")),
			request.Populate(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterUserServer(srv, user)
	return srv
}
