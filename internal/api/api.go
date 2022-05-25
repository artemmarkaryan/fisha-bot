package api

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/artemmarkaryan/fisha-facade/pkg/logy"
	"github.com/artemmarkaryan/fisha-facade/pkg/marchy"
)

const schema = "http://"

type API struct {
	host string
}

func NewAPI(host string) API {
	return API{host: host}
}

func (a API) post(ctx context.Context, path string, data any) (*http.Response, error) {
	mws := []mw{
		a.withStatusConvertor(),
		a.withLog(ctx),
	}

	r, err := http.Post(schema+a.host+path, "text/json", marchy.ForceReader(data))

	for _, m := range mws {
		r, err = m(r, err)
	}

	return r, err
}

func (a API) get(ctx context.Context, path string) (*http.Response, error) {
	mws := []mw{
		a.withStatusConvertor(),
		a.withLog(ctx),
	}

	r, err := a.withLog(ctx)(http.Get(schema + a.host + path))

	for _, m := range mws {
		r, err = m(r, err)
	}

	return r, err
}

type mw func(response *http.Response, err error) (*http.Response, error)

func (a API) withLog(ctx context.Context) mw {
	return func(response *http.Response, err error) (*http.Response, error) {
		a.log(ctx, response, err)
		return response, err
	}
}

func (a API) withStatusConvertor() mw {
	return func(response *http.Response, err error) (*http.Response, error) {
		if response.StatusCode != http.StatusOK {
			return response, fmt.Errorf(
				"status not ok: [%v] [%v:%v]",
				response.StatusCode,
				response.Request.Method,
				response.Request.RequestURI,
			)
		}

		return response, err
	}
}

func (a API) log(ctx context.Context, response *http.Response, err error) {
	if err != nil {
		logy.Log(ctx).Errorln(err)
		return
	} else if response.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(response.Body)
		logy.Log(ctx).Errorf("[%v] [%v:%v] body: `%v`", response.StatusCode, response.Request.Method, response.Request.RequestURI, b)
		return
	}

	return
}
