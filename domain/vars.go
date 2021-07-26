package domain

var PERMIT = map[string][]string{
	"owner": {"/", "/login", "/register", "/books"},
	"user":  {"/login"},
}
