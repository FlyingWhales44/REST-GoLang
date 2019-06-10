package queries

import (
	"log"

	"github.com/jmoiron/sqlx"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/attributes/models"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
)

const GET_ALL_FACTORS_ATTRIBUTES_SQL = `
SELECT
	a."AttributeId"
	,a."AttributeEnumId"
	,a."SourceId"
	,s."SourceDate"
	,s."SourceDescription"
	,ea."AttributeEnumNamePl"
	,a."Value"
	,a."ValueUnitId1"
	,a."ValueUnitId2"
	,a."ValueUnitId3"
	,u1."UnitShortName" Unit1ShortName
	,u2."UnitShortName" Unit2ShortName
	,u3."UnitShortName" Unit3ShortName
	,a."Uncertainty"
FROM
	attributes."Attributes" a
	JOIN attributes."AttributesEnum" ea ON a."AttributeEnumId" = ea."AttributeEnumId"
	JOIN files."Sources" s ON a."SourceId" = s."SourceId"
	JOIN units."Units" u1 ON a."ValueUnitId1" = u1."UnitId"
	LEFT JOIN units."Units" u2 ON a."ValueUnitId2" = u2."UnitId"
	LEFT JOIN units."Units" u3 ON a."ValueUnitId3" = u3."UnitId";
`

func GetAllAttributes() (attribute []models.FactorAttributeModel, err error) {
	db, err := sqlx.Open("postgres", configuration.ConnectionString)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return
	}
	err = db.Select(&attribute, GET_ALL_FACTORS_ATTRIBUTES_SQL)

	if err != nil {
		log.Fatal(err)
		return
	}

	return
}
