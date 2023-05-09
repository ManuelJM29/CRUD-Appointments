package responses

type Response struct {
	Msg    string      `json:"msg,omitempty"`
	Data   interface{} `json:"data,omitempty"`
	Status int         `json:"status"`
}
