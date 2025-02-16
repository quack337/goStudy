package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	mydb "github.com/quack337/goLib/db"
	mysqlite "github.com/quack337/goLib/db/sqlite"
)

func main() {
	db, err := sql.Open("sqlite3", "file:/D/DB/student2.sqlite?_foreign_keys=on")
	if err != nil {
		panic(err)
	}
	fieldInfos, err := mysqlite.GetFieldInfos(db, "testJson")
	if err != nil {
		fmt.Println(err)
	} 
	fmt.Println(fieldInfos)
	fmt.Println("----------------")

	rows, err := db.Query("SELECT * FROM testJson" )
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	fmt.Println(mydb.Jsonify(rows))
	fmt.Println("----------------")

	foreignKeyInfos, err := mysqlite.GetForeignKeyInfos(db, "student")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(foreignKeyInfos)
	fmt.Println("----------------")

	sql, err := mysqlite.MakeSelectSQLWithFkAssociation(db, "student")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sql)
	fmt.Println("----------------")

	rows, err = db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	fmt.Println(mydb.Jsonify(rows))
	fmt.Println("----------------")
}
