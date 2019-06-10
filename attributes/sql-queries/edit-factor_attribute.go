package queries

import (
	"database/sql"
	"log"

	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/attributes/models"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
)

const EDIT_FACTOR_ATTRIBUTE_SQL = `
UPDATE attributes."Attributes"
SET
	"AttributeEnumId" = $1
	,"SourceId" = $2
	,"Value" = $3
	,"Uncertainty" = $4
	,"ValueUnitId1" = $5
	,"ValueUnitId2" = $6
	,"ValueUnitId3" = $7
WHERE 
	"AttributeId" = $8;`

func EditFactorAttribute(factorAttribute models.FactorAttribute) (err error) {
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

	stmt, err := transaction.Prepare(EDIT_FACTOR_ATTRIBUTE_SQL)

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
		factorAttribute.FactorAttributeId,
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
