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

func NewActivityFromProto(message *api.ActivityMessage) Activity {
	return Activity{
		Id:      message.GetId(),
		Name:    message.GetName(),
		Address: message.GetAddress(),
		Meta:    message.GetMeta(),
		Lat:     message.GetLat(),
		Lon:     message.GetLon(),
	}
}

func (a Activity) Message() string {
	s := []string{
		format.Bold(a.Name),
		format.Bold("Адрес: ") + a.Address,
	}

	return strings.Join(s, "\n")
}
