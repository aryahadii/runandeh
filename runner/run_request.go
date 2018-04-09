package runner

// RunRequest contains info of requested run
type RunRequest struct {
	APIVersion int                 `json:"api"`
	ID         int                 `json:"id"`
	Payload    interface{}         `json:"payload"`
	CodeLang   ProgrammingLanguage `json:"codeLang"`
	Code       string              `json:"code"`
	DBMS       DBMS                `json:"dbms"`
}

// ProgrammingLanguage is a type that specifies code's language
type ProgrammingLanguage string

const (
	// LangCpp specifies C++ programming language
	LangCpp = ProgrammingLanguage("cpp")
)

// DBMS is a type that specifies a DBMS
type DBMS string

const (
	// DBPostgres specifies PostgreSQL DBMS
	DBPostgres = DBMS("postgres")
	// DBMySQL specifies MySQL DBMS
	DBMySQL = DBMS("mysql")
	// DBSqlite3 specifies Sqlite3 DBMS
	DBSqlite3 = DBMS("sqlite3")
)
