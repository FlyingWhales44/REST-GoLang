package queries

import (
	"log"

	"github.com/jmoiron/sqlx"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/units/models"
)

const GET_ALL_QUANTITIES_SQL = `
SELECT 
	q."QuantityId" 
	, q."QuantityNamePl" 
	, q."QuantityNameEn" 
	, q."BaseUnitID"
	, u."UnitShortName"
FROM 
	units."Quantities" q
	,units."Units" u
WHERE
	q."BaseUnitID"=u."UnitId";`

func GetAllQuantities() (quantities []models.GetQuantity, err error) {
	db, err := sqlx.Open("postgres", configuration.ConnectionString)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return
	}
	err = db.Select(&quantities, GET_ALL_QUANTITIES_SQL)

	if err != nil {
		log.Fatal(err)
		return
	}

	return
}
