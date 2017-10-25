package request

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"github.com/spiegel-im-spiegel/gocodic/response"
)

var (
	//ErrRequest is error in request package
	ErrRequest = errors.New("invalid parameter")
)

const (
	errMsgDo   = "error in request.requestDo() function"
	methodPost = "POST"
	methodGet  = "GET"
)

//Request interface for Post/Get classes
type Request interface {
	Add(string, string)
	Do() (*response.Response, error)
}

func requestDo(method, url, token, boundary string, body io.Reader) (*response.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, errors.Wrap(err, errMsgDo)
	}
	if len(token) > 0 {
		req.Header.Add("Authorization", "Bearer "+token)
	}
	if len(boundary) > 0 {
		req.Header.Add("Content-Type", "multipart/form-data; boundary="+boundary)
	}

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, errors.Wrap(err, errMsgDo)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, errMsgDo)
	}
	return response.New(resp.StatusCode, resp.Status, respBody), nil
}
