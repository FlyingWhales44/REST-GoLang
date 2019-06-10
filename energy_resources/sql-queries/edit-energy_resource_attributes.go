package queries

import (
	"database/sql"
	"log"

	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/energy_resources/models"
)

const EDIT_ENERGY_RESOURCE_ATTRIBUTE_SQL = `
UPDATE energy_resources."EnergyResourcesAttributes"
SET
	"EnergyResourceId" = $1
	,"SourceId" = $2
	,"CO2Value" = $3
	,"NCVValue" = $4
	,"CO2UnitId" = $5
WHERE 
"EnergyResourceAttributeId" = $6`

const EDIT_ENERGY_RESOURCE_SQL = `
UPDATE energy_resources."EnergyResources"
SET
	"EnergyResourceName" = $1
	,"GUSResourceId" = $2
WHERE 
"EnergyResourceId" = $3;`

func EditEnergyResourceAttribute(energyResourceEdit models.EnergyResourceAttributeEdit) (err error) {
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

	stmt, err := transaction.Prepare(EDIT_ENERGY_RESOURCE_ATTRIBUTE_SQL)

	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = stmt.Exec(
		energyResourceEdit.EnergyResourceId,
		energyResourceEdit.SourceId,
		energyResourceEdit.CO2Value,
		energyResourceEdit.NCVValue,
		energyResourceEdit.CO2UnitId,
		energyResourceEdit.EnergyResourceAttributeId,
	)

	if err != nil {
		log.Fatal(err)
		return
	}

	stmt, err = transaction.Prepare(EDIT_ENERGY_RESOURCE_SQL)

	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = stmt.Exec(
		energyResourceEdit.EnergyResourceName,
		energyResourceEdit.GUSResourceId,
		energyResourceEdit.EnergyResourceId,
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
