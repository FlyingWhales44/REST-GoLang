package queries

import (
	"log"

	"github.com/jmoiron/sqlx"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/attributes/models"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
)

const GET_FACTOR_ATTRIBUTE_BY_ID_SQL = `
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
	attributes."AttributesEnum" ea
	, attributes."Attributes" a
	, files."Sources" s 
	, units."Units" u1
	, units."Units" u2
	, units."Units" u3
WHERE
	a."SourceId" = s."SourceId"
	AND a."AttributeEnumId" = ea."AttributeEnumId"
	AND a."ValueUnitId1" = u1."UnitId"
	AND a."ValueUnitId2" = u2."UnitId"
	AND a."ValueUnitId3" = u3."UnitId"
	AND a."AttributeId"=$1;`

func GetFactorAttributeById(id int) (attribute *models.FactorAttributeModel, err error) {
	db, err := sqlx.Open("postgres", configuration.ConnectionString)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return
	}
	var _attribute []models.FactorAttributeModel
	err = db.Select(&_attribute, GET_FACTOR_ATTRIBUTE_BY_ID_SQL, id)

	if err != nil {
		log.Fatal(err)
		return
	}

	if len(_attribute) > 0 {
		attribute = &_attribute[0]
	}

	return
}
