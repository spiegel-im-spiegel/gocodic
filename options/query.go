package options

//Query class is "query" parameter
type Query string

//Key returns key string
func (q Query) Key() string {
	return "query"
}

//Value returns key string
func (q Query) Value() string {
	return string(q)
}
