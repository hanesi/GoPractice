package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type streamRow struct {
	playedAt         string
	albumType        string
	albumName        string
	releaseDate      string
	albumTotalTracks int
	albumTrackNumber int
	artist           string
	trackName        string
	trackPopularity  int
	trackDuration    int
	trackIsLocal     string
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	host, _ := os.LookupEnv("host")
	port := 3306
	user, _ := os.LookupEnv("user")
	password, _ := os.LookupEnv("password")
	dbname, _ := os.LookupEnv("dbname")

	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		user, password, host, port, dbname)

	db, err := sql.Open("mysql", mysqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := "select * from StreamingLog where `track.duration_ms` is not null limit 15;"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println("Failed to run query", err)
		return
	}

	rowSlice := []streamRow{}
	for rows.Next() {
		var row streamRow

		if err := rows.Scan(&row.playedAt, &row.albumType, &row.albumName, &row.releaseDate, &row.albumTotalTracks, &row.albumTrackNumber, &row.artist, &row.trackName, &row.trackPopularity, &row.trackDuration, &row.trackIsLocal); err != nil {
			log.Fatal(err)
		}
		rowSlice = append(rowSlice, row)
	}

	for _, v := range rowSlice {
		fmt.Println(v)
	}

}
