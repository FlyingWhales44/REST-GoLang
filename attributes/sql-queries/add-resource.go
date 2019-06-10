package queries

import (
	"database/sql"
	"log"
	"strconv"
	"strings"

	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/attributes/models"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
)

const ADD_RESOURCE_SQL = `
	INSERT INTO resources."Resources"
	(
		"ResourceNamePl",
		"ResourceNameEn",
		"ResourceDescriptionPl",
		"ResourceDescriptionEn"
	)
		VALUES ($1, $2, $3, $4) RETURNING "ResourceId";`

func AddResource(resource models.Resource) (err error) {
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

	stmt, err := transaction.Prepare(ADD_RESOURCE_SQL)

	if err != nil {
		log.Fatal(err)
		return
	}

	var resourceId int

	err = stmt.QueryRow(
		resource.ResourceNamePl,
		resource.ResourceNameEn,
		resource.ResourceDescriptionPl,
		resource.ResourceDescriptionEn,
	).Scan(&resourceId)

	if err != nil {
		log.Fatal(err)
		return
	}

	if len(resource.AttributeIds) != 0 {

		addresourcesattribute_sql := `
		INSERT INTO resources."ResourcesAttributes"(
			"ResourceId","AttributeId")
			VALUES `

		bufI := strconv.Itoa(resourceId)

		for _, v := range resource.AttributeIds {
			bufV := strconv.Itoa(v)
			addresourcesattribute_sql += `(` + bufI + `,` + bufV + `),`
		}

		newsql := strings.TrimRight(addresourcesattribute_sql, ",")

		newsql += `;`
		stmt, err = transaction.Prepare(newsql)

		if err != nil {
			log.Fatal(err)
			return
		}

		_, err = stmt.Exec()

		if err != nil {
			log.Fatal(err)
			return
		}

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
