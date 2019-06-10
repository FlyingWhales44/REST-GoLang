package queries

import (
	"database/sql"
	"log"

	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/energy_resources/models"
)

const ADD_ENERGY_RESOURCE_SQL = `
INSERT INTO energy_resources."EnergyResources"
(
	"EnergyResourceName"
	,"GUSResourceId"
) 
VALUES ($1, $2) RETURNING "EnergyResourceId";`

const ADD_ENERGY_RESOURCE_ATTRIBUTE_SQL = `
INSERT INTO energy_resources."EnergyResourcesAttributes"
(
	"EnergyResourceId"
	,"SourceId"
	,"CO2Value"
	,"NCVValue"
	,"CO2UnitId"
) 
VALUES ($1, $2, $3, $4, $5);`

const GET_ENERGY_RESOURCE_ID_SQL = `
SELECT 
	e."EnergyResourceId"
FROM
	energy_resources."EnergyResources" e
WHERE
	e."EnergyResourceName" = $1;
`

func AddEnergyResource(energyResource models.EnergyResource) (err error) {
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

	stmt, err := transaction.Prepare(GET_ENERGY_RESOURCE_ID_SQL)

	if err != nil {
		log.Fatal(err)
		return
	}

	var energyResourceId int

	row := stmt.QueryRow(
		energyResource.EnergyResourceAttributeName,
	)

	err = row.Scan(&energyResourceId)

	if energyResourceId == 0 {
		stmt, err = transaction.Prepare(ADD_ENERGY_RESOURCE_SQL)

		if err != nil {
			log.Fatal(err)
			return
		}

		err = stmt.QueryRow(
			energyResource.EnergyResourceAttributeName,
			energyResource.GUSResourceId,
		).Scan(&energyResourceId)

		if err != nil {
			log.Fatal(err)
			return
		}
	}

	stmt, err = transaction.Prepare(ADD_ENERGY_RESOURCE_ATTRIBUTE_SQL)

	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = stmt.Exec(
		energyResourceId,
		energyResource.SourceId,
		energyResource.CO2Value,
		energyResource.NCVValue,
		energyResource.CO2UnitId,
	)

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
