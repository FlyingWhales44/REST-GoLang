package queries

import (
	"database/sql"
	"log"

	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/categories/models"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
)

const EDIT_CATEGORY_SQL = `
	UPDATE categories."Categories"
	SET
		"CategoryNamePl" = $1
		,"CategoryNameEn" = $2
		,"CategoryDescriptionPl" = $3
		,"CategoryDescriptionEn" = $4
	WHERE 
	"CategoryId" = $5`

func EditCategory(category models.EditCategory) (err error) {
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

	stmt, err := transaction.Prepare(EDIT_CATEGORY_SQL)

	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = stmt.Exec(
		category.CategoryNamePl,
		category.CategoryNameEn,
		category.CategoryDescriptionPl,
		category.CategoryDescriptionEn,
		category.CategoryId,
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
