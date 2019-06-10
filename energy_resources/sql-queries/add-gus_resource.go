package queries

import (
	"database/sql"
	"log"

	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/energy_resources/models"
)

const ADD_GUS_RESOURCE_SQL = `
INSERT INTO energy_resources."GUSResources"
(
	"GUSResourceNamePl",
	"GUSResourceNameEn",
	"GUSId"
) 
VALUES ($1, $2, $3);`

func AddGUSResource(gusResource models.GUSResource) (err error) {
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

	stmt, err := transaction.Prepare(ADD_GUS_RESOURCE_SQL)

	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = stmt.Exec(
		gusResource.GUSResourceNamePl,
		gusResource.GUSResourceNameEn,
		gusResource.GUSId,
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
