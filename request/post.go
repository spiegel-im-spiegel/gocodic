package request

import (
	"bytes"
	"mime/multipart"

	"github.com/spiegel-im-spiegel/gocodic/response"
)

//Post class is parameters for Post request
type Post struct {
	path  string
	token string
	data  map[string]string
}

//NewPost returns Post instance
func NewPost(path, token string) (*Post, error) {
	if len(path) == 0 || len(token) == 0 {
		return nil, ErrRequest
	}
	return &Post{path: path, token: token, data: make(map[string]string)}, nil
}

//Add sets key-value data
func (r *Post) Add(key, value string) {
	if r == nil {
		return
	}
	r.data[key] = value
}

//Do return respons from codic service
func (r *Post) Do() (*response.Response, error) {
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

	return requestDo(methodPost, "https://api.codic.jp"+r.path, r.token, writer.Boundary(), buffer)
}
