package queries

import (
	"database/sql"
	"log"

	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/units/models"
)

const EDIT_BASIC_UNIT_SQL = `
	UPDATE units."Units"
	SET
		"UnitNamePl" = $1
		,"UnitNameEn" = $2
		,"UnitShortName" = $3
		,"Ratio" = $4
	WHERE 
	"UnitId" = $5`

func EditBasicUnit(unit models.BasicUnit) (err error) {
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

	stmt, err := transaction.Prepare(EDIT_BASIC_UNIT_SQL)

	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = stmt.Exec(
		unit.UnitNamePl,
		unit.UnitNameEn,
		unit.UnitShortName,
		unit.Ratio,
		unit.UnitId,
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
