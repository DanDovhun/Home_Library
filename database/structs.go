package database

type Book struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Language string `json:"language"`
	Genre    string `json:"genre"`
	Room     string `json:"room"`
	Shelf    string `json:"shelf"`
}

type Books struct {
	Book []Book `json:"books"`
}
