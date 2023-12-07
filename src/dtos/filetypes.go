package dtos

type FileTypes struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedOn   string `json:"createdat"`
	UpdatedOn   string `json:"updatedon"`
	DeletedOn   string `json:"deletedon"`
}
