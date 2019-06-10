package queries

import (
	"log"

	"github.com/jmoiron/sqlx"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/units/models"
)

const GET_QUANITY_BY_ID_SQL = `
SELECT 
	q."QuantityId" 
	,q."QuantityNamePl"
	,q."QuantityNameEn"
	,q."BaseUnitID" 
FROM 
	units."Quantities" q
WHERE 
	q."QuantityId"=$1;`

func GetQuanityById(id int) (quanity *models.Quantity, err error) {
	db, err := sqlx.Open("postgres", configuration.ConnectionString)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return
	}
	var _qunity []models.Quantity
	err = db.Select(&_qunity, GET_QUANITY_BY_ID_SQL, id)

	if err != nil {
		log.Fatal(err)
		return
	}

	if len(_qunity) > 0 {
		quanity = &_qunity[0]
	}

	return
}
