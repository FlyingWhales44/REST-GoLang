package queries

import (
	"log"

	"github.com/jmoiron/sqlx"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/attributes/models"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
)

const GET_FACTORS_SQL = `
SELECT * FROM attributes."AttributesEnum";`

func GetAllFactors() (factors []models.Factor, err error) {
	db, err := sqlx.Open("postgres", configuration.ConnectionString)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return
	}
	err = db.Select(&factors, GET_FACTORS_SQL)

	if err != nil {
		log.Fatal(err)
		return
	}

	return
}
