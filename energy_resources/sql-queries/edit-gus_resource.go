package queries

import (
	"database/sql"
	"log"

	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/energy_resources/models"
)

const EDIT_GUS_RESOURCE_SQL = `
UPDATE energy_resources."GUSResources"
SET
	"GUSResourceNamePl" = $1
	,"GUSResourceNameEn" = $2
	,"GUSId" = $3
WHERE 
"GUSResourceId" = $4;`

func EditGUSResource(gusResource models.GUSResource) (err error) {
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

	stmt, err := transaction.Prepare(EDIT_GUS_RESOURCE_SQL)

	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = stmt.Exec(
		gusResource.GUSResourceNamePl,
		gusResource.GUSResourceNameEn,
		gusResource.GUSId,
		gusResource.GUSResourceId,
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
