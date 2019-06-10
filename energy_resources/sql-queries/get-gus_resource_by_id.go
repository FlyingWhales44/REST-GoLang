package queries

import (
	"log"

	"github.com/jmoiron/sqlx"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/energy_resources/models"
)

const GET_GUS_RESOURCE_BY_ID_SQL = `
SELECT 
	*
FROM 
	energy_resources."GUSResources" e
WHERE 
	e."GUSResourceId"=$1;`

func GetGUSResourceById(id int) (gus *models.GUSResource, err error) {
	db, err := sqlx.Open("postgres", configuration.ConnectionString)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return
	}
	var _gus []models.GUSResource
	err = db.Select(&_gus, GET_GUS_RESOURCE_BY_ID_SQL, id)

	if err != nil {
		log.Fatal(err)
		return
	}

	if len(_gus) > 0 {
		gus = &_gus[0]
	}

	return
}
