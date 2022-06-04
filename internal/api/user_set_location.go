package api

import (
	"context"

	"github.com/artemmarkaryan/fisha-facade/pkg/pb/gen/api"
)

func (a API) UserSetLocation(ctx context.Context, userID int64, lon, lat float32) (err error) {
	_, err = a.post(ctx,
		"/user/set-location",
		api.SetLocationMessage{UserId: userID, Lon: lon, Lat: lat},
	)

	return
}
