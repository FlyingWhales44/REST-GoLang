package queries

import (
	"log"

	"github.com/jmoiron/sqlx"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/energy_resources/models"
)

const GET_GUS_ID_SQL = `
SELECT 
	g."GUSResourceId"
	,g."GUSResourceNamePl" 
	,g."GUSId" 
FROM 
	energy_resources."GUSResources" g;`

func GetGUSResourcesId() (gusResources []models.GUSResourceId, err error) {
	db, err := sqlx.Open("postgres", configuration.ConnectionString)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return
	}
	err = db.Select(&gusResources, GET_GUS_ID_SQL)

	if err != nil {
		log.Fatal(err)
		return
	}

	return
}
