package bot

import (
	"context"

	tele "gopkg.in/telebot.v3"
)

func (b *Bot) resetKeyboard(ctx context.Context) tele.HandlerFunc {
	return func(t tele.Context) error {
		r := &tele.ReplyMarkup{OneTimeKeyboard: true}
		r.Reply(r.Row(tele.Btn{Text: "Жмак"}))

		_ = t.Send("Нажми чтобы сбросить", r)

		b.handleForUser(ctx, t.Sender().ID, func(t2 tele.Context) error {
			_ = t2.Send("Сбросили")
			return nil
		})

		return nil
	}
}
