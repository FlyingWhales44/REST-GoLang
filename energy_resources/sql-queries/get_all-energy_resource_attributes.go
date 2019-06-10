package queries

import (
	"log"

	"github.com/jmoiron/sqlx"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/energy_resources/models"
)

const getAllEnergyResourceAttributesSQL = `
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
	AND e."CO2UnitId"=u."UnitId";`

func GetAllEnergyResourceAttributes() (energyResources []models.EnergyResourceAttributeForm, err error) {
	db, err := sqlx.Open("postgres", configuration.ConnectionString)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return
	}
	err = db.Select(&energyResources, getAllEnergyResourceAttributesSQL)

	if err != nil {
		log.Fatal(err)
		return
	}

	return
}
