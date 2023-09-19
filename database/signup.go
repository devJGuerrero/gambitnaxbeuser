package database

import (
	"fmt"
	"github.com/devJGuerrero/gambitnaxbeuser/models"
	"github.com/devJGuerrero/gambitnaxbeuser/tools"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Init func sign up")
	Failure = Connect()
	if Failure != nil {
		return Failure
	}
	defer DB.Close()

	sql := fmt.Sprintf("INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('%s', '%s', '%s')",
		sig.Email,
		sig.UUID,
		tools.CurrentDate(),
	)
	fmt.Println(sql)

	_, Failure = DB.Exec(sql)
	if Failure != nil {
		fmt.Println("Error. Insert record users")
		return Failure
	}
	fmt.Println("Success insert record")
	return nil
}
