package api

import (
	"context"

	"github.com/artemmarkaryan/fisha-facade/pkg/marchy"
	"github.com/artemmarkaryan/fisha-facade/pkg/pb/gen/api"
)

func (a API) UserHasLocation(ctx context.Context, userID int64) (hasLocation bool, err error) {
	r, err := a.post(ctx,
		"/user/has-location",
		api.IdMessage{Id: userID},
	)

	boolMessage, err := marchy.Obj[*api.BooleanMessage](ctx, r.Body)
	if err != nil {
		return
	}

	return boolMessage.GetResult(), err
}
