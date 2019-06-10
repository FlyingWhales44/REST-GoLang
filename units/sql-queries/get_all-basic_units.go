package queries

import (
	"log"

	"github.com/jmoiron/sqlx"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/units/models"
)

const GET_ALL_BASIC_UNITS_SQL = `
SELECT 
	u."UnitId"
	,u."UnitNamePl"
	,u."UnitNameEn"
	,u."UnitShortName"
	,u."Ratio"
FROM 
	units."Units" u
WHERE 
	"QuantityId" IS NULL;`

func GetAllBasicUnits() (units []models.BasicUnit, err error) {
	db, err := sqlx.Open("postgres", configuration.ConnectionString)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return
	}
	err = db.Select(&units, GET_ALL_BASIC_UNITS_SQL)

	if err != nil {
		log.Fatal(err)
		return
	}

	return
}
