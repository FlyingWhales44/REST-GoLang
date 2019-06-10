package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	httplib "gitlab.com/semestr-6/projekt-grupowy/backend/go-libs/http-lib"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
)

const (
	insert_endpoints = `
	INSERT INTO identity."Endpoints" 
    ("Url") 
	VALUES %s ON CONFLICT ON CONSTRAINT "Endpoints_Url_key" DO NOTHING RETURNING "Id";`

	insert_permission = `
	INSERT INTO identity."Permissions" 
	("Name", "Url") 
	VALUES ('%s', '%s') RETURNING "Id";`

	insert_permission_endpoints = `
	INSERT INTO identity."PermissionEndpoints" 
	("PermissionId", "EndpointId") 
	VALUES %s;`
)

func addEndpoints(routes httplib.Routes) {
	routesmap := routesToString(routes)
	db, err := sqlx.Open("postgres", configuration.ConnectionString)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
		return
	}

	ids := make(map[int][]int, 0)

	for k := range routesmap {
		if k != "OPTIONS" {
			temp := []int{}
			id := 0
			rows, _ := db.Query(fmt.Sprintf(insert_endpoints, routesmap[k]))
			for rows.Next() {
				idt := 0
				rows.Scan(&idt)
				temp = append(temp, idt)
			}
			err = db.QueryRow(fmt.Sprintf(insert_permission, strings.ToUpper(k), strings.ToLower(k))).Scan(&id)
			ids[id] = temp
		}
	}

	permissionendpoints := mapToString(ids)

	db.Query(fmt.Sprintf(insert_permission_endpoints, permissionendpoints))
}

func routesToString(routes httplib.Routes) map[string]string {
	result := make(map[string]string, 0)
	for i := range routes {
		result[routes[i].HttpMethod] += "('" + routes[i].Route + "'), "
	}

	for k := range result {
		result[k] = strings.TrimSuffix(result[k], ", ")
	}

	return result
}

func mapToString(mapa map[int][]int) string {
	result := ""
	for k := range mapa {
		for i := range mapa[k] {
			result += "(" + strconv.Itoa(k) + ", " + strconv.Itoa(mapa[k][i]) + "), "
		}
	}

	result = strings.TrimSuffix(result, ", ")

	return result
}
