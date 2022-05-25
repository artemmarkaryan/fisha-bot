package bot

import (
	"context"
	"fmt"

	"github.com/artemmarkaryan/fisha/bot/internal/bot/callback"
	tele "gopkg.in/telebot.v3"
)

var interestsPattern = callback.NewPattern("intt")

func (b *Bot) chooseInterests(ctx context.Context) tele.HandlerFunc {
	return func(t tele.Context) error {
		interests, err := b.api.Interests(ctx)
		if err != nil {
			return err
		}

		r := &tele.ReplyMarkup{}
		var rows []tele.Row
		for _, interest := range interests {
			rows = append(rows, r.Row(tele.Btn{
				Text: interest.Name,
				Data: callback.MakeCallbackData(interestsPattern, fmt.Sprint(interest.Id)),
			}))
		}

		r.Inline(rows...)

		_, _ = b.bot.Send(t.Sender(), "Выбери свои интересы", r)

		return nil
	}
}
