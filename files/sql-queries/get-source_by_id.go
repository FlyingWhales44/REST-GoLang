package queries

import (
	"log"

	"github.com/jmoiron/sqlx"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/files/models"
)

const GET_SOURCE_BY_ID_SQL = `
SELECT 
	f."SourceId" 
	,f."SourceDate"
	,f."DOI"
	,f."SourceDescription" 
	,f."BibText"
	,f."FilePathURL"
FROM 
	files."Sources" f 
WHERE 
	f."SourceId"=$1;`

func GetSourceById(id int) (source *models.Source, err error) {
	db, err := sqlx.Open("postgres", configuration.ConnectionString)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return
	}
	var _source []models.Source
	err = db.Select(&_source, GET_SOURCE_BY_ID_SQL, id)

	if err != nil {
		log.Fatal(err)
		return
	}

	if len(_source) > 0 {
		source = &_source[0]
	}

	return
}
