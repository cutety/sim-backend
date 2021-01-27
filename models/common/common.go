package common

type LoginResponse struct {
	Status int `json:"status"`
	Msg string `json:"msg"`
	Error string `json:"error"`
	Token string `json:"token"`
}

type Response struct {
	Status int `json:"status"`
	Data interface{} `json:"data"`
	Msg string `json:"msg"`
	Error string `json:"error"`
}

type DataList struct {
	Items interface{} `json:"items"`
	Total uint `json:"total"`
}

type Pagination struct {
	Page int `json:"page"`
	Limit int `json:"limit"`
	Order string `json:"order"`
}