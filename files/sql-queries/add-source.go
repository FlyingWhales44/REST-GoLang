package queries

import (
	"database/sql"
	"log"
	"time"

	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/files/models"
)

const ADD_SOURCE_SQL = `
INSERT INTO files."Sources"
(
	"SourceDate"
	,"DOI"
	,"SourceDescription"
	,"BibText"
	,"FilePathURL"
)
	VALUES ($1, $2, $3, $4, $5);`

func AddSource(source models.SourceForm) (err error) {
	db, err := sql.Open("postgres", configuration.ConnectionString)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer db.Close()

	transaction, err := db.Begin()

	if err != nil {
		log.Fatal(err)
		return
	}

	stmt, err := transaction.Prepare(ADD_SOURCE_SQL)

	if err != nil {
		log.Fatal(err)
		return
	}

	const createdFormat = "2006-01-02"
	updatedAt, _ := time.Parse(createdFormat, source.SourceDate)
	_, err = stmt.Exec(
		updatedAt,
		source.DOI,
		source.SourceDescription,
		source.BibText,
		source.FilePathURL,
	)

	if err != nil {
		log.Fatal(err)
		return
	}

	err = stmt.Close()
	if err != nil {
		log.Fatal(err)
		return
	}

	err = transaction.Commit()
	if err != nil {
		log.Fatal(err)
	}
	return
}
