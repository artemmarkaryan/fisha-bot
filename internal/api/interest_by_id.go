package api

import (
	"context"

	"github.com/artemmarkaryan/fisha-facade/pkg/marchy"
	"github.com/artemmarkaryan/fisha-facade/pkg/pb/gen/api"
)

func (a API) InterestById(ctx context.Context, id int64) (name string, err error) {
	r, err := a.post(ctx, "/interest-by-id", api.IdMessage{Id: id})
	if err != nil {
		return
	}

	ir, err := marchy.Obj[*api.StringMessage](ctx, r.Body)
	if err != nil {
		return
	}

	name = ir.GetS()

	return
}
