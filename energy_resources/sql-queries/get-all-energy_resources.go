package queries

import (
	"log"

	"github.com/jmoiron/sqlx"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/energy_resources/models"
)

const GET_ENERGY_RESOURCES_SQL = `
SELECT 
	* 
FROM 
	energy_resources."EnergyResources";`

func GetAllEnergyResources() (energyRes []models.EnergyResourceEdit, err error) {
	db, err := sqlx.Open("postgres", configuration.ConnectionString)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return
	}
	err = db.Select(&energyRes, GET_ENERGY_RESOURCES_SQL)

	if err != nil {
		log.Fatal(err)
		return
	}

	return
}
