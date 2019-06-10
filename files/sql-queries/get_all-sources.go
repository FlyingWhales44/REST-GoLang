package queries

import (
	"log"

	"github.com/jmoiron/sqlx"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/files/models"
)

const GET_ALL_SOURCES_SQL = `
SELECT 
	* 
FROM 
	files."Sources";`

func GetAllSource() (source []models.Source, err error) {
	db, err := sqlx.Open("postgres", configuration.ConnectionString)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return
	}
	err = db.Select(&source, GET_ALL_SOURCES_SQL)

	if err != nil {
		log.Fatal(err)
		return
	}

	return
}
