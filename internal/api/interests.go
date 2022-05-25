package api

import (
	"context"

	"github.com/artemmarkaryan/fisha-facade/pkg/marchy"
	"github.com/artemmarkaryan/fisha-facade/pkg/pb/gen/api"
)

func (a API) Interests(ctx context.Context) ([]*api.InterestsResponse_Interest, error) {
	r, err := a.get(ctx, "/interests")
	if err != nil {
		return nil, err
	}

	ir, err := marchy.Obj[*api.InterestsResponse](ctx, r.Body)
	if err != nil {
		return nil, err
	}

	return ir.GetInterest(), nil
}
