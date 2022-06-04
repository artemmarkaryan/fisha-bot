package api

import (
	"context"

	"github.com/artemmarkaryan/fisha-facade/pkg/marchy"
	"github.com/artemmarkaryan/fisha-facade/pkg/pb/gen/api"
	"github.com/artemmarkaryan/fisha/bot/internal/service/activity"
)

func (a API) Recommend(ctx context.Context, userID int64) (ac activity.Activity, err error) {
	r, err := a.post(ctx, "/recommend", api.IdMessage{Id: userID})
	if err != nil {
		return
	}

	protoA, err := marchy.Obj[*api.ActivityMessage](ctx, r.Body)
	if err != nil {
		return
	}

	return activity.NewActivityFromProto(protoA), nil
}
