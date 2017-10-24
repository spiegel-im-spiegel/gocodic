package options

import (
	"github.com/spiegel-im-spiegel/gocodic/request"
)

//Option interface class for codic parameter
type Option interface {
	Key() string
	Value() string
}

//Options class is codic parameter list
type Options struct {
	options []Option
	token   string
}

//NewOptions returns Options instance
func NewOptions(token string) (*Options, error) {
	if len(token) == 0 {
		return nil, ErrNoAccessToken
	}
	return &Options{token: token}, nil
}

//Add append Option instance
func (opts *Options) Add(o Option) {
	if opts != nil {
		opts.options = append(opts.options, o)
	}
}

//Token returns access token
func (opts *Options) Token() string {
	return opts.token
}

//Export append Option instance
func (opts *Options) Export(req *request.Request) {
	if opts == nil {
		return
	}
	for _, opt := range opts.options {
		req.Add(opt.Key(), opt.Value())
	}
}
