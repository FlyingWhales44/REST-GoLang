package queries

import (
	"log"

	"github.com/jmoiron/sqlx"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/units/models"
)

const GET_BASIC_UNIT_BY_ID_SQL = `
SELECT 
	u."UnitId" 
	,u."UnitNamePl"
	,u."UnitNameEn"
	,u."UnitShortName" 
	,u."Ratio"
FROM 
	units."Units" u 
WHERE 
	u."QuantityId" IS NULL
	AND u."UnitId"=$1;`

func GetBasicUnitById(id int) (unit *models.BasicUnit, err error) {
	db, err := sqlx.Open("postgres", configuration.ConnectionString)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return
	}
	var units []models.BasicUnit
	err = db.Select(&units, GET_BASIC_UNIT_BY_ID_SQL, id)

	if err != nil {
		log.Fatal(err)
		return
	}

	if len(units) > 0 {
		unit = &units[0]
	}

	return
}
