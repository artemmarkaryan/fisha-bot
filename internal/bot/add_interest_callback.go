package bot

import (
	"context"
	"fmt"
	"strconv"

	"github.com/artemmarkaryan/fisha/bot/internal/bot/callback"
	tele "gopkg.in/telebot.v3"
)

func (b *Bot) addInterestCallback(ctx context.Context) tele.HandlerFunc {
	return func(t tele.Context) (err error) {
		var interestID int
		_, interestIdString, err := callback.Parse(t.Callback().Data)
		if err != nil {
			return
		}

		interestID, err = strconv.Atoi(interestIdString)
		if err != nil {
			return
		}

		isNew, err := b.api.AddInterest(ctx, t.Sender().ID, int64(interestID))
		if err != nil {
			return
		}

		interestName, err := b.api.InterestById(ctx, int64(interestID))
		if err != nil {
			return
		}

		if isNew {
			_ = t.Send(fmt.Sprintf("✅ Интерес %q добавлен", interestName))
		} else {
			_ = t.Send(fmt.Sprintf("👌 Интерес %q у вас уже есть", interestName))
		}

		return
	}
}
