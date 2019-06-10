package queries

import (
	"log"

	"github.com/jmoiron/sqlx"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/units/models"
)

const GET_QUANTITY_ID = `
SELECT 
	q."QuantityId"
	,q."QuantityNamePl"
	,u."UnitShortName" 
FROM 
	units."Quantities" q
	,units."Units" u 
WHERE 
	q."BaseUnitID"=u."UnitId";`

func GetQuantityId() (units []models.QuantityId, err error) {
	db, err := sqlx.Open("postgres", configuration.ConnectionString)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return
	}
	err = db.Select(&units, GET_QUANTITY_ID)

	if err != nil {
		log.Fatal(err)
		return
	}

	return
}
