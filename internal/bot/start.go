package bot

import (
	"context"

	"github.com/artemmarkaryan/fisha-facade/pkg/logy"
	tele "gopkg.in/telebot.v3"
)

func (b Bot) start(ctx context.Context) tele.HandlerFunc {
	return func(t tele.Context) error {
		isNew, err := b.api.Login(ctx, t.Sender().ID)
		if err != nil {
			return err
		}

		logy.Log(ctx).Debugf("/start: isNew: %v", isNew)

		//if !isNew {
		//	_, err = b.bot.Send(t.Sender(), "Привет!")
		//	return err
		//}

		b.handleForUser(ctx, t.Sender().ID, func(t2 tele.Context) error {
			_, err = b.bot.Send(t2.Sender(), "this is custom handler")
			return err
		})

		return err
	}
}
