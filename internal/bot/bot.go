package bot

import (
	"context"
	"time"

	"github.com/artemmarkaryan/fisha/bot/internal/api"
	tele "gopkg.in/telebot.v3"
)

type Bot struct {
	bot *tele.Bot
	api api.API

	individualHandlers map[int64]tele.HandlerFunc
}

type BotConfig struct {
	Token   string
	API     api.API
	Timeout time.Duration
}

func NewBot(ctx context.Context, cfg BotConfig) (*Bot, error) {
	pref := tele.Settings{
		Token: cfg.Token,
		Poller: &tele.LongPoller{
			Timeout: cfg.Timeout,
		},
	}

	telebot, err := tele.NewBot(pref)
	if err != nil {
		return nil, err
	}

	b := &Bot{
		api:                cfg.API,
		bot:                telebot,
		individualHandlers: map[int64]tele.HandlerFunc{},
	}

	b.Register(ctx)

	return b, nil
}

func (b *Bot) Register(ctx context.Context) {
	b.bot.Handle("/start", b.start(ctx))
	b.bot.Handle("/forget", b.forget(ctx))
	b.bot.Handle(tele.OnCallback, b.callback(ctx))
	b.bot.Handle(tele.OnText, b.defaultHandler(ctx))
}

func (b Bot) Start() { b.bot.Start() }
