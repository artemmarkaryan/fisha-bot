package bot

import (
	"context"

	tele "gopkg.in/telebot.v3"
)

func (b *Bot) forget(ctx context.Context) tele.HandlerFunc {
	return func(t tele.Context) error {
		if err := b.api.Forget(ctx, t.Sender().ID); err != nil {
			b.log(ctx, t, err)
			return err
		}

		_ = t.Send("Мы про вас забыли")

		return nil
	}
}
