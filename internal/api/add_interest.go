package api

import (
	"context"

	"github.com/artemmarkaryan/fisha-facade/pkg/marchy"
	"github.com/artemmarkaryan/fisha-facade/pkg/pb/gen/api"
)

func (a API) AddInterest(ctx context.Context, user, interest int64) (isNew bool, err error) {
	resp, err := a.post(ctx, "/addInterest", api.AddInterestRequest{
		InterestId: interest,
		UserId:     user,
	})
	if err != nil {
		return
	}

	obj, err := marchy.Obj[*api.IsNewMessage](ctx, resp.Body)
	if err != nil {
		return
	}

	return obj.GetNew(), nil
}
