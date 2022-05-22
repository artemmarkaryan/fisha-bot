package bot

import (
	"context"
	"strconv"

	"github.com/artemmarkaryan/fisha-facade/pkg/logy"
	tele "gopkg.in/telebot.v3"
)

func (b Bot) callback(ctx context.Context) tele.HandlerFunc {
	return func(t tele.Context) error {
		logy.Log(ctx).Debugf("handling callback: %v", t.Callback().Data)

		if t.Callback().Data[:3] == "int" {
			return b.interestCallback(ctx)(t)
		}

		return nil
	}
}

func (b Bot) interestCallback(ctx context.Context) tele.HandlerFunc {
	return func(t tele.Context) (err error) {
		var interestID int
		interestID, err = strconv.Atoi(t.Callback().Data[3:])

		return b.api.AddInterest(ctx, t.Sender().ID, int64(interestID))
	}
}
