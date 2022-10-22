package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type User struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
}

func main() {
	urlExample := fmt.Sprintf("postgres://postgres:%s@localhost:5432/postgres", os.Getenv("password"))
	conn, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	//tmpUser := User{UserName: "Batman331"}
	//InsertUser(&tmpUser, conn)
	GetAllUsers(conn)
}

func InsertUser(u *User, conn *pgx.Conn) {
	if _, err := conn.Exec(context.Background(), "INSERT INTO USERS(USERNAME) VALUES($1)", u.UserName); err != nil {
		fmt.Println("Unable to insert due to: ", err)
		return
	}
	fmt.Println("Insertion Succesfull")
}

func GetAllUsers(conn *pgx.Conn) {
	if rows, err := conn.Query(context.Background(), "SELECT * FROM USERS"); err != nil {
		fmt.Println("Unable to insert due to: ", err)
		return
	} else {
		defer rows.Close()
		var tmp User
		for rows.Next() {
			rows.Scan(&tmp.ID, &tmp.UserName)
			fmt.Printf("%+v\n", tmp)
		}
		if rows.Err() != nil {
			fmt.Println("Error will reading user table: ", err)
			return
		}
	}
}
