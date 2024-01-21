package database

var connection string
func init()  { //pertama kali dieksekusi, saat package diakses
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}