package models

type Product struct {
	ID          int    `json:"id" binding:"required,min=1,max=100000"`
	Name        string `json:"name" binding:"required,min=1,max=150"`
	Category    string `json:"category" binding:"required,min=1,max=50"`
	Description string `json:"description" binding:"required,min=1,max=500"`
	Brand       string `json:"brand" binding:"required,min=1,max=50"`
}
