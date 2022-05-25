package bot

import (
	"context"

	"github.com/artemmarkaryan/fisha-facade/pkg/logy"
	tele "gopkg.in/telebot.v3"
)

func (b *Bot) handleForUser(ctx context.Context, uID int64, h tele.HandlerFunc) {
	logy.Log(ctx).Debugf("registered handler for user #%v", uID)
	b.individualHandlers[uID] = h
}

func (b *Bot) defaultHandler(ctx context.Context) tele.HandlerFunc {
	return func(t tele.Context) error {
		logy.Log(ctx).Debugf("default handler triggered for user #%v", t.Sender().ID)

		if h, ok := b.individualHandlers[t.Sender().ID]; ok {
			delete(b.individualHandlers, t.Sender().ID)
			return h(t)
		}

		_, _ = b.bot.Send(t.Sender(), "Нипонял(")
		return nil
	}
}
