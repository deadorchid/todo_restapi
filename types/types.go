package types

type Person struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

type Todo struct {
	Title  string `json:"title"`
	Status int    `json:"status"`
	Author Person `json:"author"`
}
