package gocodic

import (
	"github.com/spiegel-im-spiegel/gocodic/options"
	"github.com/spiegel-im-spiegel/gocodic/request"
	"github.com/spiegel-im-spiegel/gocodic/response"
)

const (
	pathTranslate = "/v1/engine/translate.json"
)

//Translate kick transration API in codic.jp
func Translate(opts *options.Options) (*response.Response, error) {
	req, err := request.New(pathTranslate, opts.Token())
	if err != nil {
		return nil, err
	}
	opts.Export(req)
	return req.Do()
}
