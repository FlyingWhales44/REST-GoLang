package queries

import (
	"database/sql"
	"log"

	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/attributes/models"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
)

const EDIT_RESOURCE_SQL = `
UPDATE resources."Resources"
SET
	"ResourceNamePl" = $1
	,"ResourceNameEn" = $2
	,"ResourceDescriptionPl" = $3
	,"ResourceDescriptionEn" = $4
WHERE 
"ResourceId" = $5;`

func EditResource(resource models.ResourceEdit) (err error) {
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

	stmt, err := transaction.Prepare(EDIT_RESOURCE_SQL)

	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = stmt.Exec(
		resource.ResourceNamePl,
		resource.ResourceNameEn,
		resource.ResourceDescriptionPl,
		resource.ResourceDescriptionEn,
		resource.ResourceId,
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
