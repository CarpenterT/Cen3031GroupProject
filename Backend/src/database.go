package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	//"os"
	"time"

	_ "github.com/godror/godror"
)

/*
 * Set connectString env
 * [bash]$setenv DB_URL "oracle://demo:demo@ip:1521/XEPDB1"
 */

func initOracle() {
	fmt.Print("yes")
	// Get db pool object
	//connectString := os.Getenv("DB_URL")
	db, err := sql.Open("godror", "oracle://lgernier:VeQy3lhlAY6Bn9z85sEJLhZf@oracle.cise.ufl.edu/orcl")
	if err != nil { // nil means no error
		log.Fatal(err)
	}
	defer db.Close()

	// Cleanup
	db.Exec("DROP TABLE test")

	// set deadline
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// create a table test with JSON column type
	_, err = db.ExecContext(ctx,
		"CREATE TABLE test (UserID NUMBER(6), username varchar(26), password varchar(26))", //nolint:gas
	)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Exec("DROP TABLE test")

}
