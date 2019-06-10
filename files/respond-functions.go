package files

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/files/models"
	queries "gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/files/sql-queries"
)

func addSource(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var body models.SourceForm
	err := decoder.Decode(&body)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	err = queries.AddSource(body)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	return
}

func getSource(res http.ResponseWriter, req *http.Request) {

	source, err := queries.GetSource()

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	sourceJson, err := json.Marshal(source)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Write(sourceJson)
	return
}
func getAllSource(res http.ResponseWriter, req *http.Request) {

	source, err := queries.GetAllSource()

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	sourceJson, err := json.Marshal(source)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Write(sourceJson)
	return
}

func editSource(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var body models.Source
	err := decoder.Decode(&body)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	err = queries.EditSource(body)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	return
}

func getSourceById(res http.ResponseWriter, req *http.Request) {

	id_s := req.URL.Query().Get("id")
	id, _ := strconv.Atoi(id_s)

	source, err := queries.GetSourceById(id)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	sourceJson, err := json.Marshal(source)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Write(sourceJson)
	return
}
