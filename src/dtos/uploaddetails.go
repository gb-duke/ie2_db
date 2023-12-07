package dtos

type UploadDetails struct {
	ID           string `json:"id"`
	UploadID     string `json:"uploadid"`
	FileName     string `json:"filename"`
	FileLocation string `json:"filelocation"`
	FileTypeID   string `json:"filetypeid"`
	URL          string `json:"url"`
	Size         int64  `json:"size"`
	MD5          string `json:"md5"`
	Version      int32  `json:"version"`
}
