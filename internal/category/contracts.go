package category

type createRequest struct {
	Name string `json:"name" validate:"required,max=255"`
}
