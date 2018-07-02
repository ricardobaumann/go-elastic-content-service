package content

import (
	"context"
)

type Service interface {
	Get(ctx context.Context, obj contentRequest) (contentResponse, error)
}

type Repository interface {
	Get(ID string) (contentResponse, error)
}
