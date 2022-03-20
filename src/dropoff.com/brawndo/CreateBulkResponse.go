package brawndo

type CreateBulkResponse struct {
	Message   string `json:"message"`
	Timestamp string `json:"timpestamp"`
	Success   bool   `json:"success"`
}
