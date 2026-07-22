package contracts

import "time"

type CategoryResponse struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type CategoriesResponse struct {
	Categories []CategoryResponse `json:"categories"`
}
