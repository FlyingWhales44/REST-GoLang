package queries

import (
	"log"

	"github.com/jmoiron/sqlx"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/categories/models"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
)

const GET_CATEGORIES_SQL = `
SELECT 
	c."CategoryId" 
	, c."CategoryNamePl" 
	, c."CategoryNameEn" 
	, c."CategoryDescriptionPl"
	, c."CategoryDescriptionEn"
FROM 
	categories."Categories" c;`

const GET_HIERARCHY_SQL = `
SELECT 
	c."CategoryId" 
	, c."ParentId" 
FROM 
	categories."HierarchyOfCategories" c;`

const GET_CATEGORY_BY_ID2_SQL = `
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

func GetCategories() (categories []models.CategoryForm, err error) {
	db, err := sqlx.Open("postgres", configuration.ConnectionString)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return
	}
	var _categories []models.EditCategory

	err = db.Select(&_categories, GET_CATEGORIES_SQL)

	if err != nil {
		log.Fatal(err)
		return
	}

	var hierarchy []models.Hierarchy

	err = db.Select(&hierarchy, GET_HIERARCHY_SQL)

	if err != nil {
		log.Fatal(err)
		return
	}

	var buf models.CategoryForm

	for i := 0; i < len(_categories); i++ {
		buf.CategoryId = _categories[i].CategoryId
		buf.CategoryNameEn = _categories[i].CategoryNameEn
		buf.CategoryNamePl = _categories[i].CategoryNamePl
		buf.CategoryDescriptionPl = _categories[i].CategoryDescriptionPl
		buf.CategoryDescriptionEn = _categories[i].CategoryDescriptionEn

		for j := 0; j < len(hierarchy); j++ {

			if _categories[i].CategoryId == hierarchy[j].CategoryId {
				err = db.Select(&buf.CategoryParents, GET_CATEGORY_BY_ID2_SQL, hierarchy[j].ParentId)

				if err != nil {
					log.Fatal(err)
					return
				}
			}
		}

		categories = append(categories, buf)
		buf.CategoryParents = nil
	}

	return
}
