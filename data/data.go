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

// PastesResp contains multiple pastes.
type PastesResp struct {
	Pastes []Paste `json:"pastes"`
}

// Paste contains core info for a single PasteBin paste.
type Paste struct {
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	PasteDate time.Time `json:"pdate"`
}

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
		values ($1, $2, $3, $4, $5)`

	stmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(author, title, content, dt, time.Now())

	return
}

// Pastes gets all pastes in the database.
func Pastes() (pResp PastesResp, err error) {

	sqlStmt := `
	    select   author, 
                 title,
                 content,
                 created_at
        from     raw_paste_data
        order by created_at desc
	`

	rows, err := Db.Query(sqlStmt)

	if err != nil {
		return
	}

	pResp = PastesResp{}

	for rows.Next() {

		p := Paste{}

		if err = rows.Scan(&p.Author, &p.Title, &p.Content, &p.PasteDate); err != nil {
			return
		}

		pResp.Pastes = append(pResp.Pastes, p)
	}

	rows.Close()

	return pResp, err
}
