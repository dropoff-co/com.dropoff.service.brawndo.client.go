package brawndo

type DriverActionsMetaData struct {
	Id          int64  `json:"id"`
	Key         string `json:"key"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	Required    bool   `json:"required"`
}

type DriverActionsMetaResponse struct {
	Data      []*DriverActionsMetaData `json:"data"`
	Success   bool                     `json:"success"`
	Timestamp string                   `json:"timestamp"`
}
