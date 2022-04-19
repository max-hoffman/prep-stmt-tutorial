package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var db *sqlx.DB

func newDb() *sqlx.DB {
	db, err := sqlx.Connect("mysql", "user:password@(localhost:3308)/tmp3")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	db.MustExec("drop table if exists objects")
	db.MustExec(`CREATE TABLE objects (
	    id int primary key auto_increment,
		label varchar(40),
	    bbox varchar (40)
	)`)
	return db
}

func init() {
	db = newDb()
}

type Object struct {
	Id    int
	Label string
	Bbox  string
}

func main() {
	rows := []struct {
		label string
		bbox  string
	}{
		{
			label: "cat",
			bbox:  "[1,2,3,4]",
		},
		{
			label: "rabbit",
			bbox:  "[9,10,11,12]",
		},
		{
			label: "dog",
			bbox:  "[5,6,7,8]",
		},
	}

	// PREPARE for inserting rows
	stmt, err := db.Preparex(`INSERT INTO objects (label, bbox) VALUES (?, ?)`)
	if err != nil {
		log.Fatal(err)
	}

	// parameterize prepared INSERT query with row values
	for _, r := range rows {
		_ = stmt.MustExec(r.label, r.bbox)
	}

	// SELECT list of rows
	var ids []int
	err = db.Select(&ids, "select id from objects")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("inserted ids: ", ids)

	// PREPARE and SELECT rows by id
	stmt, err = db.Preparex(`SELECT * FROM objects WHERE id=?`)
	if err != nil {
		log.Fatal(err)
	}
	var o Object
	for _, id := range ids {
		err := stmt.Get(&o, id)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("found row: ", o)
	}
}
