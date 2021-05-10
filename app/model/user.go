package model

// User data structure (username, password)
var User map[string]string

func init() {
	// Our users
	User = make(map[string]string)
	User["Bob"] = "bobobob"
	User["Doe"] = "dododoe"
}
