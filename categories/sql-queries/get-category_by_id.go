package queries

import (
	"log"

	"github.com/jmoiron/sqlx"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/categories/models"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
)

const GET_CATEGORY_BY_ID_SQL = `
SELECT 
	q."CategoryId"
	,q."CategoryNamePl" 
	,q."CategoryNameEn"
	,q."CategoryDescriptionPl"
	,q."CategoryDescriptionEn" 
FROM 
	categories."Categories" q
WHERE 
	q."CategoryId"=$1;`

func GetCategoryById(id int) (category *models.EditCategory, err error) {
	db, err := sqlx.Open("postgres", configuration.ConnectionString)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return
	}
	var _category []models.EditCategory
	err = db.Select(&_category, GET_CATEGORY_BY_ID_SQL, id)

	if err != nil {
		log.Fatal(err)
		return
	}

	if len(_category) > 0 {
		category = &_category[0]
	}

	return
}
