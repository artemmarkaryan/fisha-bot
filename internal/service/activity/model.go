package activity

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/artemmarkaryan/fisha-facade/pkg/logy"
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

type Meta struct {
	Id     string `json:"id"`
	Url    string `json:"url"`
	Name   string `json:"name"`
	Phones []struct {
		Type      string `json:"type"`
		Formatted string `json:"formatted"`
	} `json:"Phones"`
	Address string `json:"address"`
}

func (a Activity) Message(ctx context.Context) string {
	s := []string{
		format.Bold(a.Name),
		format.Bold("Адрес: ") + a.Address,
	}

	var m = new(Meta)
	err := json.Unmarshal([]byte(a.Meta), m)
	if err != nil {
		logy.Log(ctx).Errorf("cant parse meta: %v", err)
	}

	if m.Phones != nil {
		for _, phone := range m.Phones {
			s = append(s, "📞 "+phone.Formatted)
		}
	}

	if m.Url != "" {
		s = append(s, format.Link("\nСайт", m.Url))
	}

	return strings.Join(s, "\n")
}
