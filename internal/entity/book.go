package entity

type Book struct {
	ID     int64  `json:"id"`
	Tittle string `json:"tittle"`
	Author string `json:"author"`
}
