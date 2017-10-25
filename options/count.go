package options

import "strconv"

//Count class is "count" parameter
type Count int

//Key returns key string
func (c Count) Key() string {
	return "count"
}

//Value returns key string
func (c Count) Value() string {
	return strconv.Itoa(int(c))
}
