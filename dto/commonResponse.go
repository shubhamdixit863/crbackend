package dto

type CommonResponse struct {
	Data interface{} `json:"data"`
}

func NewCommonResponse(data interface{}) *CommonResponse {
	return &CommonResponse{
		Data: data,
	}
}
