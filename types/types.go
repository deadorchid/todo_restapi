package types

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Person struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

type Todo struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status int    `json:"status"`
	Author Person `json:"author"`
}
