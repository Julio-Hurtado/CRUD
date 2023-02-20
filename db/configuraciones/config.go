package db

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mrrobot"
	dbName   = "catalogo_con_go"
)

const ConnStr = "postgres://postgres:mrrobot@localhost/catalogos_con_go?sslmode=disable"

// var ConnStr = fmt.Sprintf("host=%s port=%d user=%s "+
// 	"password=%s dbname=%s sslmode=disable",
// 	host, port, user, password, dbName)
