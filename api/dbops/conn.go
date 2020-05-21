package dbops

import "database/sql"

var (
	dbConn *sql.DB
	err error
)

func init() {
	dbConn, err = sql.Open("mysql","root:Itzler.666(192.168.10.99:3306)/video_server?charset=utf8mb4")
	if err != nil {
		panic(err.Error())
	}
}