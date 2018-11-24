package response

type ResponseModel struct {
	Status       bool   `json:"status"`
	Data         string `json:"data"`
	MessageKey   string `json:"message"`
	ErrorCode    int16  `json:"error_code"`
	ResponseCode int    `json:"-"`
}

type SuccessfulResponseModel struct {
	Data       interface{} `json:"data"`
	MessageKey string      `json:"message_key"`
}

type successfulResponse struct {
	Status     bool        `json:"status"`
	Data       interface{} `json:"data"`
	MessageKey string      `json:"message"`
}

type FailedResponseModel struct {
	Data         string `json:"data"`
	MessageKey   string `json:"message_key"`
	ErrorCode    int16  `json:"error_code"`
	ResponseCode int    `json:"-"`
}

type SuccessfulBulk struct {
	Result     interface{} `json:"result"`
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Page       int  `json:"page"`
	PerPage    int  `json:"per_page"`
	TotalCount int  `json:"total_count"`
	PagesCount int  `json:"pages_count"`
	HasNext    bool `json:"has_next"`
}
