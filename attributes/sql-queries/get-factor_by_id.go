package queries

import (
	"log"

	"github.com/jmoiron/sqlx"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/attributes/models"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
)

const GET_FACTOR_BY_ID_SQL = `
SELECT 
	*
FROM 
	attributes."AttributesEnum" a
WHERE 
	a."AttributeEnumId"=$1;`

func GetFactorById(id int) (factor *models.Factor, err error) {
	db, err := sqlx.Open("postgres", configuration.ConnectionString)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return
	}
	var _factor []models.Factor
	err = db.Select(&_factor, GET_FACTOR_BY_ID_SQL, id)

	if err != nil {
		log.Fatal(err)
		return
	}

	if len(_factor) > 0 {
		factor = &_factor[0]
	}

	return
}
