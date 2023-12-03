package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = "5432"
	user   = "postgres"
	pwd    = "emi"
	dbname = "loft"
)

var cn *sql.DB

func Open() {
	conexion := "host=" + host + " port=" + port + " user=" + user + " pwd=" + pwd + " dbname=" + dbname + " sslmode=disable"
	db, err := sql.Open("postgres", conexion)
	if err != nil {
		panic(err.Error())
	}
	// fmt.Fprintln(os.Stdout, []any{"Se conecto a...", db}...)
	cn = db
}
func Query(consulta string, args ...interface{}) (*sql.Rows, error) {
	rows, err := cn.Query(consulta, args...)
	if err != nil {
		panic(err.Error())
	}
	return rows, err
}
func Close() {
	cn.Close()
}
