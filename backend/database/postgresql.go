package database 

const (
    username = "postgres"
    password = "postgres"
    database = "mkjeandb"
    host = "localhost"
    port = "5432"
)

// var urlExample = "postgres://username:password@localhost:5432/database_name
const CONNECT_STR = "postgres://" + username + ":" + password + "@" + host + ":" + port + "/" + database + "?sslmode=disable"
 
