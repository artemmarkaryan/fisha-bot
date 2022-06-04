package bot

import (
	"context"

	tele "gopkg.in/telebot.v3"
)

func (b *Bot) setLocation(ctx context.Context) tele.HandlerFunc {
	return func(t tele.Context) (err error) {
		location := t.Message().Location

		err = b.api.UserSetLocation(ctx, t.Chat().ID, location.Lng, location.Lat)
		if err != nil {
			return
		}

		_ = t.Send("Теперь мы знаем где ты...")

		return err
	}
}
