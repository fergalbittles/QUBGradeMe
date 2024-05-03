package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ModuleMarks struct {
	Modules []string `json:"modules"`
	Marks   []string `json:"marks"`
}

type InsertOnePayload struct {
	Collection string      `json:"collection"`
	Database   string      `json:"database"`
	DataSource string      `json:"dataSource"`
	Document   ModuleMarks `json:"document"`
}

type FilterPayload struct {
	ID IDPayload `json:"_id"`
}

type IDPayload struct {
	OID primitive.ObjectID `json:"$oid"`
}

type UpdatePayload struct {
	Set ModuleMarks `json:"$set"`
}

type UpdateOnePayload struct {
	Collection string        `json:"collection"`
	Database   string        `json:"database"`
	DataSource string        `json:"dataSource"`
	Filter     FilterPayload `json:"filter"`
	Update     UpdatePayload `json:"update"`
}

type FindOnePayload struct {
	Collection string        `json:"collection"`
	Database   string        `json:"database"`
	DataSource string        `json:"dataSource"`
	Filter     FilterPayload `json:"filter"`
}

// Load the env variables into a global config map
var config map[string]interface{}

func main() {
	loadConfiguration()

	e := echo.New()

	e.Use(middleware.CORS())
	e.GET("/", loadData)
	e.PUT("/", saveData)

	e.Logger.Fatal(e.Start(":1325"))
}

func loadData(c echo.Context) error {
	// Handle the user input
	id := strings.TrimSpace(c.QueryParam("id"))

	if id == "" {
		msg := "ID value is missing"
		return errorResponse(c, msg, http.StatusBadRequest)
	}

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		msg := "ID provided is not a valid ID"
		return errorResponse(c, msg, http.StatusBadRequest)
	}

	// Execute the db API request
	jsonBody := executeGetRequest(objId)

	if jsonBody == nil {
		msg := "Error: Something went wrong with the server"
		return errorResponse(c, msg, http.StatusInternalServerError)
	}

	// Ensure that a record was found for the ID provided
	if jsonBody["document"] == nil {
		msg := "No record was found for the ID provided"
		return errorResponse(c, msg, http.StatusBadRequest)
	}

	// Return the data to the client
	return c.JSON(http.StatusOK, jsonBody)
}

func saveData(c echo.Context) error {
	// Handle the user input
	marks := []string{}
	modules := []string{}

	id := strings.TrimSpace(c.QueryParam("id"))
	modules = append(modules, strings.TrimSpace(c.QueryParam("module_1")))
	modules = append(modules, strings.TrimSpace(c.QueryParam("module_2")))
	modules = append(modules, strings.TrimSpace(c.QueryParam("module_3")))
	modules = append(modules, strings.TrimSpace(c.QueryParam("module_4")))
	modules = append(modules, strings.TrimSpace(c.QueryParam("module_5")))
	marks = append(marks, strings.TrimSpace(c.QueryParam("mark_1")))
	marks = append(marks, strings.TrimSpace(c.QueryParam("mark_2")))
	marks = append(marks, strings.TrimSpace(c.QueryParam("mark_3")))
	marks = append(marks, strings.TrimSpace(c.QueryParam("mark_4")))
	marks = append(marks, strings.TrimSpace(c.QueryParam("mark_5")))

	// Handle module input
	for i, module := range modules {
		// Missing parameter error
		if module == "" {
			msg := "Module " + fmt.Sprint(i+1) + " value is missing"
			return errorResponse(c, msg, http.StatusBadRequest)
		}

		// Invalid module value
		_, err := strconv.Atoi(fmt.Sprint(module))
		if err == nil {
			msg := "You must provide a string value for Module " + fmt.Sprint(i+1)
			return errorResponse(c, msg, http.StatusBadRequest)
		}
	}

	// Handle mark input
	for i, mark := range marks {
		// Missing parameter error
		if mark == "" {
			msg := "Mark " + fmt.Sprint(i+1) + " value is missing"
			return errorResponse(c, msg, http.StatusBadRequest)
		}

		// Invalid input error
		mark, err := strconv.Atoi(fmt.Sprint(mark))
		if err != nil {
			msg := "You must provide a valid integer for Mark " + fmt.Sprint(i+1)
			return errorResponse(c, msg, http.StatusBadRequest)
		}
		if mark < 0 {
			msg := "You must provide a non-negative integer for Mark " + fmt.Sprint(i+1)
			return errorResponse(c, msg, http.StatusBadRequest)
		}
		if mark > 100 {
			msg := "You cannot exceed 100 marks for Mark " + fmt.Sprint(i+1)
			return errorResponse(c, msg, http.StatusBadRequest)
		}
	}

	// Check if client wants to overwrite an existing record
	if id != "" {
		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			msg := "ID provided is not a valid ID"
			return errorResponse(c, msg, http.StatusBadRequest)
		}

		// Execute the db API request
		jsonBody := executeUpdateRequest(modules, marks, objId)

		if jsonBody == nil {
			msg := "Error: Something went wrong with the server"
			return errorResponse(c, msg, http.StatusInternalServerError)
		}

		matchedCount, modifiedCount := extractMatchedCount(jsonBody)

		if matchedCount == 0 {
			msg := "No record was found for the ID provided"
			return errorResponse(c, msg, http.StatusBadRequest)
		}

		if modifiedCount == 0 {
			msg := "Error: Failed to update record for the ID provided"
			return errorResponse(c, msg, http.StatusInternalServerError)
		}

		// Return the unique identifier to the client
		return c.JSON(http.StatusOK, jsonBody)
	}

	// Execute the db API request
	jsonBody := executeInsertRequest(modules, marks)

	if jsonBody == nil {
		msg := "Error: Something went wrong with the server"
		return errorResponse(c, msg, http.StatusInternalServerError)
	}

	// Return the unique identifier to the client
	return c.JSON(http.StatusOK, jsonBody)
}

func executeGetRequest(id primitive.ObjectID) map[string]interface{} {
	// Build the payload for the db API call
	idPayload := IDPayload{
		OID: id,
	}
	filterPayload := FilterPayload{
		ID: idPayload,
	}
	jsonPayload := FindOnePayload{
		Collection: fmt.Sprint(config["collection"]),
		Database:   fmt.Sprint(config["database"]),
		DataSource: fmt.Sprint(config["datasource"]),
		Filter:     filterPayload,
	}
	marshPayload, err := json.Marshal(jsonPayload)

	if err != nil {
		return nil
	}

	url := fmt.Sprint(config["url"]) + "/findOne"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(marshPayload))

	if err != nil {
		return nil
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Access-Control-Request-Headers", "*")
	req.Header.Add("api-key", fmt.Sprint(config["apikey"]))
	req.Header.Add("Accept", "application/ejson")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil
	}

	if resp.StatusCode >= 400 {
		return nil
	}

	// Parse the response body
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		return nil
	}

	// Parse the body from string to json
	var jsonBody map[string]interface{}
	err = json.Unmarshal([]byte(body), &jsonBody)

	if err != nil {
		return nil
	}

	return jsonBody
}

func executeInsertRequest(modules []string, marks []string) map[string]interface{} {
	// Build the payload for the db API call
	modulemarks := ModuleMarks{
		Modules: modules,
		Marks:   marks,
	}
	jsonPayload := InsertOnePayload{
		Collection: fmt.Sprint(config["collection"]),
		Database:   fmt.Sprint(config["database"]),
		DataSource: fmt.Sprint(config["datasource"]),
		Document:   modulemarks,
	}
	marshPayload, err := json.Marshal(jsonPayload)

	if err != nil {
		return nil
	}

	url := fmt.Sprint(config["url"]) + "/insertOne"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(marshPayload))

	if err != nil {
		return nil
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Access-Control-Request-Headers", "*")
	req.Header.Add("api-key", fmt.Sprint(config["apikey"]))
	req.Header.Add("Accept", "application/ejson")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil
	}

	if resp.StatusCode >= 400 {
		return nil
	}

	// Parse the response body
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		return nil
	}

	// Parse the body from string to json
	var jsonBody map[string]interface{}
	err = json.Unmarshal([]byte(body), &jsonBody)

	if err != nil {
		return nil
	}

	return jsonBody
}

func executeUpdateRequest(modules []string, marks []string, id primitive.ObjectID) map[string]interface{} {
	// Build the payload for the db API call
	modulemarks := ModuleMarks{
		Modules: modules,
		Marks:   marks,
	}
	updatePayload := UpdatePayload{
		Set: modulemarks,
	}
	idPayload := IDPayload{
		OID: id,
	}
	filterPayload := FilterPayload{
		ID: idPayload,
	}
	jsonPayload := UpdateOnePayload{
		Collection: fmt.Sprint(config["collection"]),
		Database:   fmt.Sprint(config["database"]),
		DataSource: fmt.Sprint(config["datasource"]),
		Filter:     filterPayload,
		Update:     updatePayload,
	}
	marshPayload, err := json.Marshal(jsonPayload)

	if err != nil {
		return nil
	}

	url := fmt.Sprint(config["url"]) + "/updateOne"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(marshPayload))

	if err != nil {
		return nil
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Access-Control-Request-Headers", "*")
	req.Header.Add("api-key", fmt.Sprint(config["apikey"]))
	req.Header.Add("Accept", "application/ejson")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil
	}

	if resp.StatusCode >= 400 {
		return nil
	}

	// Parse the response body
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		return nil
	}

	// Parse the body from string to json
	var jsonBody map[string]interface{}
	err = json.Unmarshal([]byte(body), &jsonBody)

	if err != nil {
		return nil
	}

	return jsonBody
}

func errorResponse(c echo.Context, msg string, statusCode int) error {
	err := struct {
		Error  bool   `json:"error"`
		String string `json:"string"`
	}{
		Error:  true,
		String: msg,
	}
	return c.JSON(statusCode, err)
}

func loadConfiguration() {
	log.Println("Loading configuration...")

	config = make(map[string]interface{})
	config["url"] = os.Getenv("url")
	config["apikey"] = os.Getenv("apikey")
	config["database"] = os.Getenv("database")
	config["datasource"] = os.Getenv("datasource")
	config["collection"] = os.Getenv("collection")

	log.Println("Configuration loaded successfully")
}

func extractMatchedCount(body map[string]interface{}) (int, int) {
	extractMatched := reflect.ValueOf(body["matchedCount"])
	extractModified := reflect.ValueOf(body["modifiedCount"])

	matched, err := strconv.Atoi(fmt.Sprint(extractMatched.MapIndex(reflect.ValueOf("$numberInt"))))

	if err != nil {
		return 0, 0
	}

	modified, err := strconv.Atoi(fmt.Sprint(extractModified.MapIndex(reflect.ValueOf("$numberInt"))))

	if err != nil {
		return matched, 0
	}

	return matched, modified
}
