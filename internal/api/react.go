package api

import (
	"context"

	"github.com/artemmarkaryan/fisha-facade/pkg/pb/gen/api"
	"github.com/artemmarkaryan/fisha/bot/internal/service/reaction"
)

func (a API) React(ctx context.Context, user, activity int64, reaction reaction.Reaction) (isNew bool, err error) {
	_, err = a.post(ctx,
		"/react",
		api.ReactRequest{UserId: user, ActivityId: activity, Reaction: reaction.ReactRequest_Reaction},
	)

	return
}
