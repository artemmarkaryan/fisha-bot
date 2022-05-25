package bot

import (
	"context"
	"errors"
	"time"

	"github.com/artemmarkaryan/fisha-facade/pkg/logy"
	"github.com/artemmarkaryan/fisha/bot/internal/api"
	"github.com/artemmarkaryan/fisha/bot/internal/bot/callback"
	tele "gopkg.in/telebot.v3"
)

type Bot struct {
	debug bool

	bot *tele.Bot
	api api.API

	individualHandlers map[int64]tele.HandlerFunc
	callbackHandlers   map[callback.Pattern]tele.HandlerFunc
}

type Config struct {
	Token   string
	API     api.API
	Timeout time.Duration
	Debug   bool
}

func NewBot(ctx context.Context, cfg Config) (*Bot, error) {
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
		debug:              cfg.Debug,
		individualHandlers: map[int64]tele.HandlerFunc{},
		callbackHandlers:   map[callback.Pattern]tele.HandlerFunc{},
	}

	b.Register(ctx)

	return b, nil
}

func (b *Bot) Register(ctx context.Context) {

	// user commands
	b.Handle(ctx, "/start", b.start(ctx))
	b.Handle(ctx, "/choose_interests", b.chooseInterests(ctx))

	// callbacks
	b.HandleCallback(ctx, interestsPattern, b.addInterestCallback(ctx))

	// helpers
	b.Handle(ctx, "/forget", b.forget(ctx))
	b.Handle(ctx, "/resetKeyboard", b.resetKeyboard(ctx))

	// defaults
	b.Handle(ctx, tele.OnText, b.defaultHandler(ctx))
	b.Handle(ctx, tele.OnCallback, b.callback(ctx))

}

func (b *Bot) HandleCallback(_ context.Context, pattern [4]byte, h tele.HandlerFunc) {
	b.callbackHandlers[pattern] = h
}

func (b *Bot) callback(ctx context.Context) tele.HandlerFunc {
	return func(t tele.Context) (err error) {
		logy.Log(ctx).Debugf("handling callback: %v", t.Callback())

		if len(t.Callback().Data) < callback.PatternLen {
			logy.Log(ctx).Errorf("too short callback data: %v", t.Callback())
		}

		pt, _, err := callback.Parse(t.Callback().Data)
		if err != nil {
			return
		}

		handler, ok := b.callbackHandlers[pt]
		if !ok {
			b.log(ctx, t, errors.New("unknown callback pattern: "+string(pt[:])))
			return
		}

		return handler(t)
	}
}

func (b *Bot) Handle(ctx context.Context, ep interface{}, h tele.HandlerFunc, m ...tele.MiddlewareFunc) {
	wh := func(t tele.Context) error {
		err := h(t)
		if err != nil {
			logy.Log(ctx).Errorf("%v err: %v", ep, err.Error())
		}

		return err
	}

	b.bot.Handle(ep, wh, m...)
}

func (b *Bot) log(ctx context.Context, t tele.Context, err error) {
	logy.Log(ctx).Errorln(err)

	if b.debug == false {
		return
	}

	_ = t.Send("error: " + err.Error())
}

func (b *Bot) Start() { b.bot.Start() }
