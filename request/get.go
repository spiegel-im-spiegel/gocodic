package request

import (
	"net/url"

	"github.com/spiegel-im-spiegel/gocodic/response"
)

//Get class is parameters for Get request
type Get struct {
	path   string
	token  string
	values url.Values
}

//NewGet returns Get instance
func NewGet(path, token string) (Request, error) {
	if len(path) == 0 || len(token) == 0 {
		return nil, ErrRequest
	}
	return &Get{path: path, token: token, values: url.Values{}}, nil
}

//Add sets key-value data
func (r *Get) Add(key, value string) {
	if r == nil {
		return
	}
	//fmt.Printf("\"%s\" = \"%s\"\n", key, value)
	r.values.Add(key, value)
}

//Do return respons from codic service
func (r *Get) Do() (*response.Response, error) {
	if r == nil {
		return nil, ErrRequest
	}
	params := r.values.Encode()
	url := "https://api.codic.jp" + r.path
	if len(params) > 0 {
		url += "?" + params
	}
	//fmt.Println(url)

	return requestDo(methodGet, url, r.token, "", nil)
}
