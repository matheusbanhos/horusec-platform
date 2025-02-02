package enums

const (
	EnvAnalyticDatabaseURI          = "HORUSEC_DATABASE_SQL_URI"
	DefaultAnalyticDatabaseURIValue = "postgresql://root:root@localhost:5432/horusec_analytic_db?sslmode=disable"
	EnvHorusecDatabaseURI           = "HORUSEC_DATABASE_HORUSEC_SQL_URI"
	DefaultHorusecDatabaseURIValue  = "postgresql://root:root@localhost:5432/horusec_db?sslmode=disable"
	SummarySuccess                  = "success"
	SummaryFailed                   = "failed"
	TotalOfTables                   = 4
)
