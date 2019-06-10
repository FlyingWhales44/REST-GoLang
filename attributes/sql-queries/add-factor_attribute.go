package queries

import (
	"database/sql"
	"log"

	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/attributes/models"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
)

const ADD_FACTOR_ATTRIBUTE_SQL = `
	INSERT INTO attributes."Attributes"
	(
		"AttributeEnumId",
		"SourceId",
		"Value",
		"Uncertainty",
		"ValueUnitId1",
		"ValueUnitId2",
		"ValueUnitId3"
	)
		VALUES ($1, $2, $3, $4, $5, $6, $7);`

func AddFactorAttribute(factorAttribute models.FactorAttribute) (err error) {
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

	stmt, err := transaction.Prepare(ADD_FACTOR_ATTRIBUTE_SQL)

	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = stmt.Exec(
		factorAttribute.FactorId,
		factorAttribute.SourceId,
		factorAttribute.Value,
		factorAttribute.Uncertainty,
		factorAttribute.ValueUnitId1,
		factorAttribute.ValueUnitId2,
		factorAttribute.ValueUnitId3,
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
