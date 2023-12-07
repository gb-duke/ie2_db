package dtos

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	CreatedOn string `json:"createdon"`
	UpdatedOn string `json:"updatedon"`
	DeletedOn string `json:"deletedon"`
}
