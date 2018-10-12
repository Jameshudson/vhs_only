package repositories

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"os"
)

type FilmRepositories struct {
	
}

func (fr *FilmRepositories) GetAllFilms(limit int) {

	db, err := sql.Open("mysql", os.Getenv("database_user") + ":" + os.Getenv("database_pass") + "@/" + os.Getenv("database_name"))
	
	if err != nil{
		fmt.Println("Failed")
	}

	fmt.Println(db)
}