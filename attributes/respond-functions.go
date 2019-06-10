package attributes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/attributes/models"
	queries "gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/attributes/sql-queries"
)

func addFactorAttribute(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var body models.FactorAttribute
	err := decoder.Decode(&body)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	err = queries.AddFactorAttribute(body)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	return
}

func addFactor(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var body models.Factor
	err := decoder.Decode(&body)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	err = queries.AddFactor(body)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	return
}

func addAttributeToExistingResource(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var body models.AttributeRes
	err := decoder.Decode(&body)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	err = queries.AddAttributeToExistingResource(body)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	return
}

func addResource(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var body models.Resource
	err := decoder.Decode(&body)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	err = queries.AddResource(body)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	return
}

func getAllFactorAttributes(res http.ResponseWriter, req *http.Request) {

	getFactorAttribute, err := queries.GetAllAttributes()

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	sourceJson, err := json.Marshal(getFactorAttribute)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Write(sourceJson)
	return
}

func getAllFactors(res http.ResponseWriter, req *http.Request) {

	allFactors, err := queries.GetAllFactors()

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	sourceJson, err := json.Marshal(allFactors)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Write(sourceJson)
	return
}

func getFactorsId(res http.ResponseWriter, req *http.Request) {

	getfactorsId, err := queries.GetFactorsId()

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	sourceJson, err := json.Marshal(getfactorsId)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Write(sourceJson)
	return
}

func getResourcesWithAttributes(res http.ResponseWriter, req *http.Request) {

	getResourcesWithAttributes, err := queries.GetResourcesWithAttributes()

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	sourceJson, err := json.Marshal(getResourcesWithAttributes)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Write(sourceJson)
	return
}

func editFactorAttribute(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var body models.FactorAttribute
	err := decoder.Decode(&body)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	err = queries.EditFactorAttribute(body)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	return
}

func editFactor(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var body models.Factor
	err := decoder.Decode(&body)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	err = queries.EditFactor(body)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	return
}

func editResource(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var body models.ResourceEdit
	err := decoder.Decode(&body)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	err = queries.EditResource(body)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	return
}

func deleteFactorResource(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var body models.ResourceFactorId
	err := decoder.Decode(&body)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	err = queries.DeleteFactorResource(body)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	return
}

func getFactorById(res http.ResponseWriter, req *http.Request) {

	id_s := req.URL.Query().Get("id")
	id, _ := strconv.Atoi(id_s)

	factor, err := queries.GetFactorById(id)

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

func getResourceById(res http.ResponseWriter, req *http.Request) {

	id_s := req.URL.Query().Get("id")
	id, _ := strconv.Atoi(id_s)

	resource, err := queries.GetResourceById(id)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	sourceJson, err := json.Marshal(resource)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Write(sourceJson)
	return
}

func getFactorAttributeById(res http.ResponseWriter, req *http.Request) {

	id_s := req.URL.Query().Get("id")
	id, _ := strconv.Atoi(id_s)

	factor, err := queries.GetFactorAttributeById(id)

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
