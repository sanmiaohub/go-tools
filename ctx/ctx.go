package ctx

import (
	"context"
	"github.com/satori/go.uuid"
)

const (
	RequestIDKey string = "x-request-id"
)

type Context interface {
	ID() string
	context.Context
}

type defaultContext struct {
	id string
	context.Context
}

func (d *defaultContext) ID() string {
	return d.id
}

func New() Context {
	p, id := xReqIDCtx(context.TODO())
	return &defaultContext{Context: p, id: id}
}

func Wrap(c context.Context) Context {
	if cc, ok := c.(Context); ok {
		return cc
	}
	p, id := xReqIDCtx(c)
	return &defaultContext{Context: p, id: id}
}

func xReqIDCtx(c context.Context) (context.Context, string) {
	if c == nil {
		panic("xReqIDCtx：context.Context 不能传入 nil")
	}
	val := c.Value(RequestIDKey)
	if id, ok := val.(string); ok && id != "" {
		return c, id
	}
	id := uuid.NewV4().String()
	return context.WithValue(c, RequestIDKey, id), id
}
