package gocodic

import (
	"fmt"

	"github.com/spiegel-im-spiegel/gocodic/options"
	"github.com/spiegel-im-spiegel/gocodic/request"
	"github.com/spiegel-im-spiegel/gocodic/response"
)

const (
	pathCEDList  = "/v1/ced/lookup.json"
	pathCEDEntry = "/v1/ced/entries/"
)

//LookupCED kick lookup CED API in codic.jp
func LookupCED(opts *options.Options, eid options.EntryID) (*response.Response, error) {
	var path string
	if opts.Cmd() == options.CmdCEDQuery {
		path = pathCEDList
	} else {
		path = fmt.Sprintf("%s%d.json", pathCEDEntry, int(eid))
	}
	req, err := request.NewGet(path, opts.Token())
	if err != nil {
		return nil, err
	}
	opts.Export(req)
	return req.Do()
}
