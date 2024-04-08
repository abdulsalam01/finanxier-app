package app

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

// Private struct helper.
// Goes here:
type dbQueryTracer struct {
	log *logrus.Logger
}

type ValidatorResource interface {
	Struct(s interface{}) error
	StructCtx(ctx context.Context, s interface{}) (err error)
	StructExcept(s interface{}, fields ...string) error
	StructExceptCtx(ctx context.Context, s interface{}, fields ...string) (err error)
	StructFiltered(s interface{}, fn validator.FilterFunc) error
	StructFilteredCtx(ctx context.Context, s interface{}, fn validator.FilterFunc) (err error)
	StructPartial(s interface{}, fields ...string) error
	StructPartialCtx(ctx context.Context, s interface{}, fields ...string) (err error)
}

func (tracer *dbQueryTracer) TraceQueryStart(ctx context.Context, _ *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	tracer.log.Infof("Executing command sql: %s with args: %v", data.SQL, data.Args)
	return ctx
}
func (tracer *dbQueryTracer) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {
	// Add logic here...
}
