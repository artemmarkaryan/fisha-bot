package bot

import (
	"context"
	"fmt"

	"github.com/artemmarkaryan/fisha-facade/pkg/logy"
	tele "gopkg.in/telebot.v3"
)

func (b Bot) errorMessage(ctx context.Context, t tele.Context, err error) {
	if sErr := t.Send(fmt.Sprintf("Произошла внутрення ошибка. Уже бежим чинить\n\n```%s```", err)); sErr != nil {
		logy.Log(ctx).Errorf("cant send message: %w", err)
	}
}
