package channel

import (
	"context"

	"github.com/api-sekejap/internal/entity"
)

type channelResource interface {
	Create(ctx context.Context, params entity.Channel) (int, error)
}

type channelUsecase struct {
	channelRepo channelResource
}
