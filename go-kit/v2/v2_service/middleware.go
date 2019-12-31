package v2_service

import (
	"context"
	"go.uber.org/zap"
)

const ContextReqUUid = "req_uuid"

type NewMiddlewareServer func(Service) Service

type logMiddlewareServer struct {
	logger *zap.Logger
	next   Service
}

func NewLogMiddlewareServer(log *zap.Logger) NewMiddlewareServer {
	return func(service Service) Service {
		return logMiddlewareServer{
			logger: log,
			next:   service,
		}
	}
}

func (l logMiddlewareServer) TestAdd(ctx context.Context, in Add) (out AddAck) {
	defer func() {
		l.logger.Debug(ctx.Value(ContextReqUUid).(string), zap.Any("v2_service logMiddlewareServer", "TestAdd"), zap.Any("req", in), zap.Any("res", out))
	}()
	out = l.next.TestAdd(ctx, in)
	return out
}
