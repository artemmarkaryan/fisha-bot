package api

import (
	"net/http"

	"github.com/artemmarkaryan/fisha-facade/pkg/marchy"
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

func (a API) get(path string) (*http.Response, error) {
	return http.Get(schema + a.host + path)
}
