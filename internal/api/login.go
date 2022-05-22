package api

import (
	"context"

	"github.com/artemmarkaryan/fisha-facade/pkg/marchy"
	"github.com/artemmarkaryan/fisha-facade/pkg/pb/gen/api"
)

func (a API) Login(ctx context.Context, user int64) (isNew bool, err error) {
	r, err := a.post("/login", api.UserIdRequest{UserId: user})
	if err != nil {
		return
	}

	obj, err := marchy.Obj[*api.LoginResponse](ctx, r.Body)
	if err != nil {
		return
	}

	return obj.GetNew(), nil
}
