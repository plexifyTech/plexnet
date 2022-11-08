package error

type Response struct {
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
}

func NewResponse(code int, err error) Response {
	return Response{
		Error:      err.Error(),
		StatusCode: code,
	}
}
