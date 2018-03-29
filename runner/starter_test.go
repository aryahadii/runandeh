package runner

import "testing"

func TestRun(t *testing.T) {
	reqTable := []*RunRequest{
		&RunRequest{
			ID:                 0,
			Payload:            nil,
			CodeLang:           LangCpp,
			AppCode:            "",
			DB:                 DBPostrges,
			DBValidatorQueries: nil,
		},
	}

	InitRunner()

	for _, req := range reqTable {
		Run(req)
	}
}
