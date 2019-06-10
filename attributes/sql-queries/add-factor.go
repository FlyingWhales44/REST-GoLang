package queries

import (
	"database/sql"
	"log"

	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/attributes/models"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
)

const ADD_FACTOR_SQL = `
	INSERT INTO attributes."AttributesEnum"
	(
		"AttributeEnumNamePl",
		"AttributeEnumNameEn",
		"AttributeEnumShortName",
		"AttributeEnumDescriptionPl",
		"AttributeEnumDescriptionEn"
	)
		VALUES ($1, $2, $3, $4, $5);`

func AddFactor(factor models.Factor) (err error) {
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

	stmt, err := transaction.Prepare(ADD_FACTOR_SQL)

	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = stmt.Exec(
		factor.AttributeEnumNamePl,
		factor.AttributeEnumNameEn,
		factor.AttributeEnumShortName,
		factor.AttributeEnumDescriptionPl,
		factor.AttributeEnumDescriptionEn,
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
