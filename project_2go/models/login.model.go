package models

import (
	"database/sql"
	"fmt"
	"project_2go/db"
	"project_2go/helpers"
)

type User struct {
	Id int `json:"id"`
	Username string `json:"username"`
}

func CheckLogin(username, password string) (bool, error) {
	var obj User
	var pwd string

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM users WHERE username = ?"

	err := con.QueryRow(sqlStatement, username).Scan(
		&obj.Id, &obj.Username, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("Username Not Found")
		return false,err
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
