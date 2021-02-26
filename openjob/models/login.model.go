package models

import (
	"database/sql"
	"fmt"
	"openjob/db"
	"openjob/helpers"
)

type User struct {
	Nim      int    `json:"nim`
	Password string `json:"password`
}

func CheckLogin(Nim, password string) (bool, error) {

	var obj User
	var pwd string

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM users WHERE nim = ?"

	err := con.QueryRow(sqlStatement, Nim).Scan(
		&obj.Nim, &obj.Password, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("Nim Not Found")
		return false, err
	}
	if err != nil {
		fmt.Println("Querry Error")
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Hash and Password doesn't match")
		return false, err
	}
	return true, nil
}
