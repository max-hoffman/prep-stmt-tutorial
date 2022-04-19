package main

import (
	"testing"
)

var Result interface{}

func BenchmarkPrepStmt(b *testing.B) {
	db := newDb()
	stmt, err := db.Preparex(`INSERT INTO objects (label, bbox) VALUES (?, ?)`)
	if err != nil {
		b.Fatal(err)
	}
	var res interface{}
	for n := 0; n < b.N; n++ {
		res = stmt.MustExec("label", "box")
	}
	Result = res
}

func BenchmarkExec(b *testing.B) {
	db := newDb()
	var res interface{}
	for n := 0; n < b.N; n++ {
		res = db.MustExec(`INSERT INTO objects (label, bbox) VALUES ('label', 'bbox')`)
	}
	Result = res
}
