package request

import (
	"bytes"
	"io"
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
func NewPost(path, token string) (Request, error) {
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
	body, boundary := r.makeBody()

	return requestDo(methodPost, "https://api.codic.jp"+r.path, r.token, boundary, body)
}

func (r *Post) makeBody() (io.Reader, string) {
	buffer := new(bytes.Buffer)
	writer := multipart.NewWriter(buffer)
	defer writer.Close()
	for key, value := range r.data {
		//fmt.Printf("\"%s\" = \"%s\"\n", key, value)
		writer.WriteField(key, value)
	}
	return buffer, writer.Boundary()
}
