package api

import (
	"context"

	"github.com/artemmarkaryan/fisha-facade/pkg/pb/gen/api"
)

func (a API) Forget(ctx context.Context, user int64) error {
	_, err := a.post(ctx, "/forget", api.IdMessage{Id: user})
	if err != nil {
		return err
	}

	return nil
}
