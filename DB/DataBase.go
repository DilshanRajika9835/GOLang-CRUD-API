package DB

import
	// This should be used in conjunction with a SQL driver and this package will provide a generic interface around SQL databases .

"database/sql"

func GetConnection() *sql.DB {
	db, _ := sql.Open("mysql", "root:ijse@tcp(127.0.0.1:3306)/GOLand")
	return db
}
