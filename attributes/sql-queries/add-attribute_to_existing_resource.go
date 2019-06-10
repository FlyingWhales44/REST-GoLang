package queries

import (
	"database/sql"
	"log"

	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/attributes/models"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
)

const ADD_ATTRIBUTE_TO_EXISTING_RESOURCE_SQL = `
INSERT INTO resources."ResourcesAttributes"
(
	"ResourceId"
	,"AttributeId"
)
	VALUES ($1,$2)
`

func AddAttributeToExistingResource(attribute models.AttributeRes) (err error) {
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

	stmt, err := transaction.Prepare(ADD_ATTRIBUTE_TO_EXISTING_RESOURCE_SQL)

	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = stmt.Exec(
		attribute.ResourceId,
		attribute.AttributeId,
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
