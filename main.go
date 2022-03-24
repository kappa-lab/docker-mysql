package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbconf := "remote-user:abc@tcp(127.0.0.1:3306)"

	db, err := sql.Open("mysql", dbconf)

	defer db.Close()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(db)

	err = db.Ping()

	if err != nil {
		fmt.Println("Ping 失敗")
		return
	} else {
		fmt.Println("Ping 成功")
	}

	ctx := context.TODO()
	cnn, err := db.Conn(ctx)

	if err != nil {
		fmt.Println("Connect 失敗")
		return
	} else {
		fmt.Println("Connect 成功")
	}

	fmt.Println("cnn", cnn)

	res, _ := db.Query("SHOW DATABASES")
	fmt.Println("----------------databases----------------")
	var database string
	for res.Next() {
		res.Scan(&database)
		fmt.Println(database)
	}
	res.Close()

}
