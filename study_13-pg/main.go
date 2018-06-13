/*CREATE TABLE userinfo
(
uid serial NOT NULL,
username character varying(100) NOT NULL,
departname character varying(500) NOT NULL,
Created date,
CONSTRAINT userinfo_pkey PRIMARY KEY (uid)
)
WITH (OIDS=FALSE);

CREATE TABLE userdetail
(
uid integer,
intro character varying(100),
profile character varying(100)
)
WITH(OIDS=FALSE);

*/


package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://dbUser02:961385@localhost/pgTestDb01?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) RETURNING uid")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)

	//pg不支持这个函数，因为他没有类似MySQL的自增ID
	// id, err := res.LastInsertId()
	// checkErr(err)
	// fmt.Println(id)

	var lastInsertId int
	err = db.QueryRow("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) returning uid;", "astaxie", "研发部门", "2012-12-09").Scan(&lastInsertId)
	checkErr(err)
	fmt.Println("最后插入id =", lastInsertId)


	//更新数据
	stmt, err = db.Prepare("update userinfo set username=$1 where uid=$2")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", 1)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	//查询数据
	fmt.Println("\n\n查询数据：")
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid, username,department,created)
	}


	//删除数据
	fmt.Println("\n\n删除数据：")
	stmt, err = db.Prepare("delete from userinfo where uid=$1")
	checkErr(err)

	for i := 0; i <10; i++  {
		res, err = stmt.Exec(i) //传入要删除的uid
		checkErr(err)

	}


	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}

}