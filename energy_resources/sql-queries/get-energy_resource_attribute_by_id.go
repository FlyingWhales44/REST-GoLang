package queries

import (
	"log"

	"github.com/jmoiron/sqlx"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/energy_resources/models"
)

const GET_ENERGY_RESOURCE_ATTRIBUTE_BY_ID_SQL = `
SELECT 
	e."EnergyResourceAttributeId"
	,e."EnergyResourceId"
	,er."EnergyResourceName"
	,er."GUSResourceId"
	,g."GUSId"
	,e."SourceId"
	,s."SourceDescription"
	,s."SourceDate"
	,e."CO2Value"
	,e."NCVValue"
	,e."CO2UnitId"
	,u."UnitShortName"
	,u."Ratio"
FROM 
	energy_resources."EnergyResourcesAttributes" e
	,units."Units" u
	,energy_resources."EnergyResources" er
	,energy_resources."GUSResources" g
	,files."Sources" s
WHERE 
	e."EnergyResourceId"=er."EnergyResourceId"
	AND er."GUSResourceId"=g."GUSResourceId"
	AND e."SourceId"=s."SourceId"
	AND e."CO2UnitId"=u."UnitId"
	AND e."EnergyResourceAttributeId"=$1;`

func GetEnergyResourceAttributeById(id int) (attribute *models.EnergyResourceAttributeForm, err error) {
	db, err := sqlx.Open("postgres", configuration.ConnectionString)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return
	}
	var _attribute []models.EnergyResourceAttributeForm
	err = db.Select(&_attribute, GET_ENERGY_RESOURCE_ATTRIBUTE_BY_ID_SQL, id)

	if err != nil {
		log.Fatal(err)
		return
	}

	if len(_attribute) > 0 {
		attribute = &_attribute[0]
	}

	return
}
