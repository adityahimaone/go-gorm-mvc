package controllers

import (
	"orm-crud/models"
	"time"
)

//Request Book
type RequestBook struct {
	Title  string `json:"title"`
	Isbn   string `json:"isbn"`
	Writer string `json:"writer"`
}

//translate req book to model
func (req *RequestBook) toModel() models.Book {
	return models.Book{
		Title:  req.Title,
		Isbn:   req.Isbn,
		Writer: req.Writer,
	}
}

//Response Book
type ResponseBook struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Isbn      string    `json:"isbn"`
	Writer    string    `json:"writer"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func newResponse(modelBooks models.Book) ResponseBook {
	return ResponseBook{
		ID:        modelBooks.ID,
		Title:     modelBooks.Title,
		Isbn:      modelBooks.Isbn,
		Writer:    modelBooks.Writer,
		CreatedAt: modelBooks.CreatedAt,
		UpdatedAt: modelBooks.UpdatedAt,
	}
}

func newResponseArray(modelBooks []models.Book) []ResponseBook {
	var res []ResponseBook
	for _, v := range modelBooks {
		res = append(res, newResponse(v))
	}
	return res
}
