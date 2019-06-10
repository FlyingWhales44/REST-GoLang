package units

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/units/models"
	queries "gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/units/sql-queries"
)

func addUnit(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var body models.Unit
	err := decoder.Decode(&body)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	err = queries.AddUnit(body)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	return
}

func addQuantity(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var body models.Quantity
	err := decoder.Decode(&body)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	err = queries.AddQuantity(body)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	return
}

func addBasicUnit(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var body models.BasicUnit
	err := decoder.Decode(&body)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	body.Ratio = 1.0
	err = queries.AddBasicUnit(body)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	return
}

func getAllBasicUnits(res http.ResponseWriter, req *http.Request) {

	basicUnit, err := queries.GetAllBasicUnits()

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	sourceJson, err := json.Marshal(basicUnit)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Write(sourceJson)
	return
}

func getAllQuantities(res http.ResponseWriter, req *http.Request) {

	quantities, err := queries.GetAllQuantities()

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	sourceJson, err := json.Marshal(quantities)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Write(sourceJson)
	return
}

func getAllUnits(res http.ResponseWriter, req *http.Request) {

	unit, err := queries.GetAllUnits()

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	sourceJson, err := json.Marshal(unit)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Write(sourceJson)
	return
}

func getBasicUnitsId(res http.ResponseWriter, req *http.Request) {

	basicUnits, err := queries.GetBasicUnitsId()

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	sourceJson, err := json.Marshal(basicUnits)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Write(sourceJson)
	return
}

func getQuantityId(res http.ResponseWriter, req *http.Request) {

	allUnits, err := queries.GetQuantityId()

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	sourceJson, err := json.Marshal(allUnits)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Write(sourceJson)
	return
}

func getUnitsId(res http.ResponseWriter, req *http.Request) {

	units, err := queries.GetUnitsId()

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	sourceJson, err := json.Marshal(units)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Write(sourceJson)
	return
}

func editBasicUnit(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var body models.BasicUnit
	err := decoder.Decode(&body)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	err = queries.EditBasicUnit(body)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	return
}

func editQuantity(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var body models.Quantity
	err := decoder.Decode(&body)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	err = queries.EditQuantity(body)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	return
}

func editUnit(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var body models.Unit
	err := decoder.Decode(&body)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	err = queries.EditUnit(body)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	return
}

func getBasicUnitById(res http.ResponseWriter, req *http.Request) {

	id_s := req.URL.Query().Get("id")
	id, _ := strconv.Atoi(id_s)

	basicUnit, err := queries.GetBasicUnitById(id)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	sourceJson, err := json.Marshal(basicUnit)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Write(sourceJson)
	return
}

func getUnitById(res http.ResponseWriter, req *http.Request) {

	id_s := req.URL.Query().Get("id")
	id, _ := strconv.Atoi(id_s)

	unit, err := queries.GetUnitById(id)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	sourceJson, err := json.Marshal(unit)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Write(sourceJson)
	return
}

func getQuanityById(res http.ResponseWriter, req *http.Request) {

	id_s := req.URL.Query().Get("id")
	id, _ := strconv.Atoi(id_s)

	quanity, err := queries.GetQuanityById(id)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	sourceJson, err := json.Marshal(quanity)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Write(sourceJson)
	return
}
