package energy_resources

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/energy_resources/models"
	queries "gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/energy_resources/sql-queries"
)

func addEnergyResourceAttribute(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var body models.EnergyResource
	err := decoder.Decode(&body)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	err = queries.AddEnergyResource(body)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	return
}

func addGUSResourceAttribute(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var body models.GUSResource
	err := decoder.Decode(&body)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	err = queries.AddGUSResource(body)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	return
}

func getAllEnergyResourceAttributes(res http.ResponseWriter, req *http.Request) {

	energyResourceAttributes, err := queries.GetAllEnergyResourceAttributes()

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	sourceJson, err := json.Marshal(energyResourceAttributes)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Write(sourceJson)
	return
}

func getAllEnergyResources(res http.ResponseWriter, req *http.Request) {

	energyResources, err := queries.GetAllEnergyResources()

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	sourceJson, err := json.Marshal(energyResources)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Write(sourceJson)
	return
}

func getAllGUSResources(res http.ResponseWriter, req *http.Request) {

	GUSResources, err := queries.GetAllGUSResource()

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	sourceJson, err := json.Marshal(GUSResources)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Write(sourceJson)
	return
}
func getGUSResourcesID(res http.ResponseWriter, req *http.Request) {

	GUSResourcesID, err := queries.GetGUSResourcesId()

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	sourceJson, err := json.Marshal(GUSResourcesID)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Write(sourceJson)
	return
}

func editEnergyResourceAttribute(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var body models.EnergyResourceAttributeEdit
	err := decoder.Decode(&body)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	err = queries.EditEnergyResourceAttribute(body)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	return
}

func editGUSResource(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var body models.GUSResource
	err := decoder.Decode(&body)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	err = queries.EditGUSResource(body)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	return
}

func getGUSResourceById(res http.ResponseWriter, req *http.Request) {

	id_s := req.URL.Query().Get("id")
	id, _ := strconv.Atoi(id_s)

	gus, err := queries.GetGUSResourceById(id)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	sourceJson, err := json.Marshal(gus)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Write(sourceJson)
	return
}

func getEnergyResourceAttributeById(res http.ResponseWriter, req *http.Request) {

	id_s := req.URL.Query().Get("id")
	id, _ := strconv.Atoi(id_s)

	attribute, err := queries.GetEnergyResourceAttributeById(id)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	sourceJson, err := json.Marshal(attribute)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Write(sourceJson)
	return
}
