package main

import (
	"context"
	"os"

	"github.com/artemmarkaryan/fisha-facade/pkg/logy"
	"github.com/artemmarkaryan/fisha/bot/internal/api"
	"github.com/artemmarkaryan/fisha/bot/internal/bot"
	"github.com/artemmarkaryan/fisha/bot/internal/config"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var ctx = context.Background()

	ctx = initLogger(ctx)

	b, err := bot.NewBot(ctx, bot.BotConfig{
		Token:   os.Getenv("TELEGRAM_TOKEN"),
		API:     api.NewAPI(os.Getenv("SERVER_HOST")),
		Timeout: config.BotTimeout,
	})

	if err != nil {
		logy.Log(ctx).Errorf("failed to create bot: %v", err)
		return
	}

	b.Start()
}

func initLogger(ctx context.Context) context.Context {
	return logy.New(ctx)
}
