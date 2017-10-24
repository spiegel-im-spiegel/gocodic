package options

import "errors"

var (
	//ErrOption is error in options package
	ErrOption = errors.New("invalid option")
	//ErrNoAccessToken is error "no access token"
	ErrNoAccessToken = errors.New("require access token")
)
