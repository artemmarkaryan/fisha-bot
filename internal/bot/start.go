package bot

import (
	"context"

	"github.com/artemmarkaryan/fisha-facade/pkg/logy"
	tele "gopkg.in/telebot.v3"
)

func (b *Bot) start(ctx context.Context) tele.HandlerFunc {
	return func(t tele.Context) error {
		isNew, err := b.api.Login(ctx, t.Sender().ID)
		if err != nil {
			b.log(ctx, t, err)
			return err
		}

		logy.Log(ctx).Debugf("/start: user_id: #%v, isNew: %v", t.Sender().ID, isNew)

		if !isNew {
			_, err = b.bot.Send(t.Sender(), "Снова привет!")
			return err
		}

		if err = b.chooseInterests(ctx)(t); err != nil {
			return err
		}

		return err
	}
}
