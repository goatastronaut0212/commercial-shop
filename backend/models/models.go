package models

type Category struct {
    Id   string   `json:"id" binding:"required"`
    Name string   `json:"name" binding:"required"`
}
