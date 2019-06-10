package queries

import (
	"log"

	"github.com/jmoiron/sqlx"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/attributes/models"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
)

const GET_RESOURCE_SQL = `
SELECT 
	r."ResourceId"
	,r."ResourceNamePl"
	,r."ResourceNameEn"
	,r."ResourceDescriptionPl"
	,r."ResourceDescriptionEn"
FROM 
	resources."Resources" r;`

const GET_RESOURCE_ATTRIBUTES_SQL = `
SELECT 
	a."AttributeId"
	,re."ResourceId"
	,a."AttributeEnumId"
	,ae."AttributeEnumNamePl"
	,a."SourceId"
	,s."SourceDate"
	,s."SourceDescription"
	,a."Value"
	,a."Uncertainty"
	,a."ValueUnitId1"
	,a."ValueUnitId2"
	,a."ValueUnitId3"
	,u1."UnitShortName" Unit1ShortName
	,u2."UnitShortName" Unit2ShortName
	,u3."UnitShortName" Unit3ShortName
FROM 
	resources."ResourcesAttributes" re 
	JOIN attributes."Attributes" a ON a."AttributeId"= re."AttributeId"
	JOIN attributes."AttributesEnum" ae ON a."AttributeEnumId" = ae."AttributeEnumId"
	JOIN files."Sources" s ON a."SourceId" = s."SourceId"
	LEFT JOIN units."Units" u1 ON a."ValueUnitId1" = u1."UnitId"
	LEFT JOIN units."Units" u2 ON a."ValueUnitId2" = u2."UnitId"
	LEFT JOIN units."Units" u3 ON a."ValueUnitId3" = u3."UnitId";`

func GetResourcesWithAttributes() (resourcesWithAttributes []models.ResourceModel, err error) {
	db, err := sqlx.Open("postgres", configuration.ConnectionString)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return
	}

	var resources []models.ResourceEdit

	err = db.Select(&resources, GET_RESOURCE_SQL)

	if err != nil {
		log.Fatal(err)
		return
	}

	var attributes []models.FactorAttributeModel

	err = db.Select(&attributes, GET_RESOURCE_ATTRIBUTES_SQL)

	if err != nil {
		log.Fatal(err)
		return
	}

	var buff models.ResourceModel

	for i := 0; i < len(resources); i++ {

		buff.ResourceId = resources[i].ResourceId
		buff.ResourceNameEn = resources[i].ResourceDescriptionEn
		buff.ResourceDescriptionPl = resources[i].ResourceDescriptionPl
		buff.ResourceNameEn = resources[i].ResourceNameEn
		buff.ResourceNamePl = resources[i].ResourceNamePl

		for j := 0; j < len(attributes); j++ {
			if attributes[j].ResourceId == resources[i].ResourceId {
				buff.Attributes = append(buff.Attributes, attributes[j])
			}
		}

		resourcesWithAttributes = append(resourcesWithAttributes, buff)
		buff.Attributes = nil
	}

	return
}
