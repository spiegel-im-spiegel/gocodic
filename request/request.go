package request

import (
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/pkg/errors"
	"github.com/spiegel-im-spiegel/gocodic/response"
)

var (
	//ErrRequest is error in request package
	ErrRequest = errors.New("invalid parameter")
)

const (
	errMsgDo = "error in request.Request.Do() function"
)

//Request class is parameters for request
type Request struct {
	path  string
	token string
	data  map[string]string
}

//New returns Request instance
func New(path, token string) (*Request, error) {
	if len(path) == 0 || len(token) == 0 {
		return nil, ErrRequest
	}
	return &Request{path: path, token: token, data: make(map[string]string)}, nil
}

//Add sets key-value data
func (r *Request) Add(key, value string) {
	if r == nil {
		return
	}
	r.data[key] = value
}

//Do return respons from codic service
func (r *Request) Do() (*response.Response, error) {
	if r == nil {
		return nil, ErrRequest
	}
	buffer := new(bytes.Buffer)
	writer := multipart.NewWriter(buffer)
	for key, value := range r.data {
		//fmt.Printf("\"%s\" = \"%s\"\n", key, value)
		writer.WriteField(key, value)
	}
	writer.Close() //flush

	req, err := http.NewRequest("POST", "https://api.codic.jp"+r.path, buffer)
	if err != nil {
		return nil, errors.Wrap(err, errMsgDo)
	}
	req.Header.Add("Authorization", "Bearer "+r.token)
	req.Header.Add("Content-Type", "multipart/form-data; boundary="+writer.Boundary())

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, errors.Wrap(err, errMsgDo)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, errMsgDo)
	}
	return response.New(resp.StatusCode, resp.Status, body), nil
}
