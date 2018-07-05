package content

import (
	"context"

	tracing "github.com/ricardo-ch/go-tracing"
)

type contentTracing struct {
	next Service
}

// NewTracing ...
func NewTracing(s Service) Service {
	return contentTracing{
		next: s,
	}
}

// Get ...
func (s contentTracing) Get(ctx context.Context, obj contentRequest) (response contentResponse, err error) {
	span, ctx := tracing.CreateSpan(ctx, "content.service::Get", &map[string]interface{}{"id": obj.ID})
	defer func() {
		if err != nil {
			tracing.SetSpanError(span, err)
		}
		span.Finish()
	}()

	return s.next.Get(ctx, obj)
}

// Get ...
func (s contentTracing) Save(ctx context.Context, obj contentInput) (response contentResponse, err error) {
	span, ctx := tracing.CreateSpan(ctx, "content.service::Save", &map[string]interface{}{"id": obj.ID})
	defer func() {
		if err != nil {
			tracing.SetSpanError(span, err)
		}
		span.Finish()
	}()

	return s.next.Save(ctx, obj)
}
