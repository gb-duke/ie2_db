package dtos

type DataUpload struct {
	ID     string `json:"id"`
	UserID string `json:"userid"`
	// UploadDetails UploadDetails `json:"uploaddetails"`
	CreatedOn string `json:"createdon"`
	UpdatedOn string `json:"updatedon"`
	DeletedOn string `json:"deletedon"`
	UpdatedBy string `json:"updatedby"`
}
