package gocodic

import (
	"fmt"

	"github.com/spiegel-im-spiegel/gocodic/options"
	"github.com/spiegel-im-spiegel/gocodic/request"
	"github.com/spiegel-im-spiegel/gocodic/response"
)

const (
	pathProjList = "/v1/user_projects.json"
	pathProj     = "/v1/user_projects/"
)

//ReferProjects kick refer projects API in codic.jp
func ReferProjects(opts *options.Options, pid options.ProjectID) (*response.Response, error) {
	var path string
	if opts.Cmd() == options.CmdProjLst {
		path = pathProjList
	} else {
		path = fmt.Sprintf("%s%d.json", pathProj, int(pid))
	}
	req, err := request.NewGet(path, opts.Token())
	if err != nil {
		return nil, err
	}
	opts.Export(req)
	return req.Do()
}
