package web

type WebResponseSuccesForGet struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

type WebResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
