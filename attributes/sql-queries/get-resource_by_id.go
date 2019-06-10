package queries

import (
	"log"

	"github.com/jmoiron/sqlx"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/attributes/models"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
)

const GET_RESOURCE_BY_ID_SQL = `
SELECT 
	*
FROM 
	resources."Resources" r
WHERE 
	r."ResourceId"=$1;`

func GetResourceById(id int) (resource *models.ResourceEdit, err error) {
	db, err := sqlx.Open("postgres", configuration.ConnectionString)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return
	}
	var _resource []models.ResourceEdit
	err = db.Select(&_resource, GET_RESOURCE_BY_ID_SQL, id)

	if err != nil {
		log.Fatal(err)
		return
	}

	if len(_resource) > 0 {
		resource = &_resource[0]
	}

	return
}
