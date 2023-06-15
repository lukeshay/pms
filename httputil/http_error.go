package httputil

type HTTPErrorError struct {
	Message string `json:"message"`
}

type HTTPError struct {
	Error HTTPErrorError `json:"error"`
}
