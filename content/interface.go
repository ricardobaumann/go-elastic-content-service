package content

import (
	"context"
)

type Service interface {
	Get(ctx context.Context, obj contentRequest) (contentResponse, error)
	Delete(ctx context.Context, obj contentRequest) (contentResponse, error)
	Save(ctx context.Context, obj contentInput) (contentResponse, error)
}

type Repository interface {
	Get(ID string) (contentResponse, error)
	Delete(ID string) (contentResponse, error)
	Save(ID string, Body string) (contentResponse, error)
}
