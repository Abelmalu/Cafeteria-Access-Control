package models

type Student struct {
	IdCard     string `json:"id_card" db:"id_card"`
	FirstName  string `json:"first_name" db:"first_name"`
	MiddleName string `json:"middle_name" db:"middle_name"`
	LastName   string `json:"last_name" db:"last_name"`
	BatchId    string `json:"batch_id" db:"batch_id"`
	RFIDTag    string `json:"rfid_tag" db:"rfid_tag"`
	ImageURL   string `json:"image_url" db:"image_url"` // path or URL to the image
}
