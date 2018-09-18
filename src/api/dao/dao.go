package dao

import(
	"apiGo/src/api/domain"
	"database/sql";
)
var queryUser = "Select * from User where id= ?"
func GetUser(key string)*domain.User{
	return getUser(key,queryUser)
}

func getUser(key string , query string)*domain.User{
	db, err := sql.Open("mysql", "root:armageon1@tcp(127.0.0.1:3306)/apitest")
	if err != nil {
		panic(err.Error())
	}
	var user =domain.User{}
	err = db.QueryRow(query, key).Scan(&user)

	defer db.Close()
	return &user
}