package api

import (
	"context"

	"github.com/artemmarkaryan/fisha-facade/pkg/pb/gen/api"
)

func (a API) AckRecommendation(ctx context.Context, userID, activityID int64) (err error) {
	_, err = a.post(ctx,
		"/ack-recommendation",
		api.AckRecommendationMessage{UserId: userID, ActivityId: activityID},
	)

	return
}
