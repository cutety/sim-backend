package common

const (
	IS_MATCHED = 1
	NOT_MATCHED = 0
)

type LoginResponse struct {
	Status int `json:"status"`
	Msg string `json:"msg"`
	Error string `json:"error"`
	Token string `json:"token"`
}

type Response struct {
	Status int `json:"status" example:"10001"`
	Data interface{} `json:"data" example:""`
	Msg string `json:"msg" example:"OK"`
	Error string `json:"error" example:""`
}

type DataList struct {
	Items interface{} `json:"items"`
	Total int64 `json:"total"`
}

type Pagination struct {
	Page int `json:"page"`
	Limit int `json:"limit"`
	Order string `json:"order"`
}