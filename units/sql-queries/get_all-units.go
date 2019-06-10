package queries

import (
	"log"

	"github.com/jmoiron/sqlx"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/units/models"
)

const GET_ALL_UNITS = `
SELECT 
	u."UnitId"
	,u."UnitNamePl" 
	,u."UnitNameEn" 
	,u."UnitShortName" 
	,u."QuantityId" 
	,u."Ratio" 
	,q."QuantityNamePl"
FROM 
	units."Units" u
	,units."Quantities" q
WHERE 
	q."QuantityId" IS NOT NULL
	AND u."QuantityId"=q."QuantityId";`

func GetAllUnits() (units []models.GetUnit, err error) {
	db, err := sqlx.Open("postgres", configuration.ConnectionString)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return
	}
	err = db.Select(&units, GET_ALL_UNITS)

	if err != nil {
		log.Fatal(err)
		return
	}

	return
}
