package queries

import (
	"log"

	"github.com/jmoiron/sqlx"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/files/models"
)

const GET_SOURCE_SQL = `
SELECT 
	f."SourceId" 
	,f."SourceDate"
	,f."DOI"
	,f."SourceDescription" 
FROM 
	files."Sources" f;`

func GetSource() (source []models.SourceInfo, err error) {
	db, err := sqlx.Open("postgres", configuration.ConnectionString)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return
	}
	err = db.Select(&source, GET_SOURCE_SQL)

	if err != nil {
		log.Fatal(err)
		return
	}

	return
}
