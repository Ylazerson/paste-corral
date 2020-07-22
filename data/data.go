package data

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

// Db represents the top level DB struct.
var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	return
}

// CreateRawPaste adds new raw-paste info int the DB:
func CreateRawPaste(author, title, content, dt string) (err error) {

	statement := `	
		insert into raw_paste_data 
		       (
               author,
               title,
               content,
               paste_date,
			   created_at
			   )
		values ($1, $2, $3, $4)`

	stmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(author, title, content, dt, time.Now())

	return
}
