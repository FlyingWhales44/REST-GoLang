package queries

import (
	"log"

	"github.com/jmoiron/sqlx"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/categories/models"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
)

const GET_CATEGORY_ID_SQL = `
SELECT 
	q."CategoryId";
	,q."CategoryNamePl" 
FROM 
	categories."Categories" q;
`

func GetCategoryId() (category []models.EditCategory, err error) {
	db, err := sqlx.Open("postgres", configuration.ConnectionString)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return
	}
	err = db.Select(&category, GET_CATEGORY_ID_SQL)

	if err != nil {
		log.Fatal(err)
		return
	}

	return
}
