package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId   int    `db:"user_id"`
	UserName string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var DB *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go-test")
	if err != nil {
		fmt.Println("open mysql failed", err)
		return
	}
	DB = database
}

func main() {
	//insert()
	//selects()
	//update()
	//delete()
	translation()
}

func insert() {
	result, err := DB.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed", err)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("exec failed", err)
		return
	}
	fmt.Println("insert success:", id)
}

func selects() {
	var person []Person
	//必须传入切片
	err := DB.Select(&person, "select user_id, username, sex, email from person where user_id=?", 2)
	if err != nil {
		fmt.Println("exec failed", err)
		return
	}
	fmt.Println("select success:", person)
}

func update() {
	res, err := DB.Exec("update person set username=? where user_id=?", "stu0003", 1)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
	}
	fmt.Println("update success:", row)
}

func delete() {
	res, err := DB.Exec("delete from person where user_id=?", 2)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
	}

	fmt.Println("delete success: ", row)
}

func translation() {
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("begin failed :", err)
		return
	}
	r, err := tx.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		tx.Rollback()
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		tx.Rollback()
		return
	}
	fmt.Println("insert success:", id)

	r, err = tx.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		tx.Rollback()
		return
	}
	id, err = r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		tx.Rollback()
		return
	}
	fmt.Println("insert success:", id)
	tx.Commit()
}
