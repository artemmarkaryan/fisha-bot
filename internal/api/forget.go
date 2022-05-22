package api

import (
	"context"

	"github.com/artemmarkaryan/fisha-facade/pkg/pb/gen/api"
)

func (a API) Forget(ctx context.Context, user int64) error {
	_, err := a.post("/forget", api.UserIdRequest{UserId: user})
	if err != nil {
		return err
	}

	return nil
}
