package categories

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/categories/models"
	queries "gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/categories/sql-queries"
)

func addCategory(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var body models.Category
	err := decoder.Decode(&body)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	err = queries.AddCategory(body)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	return
}

func deleteCategory(res http.ResponseWriter, req *http.Request) {

	id_s := req.URL.Query().Get("id")
	id, _ := strconv.Atoi(id_s)

	err := queries.DeleteCategory(id)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	return
}

func getCategoryById(res http.ResponseWriter, req *http.Request) {

	id_s := req.URL.Query().Get("id")
	id, _ := strconv.Atoi(id_s)

	factor, err := queries.GetCategoryById(id)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	sourceJson, err := json.Marshal(factor)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Write(sourceJson)
	return
}

func editCategory(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var body models.EditCategory
	err := decoder.Decode(&body)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	err = queries.EditCategory(body)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	return
}

func getCategories(res http.ResponseWriter, req *http.Request) {

	getCategories, err := queries.GetCategories()

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	sourceJson, err := json.Marshal(getCategories)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Write(sourceJson)
	return
}

func getCategoryId(res http.ResponseWriter, req *http.Request) {

	getCategories, err := queries.GetCategoryId()

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	sourceJson, err := json.Marshal(getCategories)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Write(sourceJson)
	return
}
