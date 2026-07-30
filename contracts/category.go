package contracts

type CategoryResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type CategoriesResponse struct {
	Categories []CategoryResponse `json:"categories"`
}
