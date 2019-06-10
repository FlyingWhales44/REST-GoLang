package queries

import (
	"database/sql"
	"log"

	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/units/models"
)

const ADD_BASIC_UNIT_SQL = `
	INSERT INTO units."Units"
	(
		"UnitNamePl"
		,"UnitNameEn"
		,"UnitShortName"
		,"Ratio"
	)
		VALUES ($1, $2, $3, $4);`

func AddBasicUnit(unit models.BasicUnit) (err error) {
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

	stmt, err := transaction.Prepare(ADD_BASIC_UNIT_SQL)

	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = stmt.Exec(
		unit.UnitNamePl,
		unit.UnitNameEn,
		unit.UnitShortName,
		unit.Ratio,
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
