package runner

// AppResponse contains output of app
type AppResponse struct {
	StdOut string `json:"stdout"`
	StdErr string `json:"stderr"`
	Error  string `json:"error"`
}
