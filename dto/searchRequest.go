package dto

type SearchRequest struct {
	Page     uint   `json:"page"`
	Limit    uint   `json:"limit"`
	Search   string `json:"search"`
	Location string `json:"location"`
	Category string `json:"category"`
}
