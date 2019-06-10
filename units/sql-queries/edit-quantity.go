package queries

import (
	"database/sql"
	"log"

	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/units/models"
)

const EDIT_QUANTITY_SQL = `
	UPDATE units."Quantities"
	SET
		"QuantityNamePl" = $1
		,"QuantityNameEn" = $2
		,"BaseUnitID" = $3
	WHERE 
	"QuantityId" = $4`

func EditQuantity(quantity models.Quantity) (err error) {
	db, err := sql.Open("postgres", configuration.ConnectionString)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer db.Close()

	transaction, err := db.Begin()

	if err != nil {
		log.Fatal(err)
		return
	}

	stmt, err := transaction.Prepare(EDIT_QUANTITY_SQL)

	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = stmt.Exec(
		quantity.QuantityNamePl,
		quantity.QuantityNameEn,
		quantity.BaseUnitID,
		quantity.QuantityId,
	)

	if err != nil {
		log.Fatal(err)
		return
	}

	err = stmt.Close()
	if err != nil {
		log.Fatal(err)
		return
	}

	err = transaction.Commit()
	if err != nil {
		log.Fatal(err)
	}
	return
}
