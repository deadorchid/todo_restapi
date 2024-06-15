package types

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
