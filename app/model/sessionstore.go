package model

// Session store data structure (sid, username)
var Session map[string]string

func init() {
	// Session store is empty after server boot up
	Session = make(map[string]string)
}
