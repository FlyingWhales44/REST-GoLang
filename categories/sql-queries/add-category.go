package queries

import (
	"database/sql"
	"log"
	"strconv"
	"strings"

	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/categories/models"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
)

const ADD_CATEGORY_SQL = `
	INSERT INTO categories."Categories"
	(
		"CategoryNamePl", 
		"CategoryNameEn",
		"CategoryDescriptionPl", 
		"CategoryDescriptionEn"
	)
		VALUES ($1, $2, $3, $4) RETURNING "CategoryId";`

//t
func AddCategory(category models.Category) (err error) {

	db, err := sql.Open("postgres", configuration.ConnectionString)

	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return
	}

	transaction, err := db.Begin()

	if err != nil {
		log.Fatal(err)
		return
	}

	statement, err := transaction.Prepare(ADD_CATEGORY_SQL)

	if err != nil {
		log.Fatal(err)
		return
	}

	var categoryId int

	err = statement.QueryRow(
		category.CategoryNamePl,
		category.CategoryNameEn,
		category.CategoryDescriptionPl,
		category.CategoryDescriptionEn,
	).Scan(&categoryId)

	if err != nil {
		log.Fatal(err)
		return
	}

	if len(category.CategoryParentsIds) != 0 {

		hierarchyOfCategories_sql := `
			INSERT INTO categories."HierarchyOfCategories"(
				"CategoryId","ParentId")
				VALUES `

		bufI := strconv.Itoa(categoryId)

		for _, v := range category.CategoryParentsIds {
			bufV := strconv.Itoa(v)
			hierarchyOfCategories_sql += `(` + bufI + `,` + bufV + `),`
		}

		newsql := strings.TrimRight(hierarchyOfCategories_sql, ",")

		newsql += `;`

		statement, err = transaction.Prepare(newsql)

		if err != nil {
			log.Fatal(err)
			return
		}

		_, err = statement.Exec()
	}

	err = statement.Close()
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
