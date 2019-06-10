package queries

import (
	"database/sql"
	"log"

	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/attributes/models"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
)

const EDIT_FACTOR_SQL = `
UPDATE attributes."AttributesEnum"
SET
	"AttributeEnumNamePl" = $1
	,"AttributeEnumNameEn" = $2
	,"AttributeEnumShortName" = $3
	,"AttributeEnumDescriptionPl" = $4
	,"AttributeEnumDescriptionEn" = $5
WHERE 
"AttributeEnumId" = $6;`

func EditFactor(factor models.Factor) (err error) {
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

	stmt, err := transaction.Prepare(EDIT_FACTOR_SQL)

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
		factor.AttributeEnumId,
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
