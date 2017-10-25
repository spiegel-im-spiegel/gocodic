package options

import "strconv"

//EntryID class is "entry_id" parameter
type EntryID int

//Key returns key string
func (eid EntryID) Key() string {
	return "entry_id"
}

//Value returns key string
func (eid EntryID) Value() string {
	return strconv.Itoa(int(eid))
}
