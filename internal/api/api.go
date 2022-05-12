package api

import (
	"context"
	"net/http"

	"github.com/artemmarkaryan/fisha-facade/pkg/marchy"
	"github.com/artemmarkaryan/fisha-facade/pkg/pb/gen/api"
)

const schema = "http://"

type API struct {
	host string
}

func NewAPI(host string) API {
	return API{host: host}
}

func (a API) post(path string, data any) (*http.Response, error) {
	return http.Post(schema+a.host+path, "text/json", marchy.ForceReader(data))
}

func (a API) Login(ctx context.Context, user int64) (isNew bool, err error) {
	r, err := a.post("/login", api.UserIdRequest{UserId: user})
	if err != nil {
		return
	}

	obj, err := marchy.Obj[*api.LoginResponse](ctx, r.Body)
	if err != nil {
		return
	}

	return obj.GetNew(), nil
}
