package queries

import (
	"database/sql"
	"log"

	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/files/models"
)

const EDIT_SOURCE_SQL = `
UPDATE files."Sources"
	SET
		"SourceDate" = $1
		,"DOI" = $2
		,"SourceDescription" = $3
		,"BibText" = $4
		,"FilePathURL" = $5
	WHERE 
	"SourceId" = $6`

func EditSource(source models.Source) (err error) {
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

	stmt, err := transaction.Prepare(EDIT_SOURCE_SQL)

	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = stmt.Exec(
		source.SourceDate,
		source.DOI,
		source.SourceDescription,
		source.BibText,
		source.FilePathURL,
		source.SourceId,
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
