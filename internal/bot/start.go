package bot

import (
	"context"
	"fmt"

	"github.com/artemmarkaryan/fisha-facade/pkg/logy"
	tele "gopkg.in/telebot.v3"
)

func (b Bot) start(ctx context.Context) tele.HandlerFunc {
	return func(t tele.Context) error {
		isNew, err := b.api.Login(ctx, t.Sender().ID)
		if err != nil {
			return err
		}

		logy.Log(ctx).Debugf("/start: isNew: %v", isNew)

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

func (b Bot) chooseInterests(ctx context.Context) tele.HandlerFunc {
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
				Data: fmt.Sprintf("int%v", interest.Id),
			}))
		}

		r.Inline(rows...)

		_, _ = b.bot.Send(t.Sender(), "Выбери свои интересы", r)
		//b.handleForUser(ctx, t.Sender().ID, b.chooseInterestsCB(ctx))

		return nil
	}
}

func (b Bot) chooseInterestsCB(ctx context.Context) tele.HandlerFunc {
	return func(t tele.Context) error {
		logy.Log(ctx).Infof("interest chosen; user: %v, interest: %v", t.Sender().ID, t.Text())

		return b.chooseInterests(ctx)(t)
	}
}
