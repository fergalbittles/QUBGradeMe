package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Load the env variables into a global config map
var config map[string]interface{}

// The struct that will store test results
type TestResult struct {
	Name             string                 `json:"name"`
	Date             time.Time              `json:"date"`
	Result           string                 `json:"result"`
	TimeTaken        string                 `json:"time_taken"`
	ExpectedStatus   string                 `json:"expected_status"`
	ActualStatus     string                 `json:"actual_status"`
	ExpectedResponse string                 `json:"expected_response"`
	ActualResponse   string                 `json:"actual_response"`
	RespBody         map[string]interface{} `json:"response_body"`
}

// The struct that will be returned to the client
type TestResponse struct {
	Date    time.Time     `json:"date"`
	Results []*TestResult `json:"results"`
}

// Database filter payload
type FilterPayload struct {
	ID IDPayload `json:"_id"`
}

// Database ID payload
type IDPayload struct {
	OID primitive.ObjectID `json:"$oid"`
}

// Logs which get inserted into DB
type Logs struct {
	Logs []TestResponse `json:"logs"`
}

// Database update payload
type UpdatePayload struct {
	Set Logs `json:"$set"`
}

// Database update one payload
type UpdateOnePayload struct {
	Collection string        `json:"collection"`
	Database   string        `json:"database"`
	DataSource string        `json:"dataSource"`
	Filter     FilterPayload `json:"filter"`
	Update     UpdatePayload `json:"update"`
}

// Databse find one payload
type FindOnePayload struct {
	Collection string        `json:"collection"`
	Database   string        `json:"database"`
	DataSource string        `json:"dataSource"`
	Filter     FilterPayload `json:"filter"`
}

// Document to store the logs which are fetched from the database
type Document struct {
	ID   IDPayload      `json:"_id"`
	Logs []TestResponse `json:"logs"`
}

// The logs that will be kept in the database
var testLogs []TestResponse

type TestLogResponse struct {
	Logs []TestResponse `json:"logs"`
}

func periodicTesting() {
	// Tests will periodically run every 12 hours
	ticker := time.NewTicker(12 * time.Hour)

	for range ticker.C {
		runTests()
	}
}

func main() {
	loadConfiguration()
	loadTestLogs()

	// Create a goroutine to run tests on an interval
	go periodicTesting()

	e := echo.New()

	e.Use(middleware.CORS())

	e.GET("/", getMetrics)
	e.GET("/logs", getLogs)

	e.Logger.Fatal(e.Start(":1326"))
}

func getMetrics(c echo.Context) error {
	// Run the tests
	startTime := time.Now()
	testResults := runTests()
	executionTime := time.Since(startTime)

	// Log the duration
	log.Println("Duration of tests and storage operation (s): ", executionTime.Seconds())

	// Return the results
	return c.JSON(http.StatusOK, testResults)
}

func getLogs(c echo.Context) error {
	testLogResponse := TestLogResponse{
		Logs: testLogs,
	}

	// Return the logs
	return c.JSON(http.StatusOK, testLogResponse)
}

func runTests() TestResponse {
	log.Println("Running tests...")

	proxyResult := testProxy()
	maxMinResult := testMaxMin()
	sortModulesResult := testSortModules()
	totalMarksResult := testTotalMarks()
	averageMarkResult := testAverageMark()
	classifyResult := testClassify()
	classifyModulesResult := testClassifyModules()
	databaseStoreResult := testDatabaseStore()
	databaseFetchResult := testDatabaseFetch()

	response := []*TestResult{}
	response = append(response, proxyResult)
	response = append(response, maxMinResult)
	response = append(response, sortModulesResult)
	response = append(response, totalMarksResult)
	response = append(response, averageMarkResult)
	response = append(response, classifyResult)
	response = append(response, classifyModulesResult)
	response = append(response, databaseStoreResult)
	response = append(response, databaseFetchResult)

	sendResponse := TestResponse{
		Date:    time.Now(),
		Results: response,
	}

	storeResults(sendResponse)

	return sendResponse
}

func testProxy() *TestResult {
	url := "https://qubgrademe-proxy.up.railway.app/?"

	url += "route=averagemark&"
	url += "mark_1=78&"
	url += "mark_2=65&"
	url += "mark_3=82&"
	url += "mark_4=54&"
	url += "mark_5=68&"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println("Error while initialising http request: ", err.Error())
		return nil
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Access-Control-Request-Headers", "*")

	client := &http.Client{
		Timeout: 20 * time.Second,
	}

	// Send request and time the response
	startTime := time.Now()
	resp, err := client.Do(req)
	executionTime := time.Since(startTime)

	if err != nil {
		log.Println("Error while sending http request: ", err.Error())
		return nil
	}

	// Parse the response body
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Println("Error while reading response body: ", err.Error())
		return nil
	}

	// Parse the body from string to json
	var jsonBody map[string]interface{}
	err = json.Unmarshal([]byte(body), &jsonBody)

	if err != nil {
		log.Println("Error while unmarshalling response body: ", err.Error())
		return nil
	}

	// Compare actual result with expected result
	var testResult string
	resultCheck := false
	statusCheck := false

	actualStatus := resp.Status
	expectedStatus := "200 OK"
	actualResponse := fmt.Sprint(reflect.ValueOf(jsonBody["string"]))
	expectedResponse := "Your average mark is 69.4"

	if actualStatus == expectedStatus {
		statusCheck = true
	}
	if actualResponse == expectedResponse {
		resultCheck = true
	}
	if resultCheck && statusCheck {
		testResult = "PASSED"
	} else {
		testResult = "FAILED"
	}

	testResponse := buildTestResponse("Proxy 1 - Route to Average Mark", testResult, expectedStatus, actualStatus, expectedResponse, actualResponse, executionTime.Seconds(), jsonBody)

	// Send email alert on failure
	if testResponse.Result == "FAILED" {
		sendEmailAlert(testResponse)
	}

	// Return a response
	return &testResponse
}

func testMaxMin() *TestResult {
	url := "https://qubgrademe-maxmin.up.railway.app/?"

	url += "module_1=Object%20Oriented%20Programming&"
	url += "module_2=Cloud%20Computing&"
	url += "module_3=Databases&"
	url += "module_4=Concurrent%20Programming&"
	url += "module_5=Cyber%20Security&"
	url += "mark_1=78&"
	url += "mark_2=65&"
	url += "mark_3=82&"
	url += "mark_4=54&"
	url += "mark_5=68&"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println("Error while initialising http request: ", err.Error())
		return nil
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Access-Control-Request-Headers", "*")

	client := &http.Client{
		Timeout: 20 * time.Second,
	}

	// Send request and time the response
	startTime := time.Now()
	resp, err := client.Do(req)
	executionTime := time.Since(startTime)

	if err != nil {
		log.Println("Error while sending http request: ", err.Error())
		return nil
	}

	// Parse the response body
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Println("Error while reading response body: ", err.Error())
		return nil
	}

	// Parse the body from string to json
	var jsonBody map[string]interface{}
	err = json.Unmarshal([]byte(body), &jsonBody)

	if err != nil {
		log.Println("Error while unmarshalling response body: ", err.Error())
		return nil
	}

	// Compare actual result with expected result
	var testResult string
	resultCheck := false
	statusCheck := false

	extractMaxModule := reflect.ValueOf(jsonBody["max_module"])
	extractMinModule := reflect.ValueOf(jsonBody["min_module"])

	actualStatus := resp.Status
	expectedStatus := "200 OK"
	actualResponse := "Max Module: " + fmt.Sprint(extractMaxModule) + " | Min Module: " + fmt.Sprint(extractMinModule)
	expectedResponse := "Max Module: Databases - 82 | Min Module: Concurrent Programming - 54"

	if actualStatus == expectedStatus {
		statusCheck = true
	}
	if actualResponse == expectedResponse {
		resultCheck = true
	}
	if resultCheck && statusCheck {
		testResult = "PASSED"
	} else {
		testResult = "FAILED"
	}

	testResponse := buildTestResponse("Highest & Lowest Marks", testResult, expectedStatus, actualStatus, expectedResponse, actualResponse, executionTime.Seconds(), jsonBody)

	// Send email alert on failure
	if testResponse.Result == "FAILED" {
		sendEmailAlert(testResponse)
	}

	// Return a response
	return &testResponse
}

func testSortModules() *TestResult {
	url := "https://qubgrademe-sortmodules.up.railway.app/?"

	url += "module_1=Object%20Oriented%20Programming&"
	url += "module_2=Cloud%20Computing&"
	url += "module_3=Databases&"
	url += "module_4=Concurrent%20Programming&"
	url += "module_5=Cyber%20Security&"
	url += "mark_1=78&"
	url += "mark_2=65&"
	url += "mark_3=82&"
	url += "mark_4=54&"
	url += "mark_5=68&"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println("Error while initialising http request: ", err.Error())
		return nil
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Access-Control-Request-Headers", "*")

	client := &http.Client{
		Timeout: 20 * time.Second,
	}

	// Send request and time the response
	startTime := time.Now()
	resp, err := client.Do(req)
	executionTime := time.Since(startTime)

	if err != nil {
		log.Println("Error while sending http request: ", err.Error())
		return nil
	}

	// Parse the response body
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Println("Error while reading response body: ", err.Error())
		return nil
	}

	// Parse the body from string to json
	var jsonBody map[string]interface{}
	err = json.Unmarshal([]byte(body), &jsonBody)

	if err != nil {
		log.Println("Error while unmarshalling response body: ", err.Error())
		return nil
	}

	// Compare actual result with expected result
	var testResult string
	resultCheck := false
	statusCheck := false

	actualStatus := resp.Status
	expectedStatus := "200 OK"
	actualResponse := fmt.Sprint(reflect.ValueOf(jsonBody["sorted_modules"]))
	expectedResponse := "[map[marks:82 module:Databases] map[marks:78 module:Object Oriented Programming] map[marks:68 module:Cyber Security] map[marks:65 module:Cloud Computing] map[marks:54 module:Concurrent Programming]]"

	if actualStatus == expectedStatus {
		statusCheck = true
	}
	if actualResponse == expectedResponse {
		resultCheck = true
	}
	if resultCheck && statusCheck {
		testResult = "PASSED"
	} else {
		testResult = "FAILED"
	}

	testResponse := buildTestResponse("Sort Modules", testResult, expectedStatus, actualStatus, expectedResponse, actualResponse, executionTime.Seconds(), jsonBody)

	// Send email alert on failure
	if testResponse.Result == "FAILED" {
		sendEmailAlert(testResponse)
	}

	// Return a response
	return &testResponse
}

func testTotalMarks() *TestResult {
	url := "https://qubgrademe-totalmarks.up.railway.app/?"

	url += "mark_1=78&"
	url += "mark_2=65&"
	url += "mark_3=82&"
	url += "mark_4=54&"
	url += "mark_5=68&"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println("Error while initialising http request: ", err.Error())
		return nil
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Access-Control-Request-Headers", "*")

	client := &http.Client{
		Timeout: 20 * time.Second,
	}

	// Send request and time the response
	startTime := time.Now()
	resp, err := client.Do(req)
	executionTime := time.Since(startTime)

	if err != nil {
		log.Println("Error while sending http request: ", err.Error())
		return nil
	}

	// Parse the response body
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Println("Error while reading response body: ", err.Error())
		return nil
	}

	// Parse the body from string to json
	var jsonBody map[string]interface{}
	err = json.Unmarshal([]byte(body), &jsonBody)

	if err != nil {
		log.Println("Error while unmarshalling response body: ", err.Error())
		return nil
	}

	// Compare actual result with expected result
	var testResult string
	resultCheck := false
	statusCheck := false

	actualStatus := resp.Status
	expectedStatus := "200 OK"
	actualResponse := fmt.Sprint(reflect.ValueOf(jsonBody["string"]))
	expectedResponse := "Total Marks Acquired = 347"

	if actualStatus == expectedStatus {
		statusCheck = true
	}
	if actualResponse == expectedResponse {
		resultCheck = true
	}
	if resultCheck && statusCheck {
		testResult = "PASSED"
	} else {
		testResult = "FAILED"
	}

	testResponse := buildTestResponse("Total Marks", testResult, expectedStatus, actualStatus, expectedResponse, actualResponse, executionTime.Seconds(), jsonBody)

	// Send email alert on failure
	if testResponse.Result == "FAILED" {
		sendEmailAlert(testResponse)
	}

	// Return a response
	return &testResponse
}

func testAverageMark() *TestResult {
	url := "https://qubgrademe-average.up.railway.app/?"

	url += "mark_1=78&"
	url += "mark_2=65&"
	url += "mark_3=82&"
	url += "mark_4=54&"
	url += "mark_5=68&"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println("Error while initialising http request: ", err.Error())
		return nil
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Access-Control-Request-Headers", "*")

	client := &http.Client{
		Timeout: 20 * time.Second,
	}

	// Send request and time the response
	startTime := time.Now()
	resp, err := client.Do(req)
	executionTime := time.Since(startTime)

	if err != nil {
		log.Println("Error while sending http request: ", err.Error())
		return nil
	}

	// Parse the response body
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Println("Error while reading response body: ", err.Error())
		return nil
	}

	// Parse the body from string to json
	var jsonBody map[string]interface{}
	err = json.Unmarshal([]byte(body), &jsonBody)

	if err != nil {
		log.Println("Error while unmarshalling response body: ", err.Error())
		return nil
	}

	// Compare actual result with expected result
	var testResult string
	resultCheck := false
	statusCheck := false

	actualStatus := resp.Status
	expectedStatus := "200 OK"
	actualResponse := fmt.Sprint(reflect.ValueOf(jsonBody["string"]))
	expectedResponse := "Your average mark is 69.4"

	if actualStatus == expectedStatus {
		statusCheck = true
	}
	if actualResponse == expectedResponse {
		resultCheck = true
	}
	if resultCheck && statusCheck {
		testResult = "PASSED"
	} else {
		testResult = "FAILED"
	}

	testResponse := buildTestResponse("Average Mark", testResult, expectedStatus, actualStatus, expectedResponse, actualResponse, executionTime.Seconds(), jsonBody)

	// Send email alert on failure
	if testResponse.Result == "FAILED" {
		sendEmailAlert(testResponse)
	}

	// Return a response
	return &testResponse
}

func testClassify() *TestResult {
	url := "https://qubgrademe-classify.up.railway.app/?"

	url += "mark_1=78&"
	url += "mark_2=65&"
	url += "mark_3=82&"
	url += "mark_4=54&"
	url += "mark_5=68&"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println("Error while initialising http request: ", err.Error())
		return nil
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Access-Control-Request-Headers", "*")

	client := &http.Client{
		Timeout: 20 * time.Second,
	}

	// Send request and time the response
	startTime := time.Now()
	resp, err := client.Do(req)
	executionTime := time.Since(startTime)

	if err != nil {
		log.Println("Error while sending http request: ", err.Error())
		return nil
	}

	// Parse the response body
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Println("Error while reading response body: ", err.Error())
		return nil
	}

	// Parse the body from string to json
	var jsonBody map[string]interface{}
	err = json.Unmarshal([]byte(body), &jsonBody)

	if err != nil {
		log.Println("Error while unmarshalling response body: ", err.Error())
		return nil
	}

	// Compare actual result with expected result
	var testResult string
	resultCheck := false
	statusCheck := false

	actualStatus := resp.Status
	expectedStatus := "200 OK"
	actualResponse := fmt.Sprint(reflect.ValueOf(jsonBody["string"]))
	expectedResponse := "Your overall classification is Upper Second-Class Honours (2:1)"

	if actualStatus == expectedStatus {
		statusCheck = true
	}
	if actualResponse == expectedResponse {
		resultCheck = true
	}
	if resultCheck && statusCheck {
		testResult = "PASSED"
	} else {
		testResult = "FAILED"
	}

	testResponse := buildTestResponse("Overall Classification", testResult, expectedStatus, actualStatus, expectedResponse, actualResponse, executionTime.Seconds(), jsonBody)

	// Send email alert on failure
	if testResponse.Result == "FAILED" {
		sendEmailAlert(testResponse)
	}

	// Return a response
	return &testResponse
}

func testClassifyModules() *TestResult {
	url := "https://qubgrademe-classifymodules.up.railway.app/?"

	url += "module_1=Object%20Oriented%20Programming&"
	url += "module_2=Cloud%20Computing&"
	url += "module_3=Databases&"
	url += "module_4=Concurrent%20Programming&"
	url += "module_5=Cyber%20Security&"
	url += "mark_1=78&"
	url += "mark_2=65&"
	url += "mark_3=82&"
	url += "mark_4=54&"
	url += "mark_5=68&"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println("Error while initialising http request: ", err.Error())
		return nil
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Access-Control-Request-Headers", "*")

	client := &http.Client{
		Timeout: 20 * time.Second,
	}

	// Send request and time the response
	startTime := time.Now()
	resp, err := client.Do(req)
	executionTime := time.Since(startTime)

	if err != nil {
		log.Println("Error while sending http request: ", err.Error())
		return nil
	}

	// Parse the response body
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Println("Error while reading response body: ", err.Error())
		return nil
	}

	// Parse the body from string to json
	var jsonBody map[string]interface{}
	err = json.Unmarshal([]byte(body), &jsonBody)

	if err != nil {
		log.Println("Error while unmarshalling response body: ", err.Error())
		return nil
	}

	// Compare actual result with expected result
	var testResult string
	resultCheck := false
	statusCheck := false

	actualResponse := ""
	actualResponse += fmt.Sprint(reflect.ValueOf(jsonBody["module_1"])) + " - "
	actualResponse += fmt.Sprint(reflect.ValueOf(jsonBody["mark_1"])) + " | "
	actualResponse += fmt.Sprint(reflect.ValueOf(jsonBody["module_2"])) + " - "
	actualResponse += fmt.Sprint(reflect.ValueOf(jsonBody["mark_2"])) + " | "
	actualResponse += fmt.Sprint(reflect.ValueOf(jsonBody["module_3"])) + " - "
	actualResponse += fmt.Sprint(reflect.ValueOf(jsonBody["mark_3"])) + " | "
	actualResponse += fmt.Sprint(reflect.ValueOf(jsonBody["module_4"])) + " - "
	actualResponse += fmt.Sprint(reflect.ValueOf(jsonBody["mark_4"])) + " | "
	actualResponse += fmt.Sprint(reflect.ValueOf(jsonBody["module_5"])) + " - "
	actualResponse += fmt.Sprint(reflect.ValueOf(jsonBody["mark_5"]))

	actualStatus := resp.Status
	expectedStatus := "200 OK"
	expectedResponse := "Object Oriented Programming - First-Class Honours (1st) | Cloud Computing - Upper Second-Class Honours (2:1) | Databases - First-Class Honours (1st) | Concurrent Programming - Lower Second-Class Honours (2:2) | Cyber Security - Upper Second-Class Honours (2:1)"

	if actualStatus == expectedStatus {
		statusCheck = true
	}
	if actualResponse == expectedResponse {
		resultCheck = true
	}
	if resultCheck && statusCheck {
		testResult = "PASSED"
	} else {
		testResult = "FAILED"
	}

	testResponse := buildTestResponse("Module Classification", testResult, expectedStatus, actualStatus, expectedResponse, actualResponse, executionTime.Seconds(), jsonBody)

	// Send email alert on failure
	if testResponse.Result == "FAILED" {
		sendEmailAlert(testResponse)
	}

	// Return a response
	return &testResponse
}

func testDatabaseStore() *TestResult {
	url := "https://qubgrademe-database.up.railway.app/?"

	url += "id=6628310d466ef72ac24b45c3&"
	url += "module_1=Object%20Oriented%20Programming&"
	url += "module_2=Cloud%20Computing&"
	url += "module_3=Databases&"
	url += "module_4=Concurrent%20Programming&"
	url += "module_5=Cyber%20Security&"
	url += "mark_1=78&"
	url += "mark_2=65&"
	url += "mark_3=82&"
	url += "mark_4=54&"
	url += "mark_5=68&"

	req, err := http.NewRequest("PUT", url, nil)

	if err != nil {
		log.Println("Error while initialising http request: ", err.Error())
		return nil
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Access-Control-Request-Headers", "*")

	client := &http.Client{
		Timeout: 20 * time.Second,
	}

	// Send request and time the response
	startTime := time.Now()
	resp, err := client.Do(req)
	executionTime := time.Since(startTime)

	if err != nil {
		log.Println("Error while sending http request: ", err.Error())
		return nil
	}

	// Parse the response body
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Println("Error while reading response body: ", err.Error())
		return nil
	}

	// Parse the body from string to json
	var jsonBody map[string]interface{}
	err = json.Unmarshal([]byte(body), &jsonBody)

	if err != nil {
		log.Println("Error while unmarshalling response body: ", err.Error())
		return nil
	}

	// Compare actual result with expected result
	var testResult string
	resultCheck := false
	statusCheck := false

	actualStatus := resp.Status
	expectedStatus := "200 OK"
	actualResponse := "Matched Count: " + fmt.Sprint(reflect.ValueOf(jsonBody["matchedCount"]).MapIndex(reflect.ValueOf("$numberInt"))) + " | " + "Modified Count: " + fmt.Sprint(reflect.ValueOf(jsonBody["modifiedCount"]).MapIndex(reflect.ValueOf("$numberInt")))
	expectedResponse := "Matched Count: 1 | Modified Count: 1"

	if actualStatus == expectedStatus {
		statusCheck = true
	}
	if actualResponse == expectedResponse {
		resultCheck = true
	}
	if resultCheck && statusCheck {
		testResult = "PASSED"
	} else {
		testResult = "FAILED"
	}

	testResponse := buildTestResponse("Database Store Operation", testResult, expectedStatus, actualStatus, expectedResponse, actualResponse, executionTime.Seconds(), jsonBody)

	// Send email alert on failure
	if testResponse.Result == "FAILED" {
		sendEmailAlert(testResponse)
	}

	// Return a response
	return &testResponse
}

func testDatabaseFetch() *TestResult {
	url := "https://qubgrademe-database.up.railway.app/?"

	url += "id=6628310d466ef72ac24b45c3"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println("Error while initialising http request: ", err.Error())
		return nil
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Access-Control-Request-Headers", "*")

	client := &http.Client{
		Timeout: 20 * time.Second,
	}

	// Send request and time the response
	startTime := time.Now()
	resp, err := client.Do(req)
	executionTime := time.Since(startTime)

	if err != nil {
		log.Println("Error while sending http request: ", err.Error())
		return nil
	}

	// Parse the response body
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Println("Error while reading response body: ", err.Error())
		return nil
	}

	// Parse the body from string to json
	var jsonBody map[string]interface{}
	err = json.Unmarshal([]byte(body), &jsonBody)

	if err != nil {
		log.Println("Error while unmarshalling response body: ", err.Error())
		return nil
	}

	// Compare actual result with expected result
	var testResult string
	resultCheck := false
	statusCheck := false

	actualStatus := resp.Status
	expectedStatus := "200 OK"
	actualResponse := fmt.Sprint(reflect.ValueOf(jsonBody["document"]))
	actualResponse = strings.ReplaceAll(actualResponse, "map[_id:map[$oid:6628310d466ef72ac24b45c3] ", "")
	expectedResponse := "marks:[78 65 82 54 68] modules:[Object Oriented Programming Cloud Computing Databases Concurrent Programming Cyber Security]]"

	if actualStatus == expectedStatus {
		statusCheck = true
	}
	if actualResponse == expectedResponse {
		resultCheck = true
	}
	if resultCheck && statusCheck {
		testResult = "PASSED"
	} else {
		testResult = "FAILED"
	}

	testResponse := buildTestResponse("Database Fetch Operation", testResult, expectedStatus, actualStatus, expectedResponse, actualResponse, executionTime.Seconds(), jsonBody)

	// Send email alert on failure
	if testResponse.Result == "FAILED" {
		sendEmailAlert(testResponse)
	}

	// Return a response
	return &testResponse
}

func buildTestResponse(name, testResult, expectedStatus, actualStatus, expectedResponse, actualResponse string, executionTime float64, body map[string]interface{}) TestResult {
	testResponse := TestResult{
		Name:             name,
		Date:             time.Now(),
		Result:           testResult,
		TimeTaken:        fmt.Sprint(executionTime),
		ExpectedStatus:   expectedStatus,
		ActualStatus:     actualStatus,
		ExpectedResponse: expectedResponse,
		ActualResponse:   actualResponse,
		RespBody:         body,
	}

	return testResponse
}

func sendEmailAlert(testResult TestResult) {
	// Sender data.
	from := fmt.Sprint(config["senderemail"])
	password := fmt.Sprint(config["password"])

	// Receiver email address.
	to := []string{
		fmt.Sprint(config["receiveremail"]),
	}

	// Build the message
	msg := fmt.Sprintf("From: %s\r\n", from)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(to, ";"))
	msg += "Subject: TEST FAILURE\r\n"
	msg += "\r\nQUB GradeMe App\n"
	msg += fmt.Sprintf("\nThe following test has failed: %s\n", testResult.Name)
	msg += fmt.Sprintf("\nDate of test: %s\n", testResult.Date.Format(time.RFC1123))
	msg += fmt.Sprintf("\nResult: %s", testResult.Result)
	msg += fmt.Sprintf("\nTime taken (s): %s", testResult.TimeTaken)
	msg += fmt.Sprintf("\nExpected status: %s", testResult.ExpectedStatus)
	msg += fmt.Sprintf("\nActual status: %s", testResult.ActualStatus)
	msg += fmt.Sprintf("\nExpected Response: %s", testResult.ExpectedResponse)
	msg += fmt.Sprintf("\nActual Response: %s\n\n", testResult.ActualResponse)
	msg += fmt.Sprintf("Response Body: \n\n%s", testResult.RespBody)

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte(msg)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Println("Error while sending email: ", err.Error())
		return
	}
}

func storeResults(testResults TestResponse) {
	log.Println("Storing test results in database...")

	// Ensure the max number of logs has not been reached
	if len(testLogs) >= 30 {
		testLogs = append(testLogs[1:], testResults)
	} else {
		testLogs = append(testLogs, testResults)
	}

	// Parse the database ID
	id := fmt.Sprint(reflect.ValueOf(config["db_id"]))
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println("Error while reading database ID:", err)
		return
	}

	// Build the payload for the db API call
	logs := Logs{
		Logs: testLogs,
	}
	updatePayload := UpdatePayload{
		Set: logs,
	}
	idPayload := IDPayload{
		OID: objId,
	}
	filterPayload := FilterPayload{
		ID: idPayload,
	}
	jsonPayload := UpdateOnePayload{
		Collection: fmt.Sprint(config["db_collection"]),
		Database:   fmt.Sprint(config["db_database"]),
		DataSource: fmt.Sprint(config["db_datasource"]),
		Filter:     filterPayload,
		Update:     updatePayload,
	}
	marshPayload, err := json.Marshal(jsonPayload)

	if err != nil {
		log.Println("Error while marshalling payload:", err)
		return
	}

	// Send the DB request
	url := fmt.Sprint(config["db_url"]) + "/updateOne"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(marshPayload))

	if err != nil {
		log.Println("Error while building http request:", err)
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Access-Control-Request-Headers", "*")
	req.Header.Add("api-key", fmt.Sprint(config["db_apikey"]))
	req.Header.Add("Accept", "application/ejson")
	client := &http.Client{
		Timeout: 20 * time.Second,
	}
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error while sending http request:", err)
		return
	}

	if resp.StatusCode >= 400 {
		log.Println("Error during database operation:", resp.Status)
		return
	}

	// Parse the response body
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Println("Error while reading database response:", err)
		return
	}

	// Parse the body from string to json
	var jsonBody map[string]interface{}
	err = json.Unmarshal([]byte(body), &jsonBody)

	if err != nil {
		log.Println("Error while unmarshalling database response:", err)
		return
	}

	matched := fmt.Sprint(reflect.ValueOf(jsonBody["matchedCount"]).MapIndex(reflect.ValueOf("$numberInt")))
	modified := fmt.Sprint(reflect.ValueOf(jsonBody["modifiedCount"]).MapIndex(reflect.ValueOf("$numberInt")))

	if matched == "0" || modified == "0" {
		log.Println("Error occurred while trying to find a match for the povided database ID")
		return
	}

	log.Println("Results were successfully stored in the database")
}

func loadConfiguration() {
	log.Println("Loading configuration...")

	config = make(map[string]interface{})
	config["db_apikey"] = os.Getenv("db_apikey")
	config["db_collection"] = os.Getenv("db_collection")
	config["db_database"] = os.Getenv("db_database")
	config["db_datasource"] = os.Getenv("db_datasource")
	config["db_id"] = os.Getenv("db_id")
	config["db_url"] = os.Getenv("db_url")
	config["password"] = os.Getenv("password")
	config["receiveremail"] = os.Getenv("receiveremail")
	config["senderemail"] = os.Getenv("senderemail")

	log.Println("Configuration loaded successfully")
}

func loadTestLogs() {
	log.Println("Fetching test logs from database...")

	// Parse the database ID
	id := fmt.Sprint(reflect.ValueOf(config["db_id"]))
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println("Error while reading database ID:", err)
		return
	}

	// Build the payload for the db API call
	idPayload := IDPayload{
		OID: objId,
	}
	filterPayload := FilterPayload{
		ID: idPayload,
	}
	jsonPayload := FindOnePayload{
		Collection: fmt.Sprint(config["db_collection"]),
		Database:   fmt.Sprint(config["db_database"]),
		DataSource: fmt.Sprint(config["db_datasource"]),
		Filter:     filterPayload,
	}
	marshPayload, err := json.Marshal(jsonPayload)

	if err != nil {
		log.Println("Error while marshalling payload:", err)
		return
	}

	// Send the DB request
	url := fmt.Sprint(config["db_url"]) + "/findOne"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(marshPayload))

	if err != nil {
		log.Println("Error while building http request:", err)
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Access-Control-Request-Headers", "*")
	req.Header.Add("api-key", fmt.Sprint(config["db_apikey"]))
	req.Header.Add("Accept", "application/ejson")
	client := &http.Client{
		Timeout: 20 * time.Second,
	}
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error while sending http request:", err)
		return
	}

	if resp.StatusCode >= 400 {
		log.Println("Error during database operation:", resp.Status)
		return
	}

	// Parse the response body
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Println("Error while reading database response:", err)
		return
	}

	// Parse the body from string to json
	var jsonBody map[string]interface{}
	err = json.Unmarshal([]byte(body), &jsonBody)

	if err != nil {
		log.Println("Error while unmarshalling database response:", err)
		return
	}

	document := fmt.Sprint(reflect.ValueOf(jsonBody["document"]))

	if !strings.Contains(document, "_id:map[$oid:6628319e944e745026e1a31e]") {
		log.Println("Failed to find a record with the provided database ID")
		return
	}

	// Store result in global testLogs variable
	jsonString, err := json.Marshal(jsonBody["document"])

	if err != nil {
		log.Println("Error while storing test logs in local storage:", err)
		return
	}

	documentStruct := Document{}
	json.Unmarshal(jsonString, &documentStruct)

	testLogs = documentStruct.Logs

	log.Println("Test logs were successfully loaded from the database")
}
