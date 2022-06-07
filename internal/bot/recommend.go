package bot

import (
	"context"
	"fmt"

	"github.com/artemmarkaryan/fisha/bot/internal/bot/callback"
	"github.com/artemmarkaryan/fisha/bot/internal/service/reaction"
	tele "gopkg.in/telebot.v3"
)

var reactionPattern = callback.NewPattern("reac")
var recommendPattern = callback.NewPattern("rcmd")

func (b *Bot) recommend(ctx context.Context) tele.HandlerFunc {
	return func(t tele.Context) error {
		hl, err := b.api.UserHasLocation(ctx, t.Chat().ID)
		if err != nil {
			b.log(ctx, t, err)
			return err
		}

		if !hl {
			_ = t.Send("А мы не знаем, где вы 😣\nПожалуйста, отправьте геопозиацию, вокруг которой нам искать места")
			return nil
		}

		found, activity, err := b.api.Recommend(ctx, t.Chat().ID)
		if err != nil {
			b.log(ctx, t, err)
			return err
		}

		if !found {
			return t.Send("Нет подходящих объектов около вас 😣\nПожалуйста, отправьте другую локацию, либо выберите больше интересов")
		}

		if err = b.api.AckRecommendation(ctx, t.Chat().ID, activity.Id); err != nil {
			b.log(ctx, t, err)
			return err
		}

		r := &tele.ReplyMarkup{}
		var rows = []tele.Row{
			r.Row(
				tele.Btn{Text: "👍", Data: callback.MakeCallbackData(reactionPattern, reaction.LikeReaction.Code)},
				tele.Btn{Text: "👎", Data: callback.MakeCallbackData(reactionPattern, reaction.DislikeReaction.Code)},
			),
			r.Row(tele.Btn{Text: "Ещё", Data: callback.MakeCallbackData(recommendPattern, "")}),
		}

		r.Inline(rows...)

		err1 := t.Send(&tele.Location{Lat: activity.Lat, Lng: activity.Lon})
		err2 := t.Send(activity.Message(ctx), r, tele.ModeHTML)

		if err1 != nil {
			b.log(ctx, t, fmt.Errorf("send location: %v", err1))
		}
		if err2 != nil {
			b.log(ctx, t, fmt.Errorf("send message: %v", err2))
		}

		return nil
	}
}
