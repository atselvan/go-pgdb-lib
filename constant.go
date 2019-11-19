package pgdb

const (
	dbConnectionErr       = "DB_CONNECTION_FAILED"
	dbConnectionErrStr    = "there was an error connecting to the database : %v"
	dbCloseErr            = "DB_CLOSE_FAILED"
	dbCloseErrStr         = "there was an error closing the connection to the database : %v"
	dbPingErr             = "DB_PING_FAILED"
	dbPingErrStr          = "there was an error pinging to the database : %v"
	dbQueryExecErr        = "DB_QUERY_EXEC_FAILED"
	dbQueryExecErrStr     = "there was an error while executing the query on the database : %v"
	dbRowScanErr          = "DB_ROW_SCAN_ERROR"
	dbRowScanErrStr       = "there was an error while scanning the rows : %v"
	enumNameEmptyErr      = "EMPTY_ENUM_NAME"
	enumNameEmptyErrStr   = "enum name cannot be empty"
	tableNameEmptyErr     = "TABLE_NAME_EMPTY"
	tableNameEmptyErrStr  = "table name cannot be empty"
	tableColumnErr        = "TABLE_COLUMN_ERROR"
	tableNoColumnErrStr   = "a table should have at least one column"
	tableColumnInfoErrStr = "column name and column datatype cannot be empty"
)
