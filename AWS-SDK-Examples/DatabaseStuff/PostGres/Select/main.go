package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	host, _ := os.LookupEnv("PGhost")
	port := 5432
	user, _ := os.LookupEnv("PGuser")
	password, _ := os.LookupEnv("PGpassword")
	dbname, _ := os.LookupEnv("PGdbname")

	pgConString := fmt.Sprintf("port=%d host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		port, host, user, password, dbname)

	db, err := sql.Open("postgres", pgConString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := "select * from orders limit 15"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println("Failed to run query", err)
		return
	}

	fmt.Println(rows)
	//
	// cols, err := rows.Columns()
	// if err != nil {
	// 	fmt.Println("Failed to get columns", err)
	// 	return
	// }
	// // Result is your slice string.
	// rawResult := make([][]byte, len(cols))
	// result := make([]string, len(cols))
	//
	// dest := make([]interface{}, len(cols)) // A temporary interface{} slice
	// for i, _ := range rawResult {
	// 	dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
	// }
	//
	// for rows.Next() {
	// 	err = rows.Scan(dest...)
	// 	if err != nil {
	// 		fmt.Println("Failed to scan row", err)
	// 		return
	// 	}
	//
	// 	for i, raw := range rawResult {
	// 		if raw == nil {
	// 			result[i] = "\\N"
	// 		} else {
	// 			result[i] = string(raw)
	// 		}
	// 	}
	//
	// 	fmt.Printf("%#v\n", result[3])
	// }

}
