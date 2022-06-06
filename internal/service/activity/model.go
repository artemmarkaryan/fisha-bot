package activity

import (
	"strings"

	"github.com/artemmarkaryan/fisha-facade/pkg/pb/gen/api"
	"github.com/artemmarkaryan/fisha/bot/pkg/format"
)

type Activity struct {
	Id        int64
	Name      string
	CreatedAt string
	UpdatedAt string
	Address   string
	Lon       float32
	Lat       float32
	Meta      string
}

func NewActivityFromProto(message *api.ActivityMessage) (bool, Activity) {
	activity := message.GetActivity()
	return message.GetFound(), Activity{
		Id:      activity.GetId(),
		Name:    activity.GetName(),
		Address: activity.GetAddress(),
		Meta:    activity.GetMeta(),
		Lat:     activity.GetLat(),
		Lon:     activity.GetLon(),
	}
}

func (a Activity) Message() string {
	s := []string{
		format.Bold(a.Name),
		format.Bold("Адрес: ") + a.Address,
	}

	return strings.Join(s, "\n")
}
