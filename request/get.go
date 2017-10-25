package request

import (
	"net/url"

	"github.com/spiegel-im-spiegel/gocodic/response"
)

//Get class is parameters for Get request
type Get struct {
	path  string
	token string
	data  map[string]string
}

//NewGet returns Get instance
func NewGet(path, token string) (*Get, error) {
	if len(path) == 0 || len(token) == 0 {
		return nil, ErrRequest
	}
	return &Get{path: path, token: token, data: make(map[string]string)}, nil
}

//Add sets key-value data
func (r *Get) Add(key, value string) {
	if r == nil {
		return
	}
	r.data[key] = value
}

//Do return respons from codic service
func (r *Get) Do() (*response.Response, error) {
	if r == nil {
		return nil, ErrRequest
	}
	values := url.Values{}
	for key, value := range r.data {
		//fmt.Printf("\"%s\" = \"%s\"\n", key, value)
		values.Add(key, value)
	}
	params := values.Encode()
	url := "https://api.codic.jp" + r.path
	if len(params) > 0 {
		url += "?" + params
	}
	//fmt.Println(url)

	return requestDo(methodGet, url, r.token, "", nil)
}
