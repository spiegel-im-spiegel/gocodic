package options

import (
	"github.com/spiegel-im-spiegel/gocodic/request"
)

//CmdID is kind of commands
type CmdID int

const (
	//CmdUnknown is unknown command
	CmdUnknown CmdID = iota
	//CmdTrans is transration command
	CmdTrans
	//CmdProjLst is listing project command
	CmdProjLst
	//CmdProj is project info command
	CmdProj
	//CmdCEDQuery is CED lookup command
	CmdCEDQuery
	//CmdCED is CED info command
	CmdCED
)

//Option interface class for codic parameter
type Option interface {
	Key() string
	Value() string
}

//Options class is codic parameter list
type Options struct {
	cmd     CmdID
	options []Option
	token   string
}

//NewOptions returns Options instance
func NewOptions(cmd CmdID, token string) (*Options, error) {
	if len(token) == 0 {
		return nil, ErrNoAccessToken
	}
	return &Options{cmd: cmd, token: token}, nil
}

//Add append Option instance
func (opts *Options) Add(o Option) {
	if opts != nil {
		opts.options = append(opts.options, o)
	}
}

//Cmd returns access CmdID
func (opts *Options) Cmd() CmdID {
	return opts.cmd
}

//Token returns access token
func (opts *Options) Token() string {
	return opts.token
}

//Export append Option instance
func (opts *Options) Export(req request.Request) {
	if opts == nil {
		return
	}
	for _, opt := range opts.options {
		switch opt.(type) {
		case Casing:
			if opts.cmd == CmdTrans {
				req.Add(opt.Key(), opt.Value())
			}
		case AcronymStyle:
			if opts.cmd == CmdTrans {
				req.Add(opt.Key(), opt.Value())
			}
		case Text:
			if opts.cmd == CmdTrans {
				req.Add(opt.Key(), opt.Value())
			}
		case ProjectID:
			if opts.cmd == CmdTrans {
				req.Add(opt.Key(), opt.Value())
			}
		case Query:
			if opts.cmd == CmdCEDQuery {
				req.Add(opt.Key(), opt.Value())
			}
		case Count:
			if opts.cmd == CmdCEDQuery {
				req.Add(opt.Key(), opt.Value())
			}
		default:
		}
	}
}
