package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/funayman/manticore"
)

func main() {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}
	manti := manticore.Wrap(db)

	for _, mdb := range []manticore.DB{db, manti} {
		fmt.Printf("%#v\n", mdb.Driver())
		if err := ping(mdb); err != nil {
			fmt.Printf("ping err: %v\n", err)
		}
	}
}

func ping(db manticore.DB) error {
	return db.Ping()
}
