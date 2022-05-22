package api

import (
	"context"

	"github.com/artemmarkaryan/fisha-facade/pkg/logy"
)

func (a API) AddInterest(ctx context.Context, user, interest int64) error {
	// stub

	logy.Log(ctx).Infoln("add interest", user, interest)

	return nil
}
