package bot

import (
	"context"
	"errors"

	"github.com/artemmarkaryan/fisha/bot/internal/bot/callback"
	"github.com/artemmarkaryan/fisha/bot/internal/service/reaction"
	tele "gopkg.in/telebot.v3"
)

func (b *Bot) reactionCallback(ctx context.Context) tele.HandlerFunc {
	return func(t tele.Context) (err error) {
		_, reactionCode, err := callback.Parse(t.Callback().Data)
		if err != nil {
			return
		}

		r, ok := reaction.Mapping[reactionCode]
		if !ok {
			return errors.New("unknown reaction")
		}

		_ = t.Send("Поставили " + r.Emoji)

		return
	}
}
