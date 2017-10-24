package options

//Text class is "text" parameter
type Text string

//Key returns key string
func (t Text) Key() string {
	return "text"
}

//Value returns key string
func (t Text) Value() string {
	return string(t)
}
