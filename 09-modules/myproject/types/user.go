package types

// Username: public access everywhere
// username: private access but public in its own package
// Age: public access everywhere
// age: private access but public in its own package

type User struct {
	Username string
	Age      int
}
