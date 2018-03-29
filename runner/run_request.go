package runner

// RunRequest contains info of requested run
type RunRequest struct {
	ID                 int         `json:"id"`
	Payload            interface{} `json:"payload"`
	CodeLang           string      `json:"codeLang"`
	AppCode            string      `json:"appCode"`
	DB                 string      `json:"db"`
	DBValidatorQueries []string    `json:"dbValidatorQueries"`
}

const (
	LangCpp = "cpp"

	DBPostrges = "postgres"
)
