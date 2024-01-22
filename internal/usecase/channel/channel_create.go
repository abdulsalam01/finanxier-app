package channel

import (
	"context"

	"github.com/api-sekejap/internal/entity"
)

func (c *channelUsecase) Create(ctx context.Context, params entity.Channel) error {
	_, err := c.channelRepo.Create(ctx, params)
	if err != nil {
		return err
	}

	return nil
}
