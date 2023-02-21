package db

import "fmt"

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mrrobot"
	dbName   = "catalogos_con_go"
)

var ConnStr = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, host, dbName)
