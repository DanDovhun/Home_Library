package pages

import "Home_Library/database"

type Radios struct {
	Value       string
	Name        string
	Placeholder string
}

type HomePage struct {
	Title   string
	Buttons []Radios
	Books   database.Book
}
