package queries

import (
	"log"

	"github.com/jmoiron/sqlx"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/units/models"
)

const GET_UNITS_ID_SQL = `
SELECT 
	u."UnitId" 
	,u."UnitShortName" 
FROM 
	units."Units" u;`

func GetUnitsId() (units []models.UnitId, err error) {
	db, err := sqlx.Open("postgres", configuration.ConnectionString)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return
	}
	err = db.Select(&units, GET_UNITS_ID_SQL)

	if err != nil {
		log.Fatal(err)
		return
	}

	return
}
