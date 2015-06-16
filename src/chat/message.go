package chat

type User struct {
	Name  string
	Email string
}

type Message struct {
	MType    string
	Content  string
	Userinfo *User
	Time     string
}
