package bot

import (
	"context"
	"fmt"

	"github.com/artemmarkaryan/fisha/bot/internal/bot/callback"
	"github.com/artemmarkaryan/fisha/bot/internal/service/reaction"
	tele "gopkg.in/telebot.v3"
)

var reactionPattern = callback.NewPattern("reac")

func (b *Bot) recommend(ctx context.Context) tele.HandlerFunc {
	return func(t tele.Context) error {
		activity, err := b.api.Recommend(ctx, t.Chat().ID)
		if err != nil {
			b.log(ctx, t, err)
			return err
		}

		r := &tele.ReplyMarkup{}
		var rows = []tele.Row{r.Row(
			tele.Btn{Text: "👍", Data: callback.MakeCallbackData(reactionPattern, reaction.Like.Code)},
			tele.Btn{Text: "👎", Data: callback.MakeCallbackData(reactionPattern, reaction.Dislike.Code)},
		)}

		r.Inline(rows...)

		err1 := t.Send(&tele.Location{Lat: activity.Lat, Lng: activity.Lon})
		err2 := t.Send(activity.Message(), r, tele.ModeHTML)

		if err1 != nil {
			b.log(ctx, t, fmt.Errorf("send location: %v", err1))
		}
		if err2 != nil {
			b.log(ctx, t, fmt.Errorf("send message: %v", err2))
		}

		return nil
	}
}
