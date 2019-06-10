package queries

import (
	"log"

	"github.com/jmoiron/sqlx"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/attributes/models"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
)

const GET_ALL_FACTORS_SHORT_NAME_SQL = `
SELECT 
	a."AttributeEnumId" ,
	a."AttributeEnumShortName" 
FROM 
	attributes."AttributesEnum" a;`

func GetFactorsId() (factors []models.FactorId, err error) {
	db, err := sqlx.Open("postgres", configuration.ConnectionString)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return
	}
	err = db.Select(&factors, GET_ALL_FACTORS_SHORT_NAME_SQL)

	if err != nil {
		log.Fatal(err)
		return
	}

	return
}
