package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Load the env variables into a global config map
var config map[string]interface{}

func main() {
	loadConfiguration()

	// Attempt to retrieve an updated config from a live proxy
	loadFromLiveProxy()

	// Start server
	e := echo.New()

	e.Use(middleware.CORS())
	e.GET("/", routeRequest)
	e.PUT("/", storeData)

	// Admin endpoints
	admin := e.Group("/admin")
	admin.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == config["username"] && password == config["password"] {
			return true, nil
		}
		return false, nil
	}))
	admin.GET("/routes", getRoutes)
	admin.POST("/routes", postRoute)
	admin.DELETE("/routes", deleteRoute)
	admin.GET("/login", login)

	e.Logger.Fatal(e.Start(":1324"))
}

func routeRequest(c echo.Context) error {
	// Get route query param
	route := c.QueryParam("route")
	route = strings.TrimSpace(route)

	// Ensure route param is not empty
	if route == "" {
		msg := "Route parameter is missing"
		return errorResponse(c, msg, http.StatusBadRequest)
	}

	// Ensure route param exists in the available routes
	if _, ok := config[route]; !ok || route == "password" || route == "username" {
		msg := "The specified route does not exist"
		return errorResponse(c, msg, http.StatusBadRequest)
	}

	// Iterate through the available instances of the route until a request is successfully sent
	var err error
	var resp *http.Response
	routeInstances := reflect.ValueOf(config[route])
	for i := 0; i < routeInstances.Len(); i++ {
		req := fmt.Sprint(routeInstances.Index(i)) + fmt.Sprint(c.Request().URL)

		if route == "monitorlogs" {
			req = fmt.Sprint(routeInstances.Index(i))
		}

		resp, err = http.Get(req)

		if err != nil && i == routeInstances.Len()-1 {
			msg := "Failed to forward the request to the appropriate route"
			return errorResponse(c, msg, http.StatusMisdirectedRequest)
		}

		if err != nil && i < routeInstances.Len()-1 {
			continue
		}

		if resp.StatusCode >= 500 && i == routeInstances.Len()-1 {
			msg := "Failed to forward the request to the appropriate route"
			return errorResponse(c, msg, http.StatusMisdirectedRequest)
		}

		if resp.StatusCode >= 500 && i < routeInstances.Len()-1 {
			continue
		}

		break
	}

	// Parse the response body and status code
	status := resp.StatusCode
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		msg := "Failed to read the response from the specified route"
		return errorResponse(c, msg, http.StatusBadRequest)
	}

	// Parse the body from string to json
	var jsonBody map[string]interface{}
	err = json.Unmarshal([]byte(body), &jsonBody)
	if err != nil {
		msg := "Failed to parse the response body"
		return errorResponse(c, msg, http.StatusBadRequest)
	}

	// Return the response to the client
	return c.JSON(status, jsonBody)
}

func storeData(c echo.Context) error {
	// Get route query param
	route := c.QueryParam("route")
	route = strings.TrimSpace(route)

	// Ensure route param is not empty
	if route == "" {
		msg := "Route parameter is missing"
		return errorResponse(c, msg, http.StatusBadRequest)
	}

	// Ensure route param exists in the available routes
	if _, ok := config[route]; !ok || route == "password" || route == "username" {
		msg := "The specified route does not exist"
		return errorResponse(c, msg, http.StatusBadRequest)
	}

	// The route must be to the database service
	if route != "database" {
		msg := "The specified route is invalid"
		return errorResponse(c, msg, http.StatusBadRequest)
	}

	// Iterate through the available instances of the route until a request is successfully sent
	var err error
	var resp *http.Response
	routeInstances := reflect.ValueOf(config[route])
	for i := 0; i < routeInstances.Len(); i++ {
		url := fmt.Sprint(routeInstances.Index(i)) + fmt.Sprint(c.Request().URL)
		req, err := http.NewRequest("PUT", url, nil)

		if err != nil && i == routeInstances.Len()-1 {
			msg := "Something went wrong with the server"
			return errorResponse(c, msg, http.StatusInternalServerError)
		}

		if err != nil && i < routeInstances.Len()-1 {
			continue
		}

		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Access-Control-Request-Headers", "*")
		req.Header.Add("api-key", fmt.Sprint(config["apikey"]))
		req.Header.Add("Accept", "application/ejson")
		client := &http.Client{}
		resp, err = client.Do(req)

		if err != nil && i == routeInstances.Len()-1 {
			msg := "Something went wrong with the server"
			return errorResponse(c, msg, http.StatusInternalServerError)
		}

		if err != nil && i < routeInstances.Len()-1 {
			continue
		}

		if resp.StatusCode >= 500 && i == routeInstances.Len()-1 {
			msg := "Something went wrong with the server"
			return errorResponse(c, msg, http.StatusInternalServerError)
		}

		if resp.StatusCode >= 500 && i < routeInstances.Len()-1 {
			continue
		}

		break
	}

	// Parse the response body and status code
	status := resp.StatusCode
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		msg := "Failed to read the response from the specified route"
		return errorResponse(c, msg, http.StatusBadRequest)
	}

	// Parse the body from string to json
	var jsonBody map[string]interface{}
	err = json.Unmarshal([]byte(body), &jsonBody)
	if err != nil {
		msg := "Failed to parse the response body"
		return errorResponse(c, msg, http.StatusBadRequest)
	}

	// Return the response to the client
	return c.JSON(status, jsonBody)
}

func getRoutes(c echo.Context) error {
	response := map[string]interface{}{}

	for key, value := range config {
		if key == "username" || key == "password" {
			continue
		}

		response[key] = value
	}

	return c.JSON(http.StatusOK, response)
}

func postRoute(c echo.Context) error {
	// Check if this proxy has already completed the request
	host := "http://" + c.Request().Host
	proxiesVisited := c.QueryParam("visited")
	if proxiesVisited != "" {
		proxyList := strings.Split(proxiesVisited, ",")

		if contains(proxyList, host) {
			return c.NoContent(http.StatusNoContent)
		}
	}

	// Continue to add new route
	service := c.QueryParam("service")
	route := c.QueryParam("route")

	service = strings.TrimSpace(service)
	route = strings.TrimSpace(route)

	// Ensure service param is not empty
	if service == "" {
		msg := "Service parameter is missing"
		return errorResponse(c, msg, http.StatusBadRequest)
	}

	// Ensure route param is not empty
	if route == "" {
		msg := "Route parameter is missing"
		return errorResponse(c, msg, http.StatusBadRequest)
	}

	_, exists := config[service]

	// Add route to existing service
	if exists {
		tempServiceRoute := []string{}
		routes := reflect.ValueOf(config[service])

		for i := 0; i < routes.Len(); i++ {
			// Ensure the route does not already exist
			if route == fmt.Sprint(routes.Index(i)) {
				msg := "The specified route already exists"
				return errorResponse(c, msg, http.StatusBadRequest)
			}

			tempServiceRoute = append(tempServiceRoute, fmt.Sprint(routes.Index(i)))
		}

		tempServiceRoute = append(tempServiceRoute, route)
		config[service] = tempServiceRoute
	}

	// Create a new service and add the route
	if !exists {
		newServiceRoute := []string{route}
		config[service] = newServiceRoute
	}

	// Forward this request to a proxy which hasn't completed it yet
	knownProxies := []string{}
	routes := reflect.ValueOf(config["proxy"])
	for i := 0; i < routes.Len(); i++ {
		knownProxies = append(knownProxies, fmt.Sprint(routes.Index(i)))
	}

	// No other proxies have been visited
	if proxiesVisited == "" {
		for _, proxy := range knownProxies {
			if proxy != host {
				url := proxy + "/admin/routes?service=" + service + "&route=" + route + "&visited=" + host
				req, err := http.NewRequest("POST", url, nil)

				if err != nil {
					log.Println("Error:", err.Error())
					continue
				}

				req.Header.Add("Authorization", "Basic "+basicAuth(fmt.Sprint(config["username"]), fmt.Sprint(config["password"])))
				client := &http.Client{}
				resp, err := client.Do(req)

				if err != nil {
					log.Println("Error:", err.Error())
					continue
				}

				if resp.StatusCode >= 400 {
					continue
				}

				break
			}
		}
	}

	// Some proxies have already been visited
	if proxiesVisited != "" {
		proxyList := strings.Split(proxiesVisited, ",")

		for _, proxy := range knownProxies {
			if proxy != host && !contains(proxyList, proxy) {
				proxyList = append(proxyList, host)
				visited := strings.Join(proxyList, ",")

				url := proxy + "/admin/routes?service=" + service + "&route=" + route + "&visited=" + visited
				req, err := http.NewRequest("POST", url, nil)

				if err != nil {
					log.Println("Error:", err.Error())
					continue
				}

				req.Header.Add("Authorization", "Basic "+basicAuth(fmt.Sprint(config["username"]), fmt.Sprint(config["password"])))
				client := &http.Client{}
				resp, err := client.Do(req)

				if err != nil {
					log.Println("Error:", err.Error())
					continue
				}

				if resp.StatusCode >= 400 {
					continue
				}

				break
			}
		}
	}

	response := struct {
		Message string `json:"message"`
	}{
		Message: "Endpoint successfully added",
	}

	return c.JSON(http.StatusOK, response)
}

func deleteRoute(c echo.Context) error {
	// Check if this proxy has already completed the request
	host := "http://" + c.Request().Host
	proxiesVisited := c.QueryParam("visited")
	if proxiesVisited != "" {
		proxyList := strings.Split(proxiesVisited, ",")

		if contains(proxyList, host) {
			return c.NoContent(http.StatusNoContent)
		}
	}

	service := c.QueryParam("service")
	route := c.QueryParam("route")

	service = strings.TrimSpace(service)
	route = strings.TrimSpace(route)

	// Ensure service param is not empty
	if service == "" {
		msg := "Service parameter is missing"
		return errorResponse(c, msg, http.StatusBadRequest)
	}

	// Ensure route param is not empty
	if route == "" {
		msg := "Route parameter is missing"
		return errorResponse(c, msg, http.StatusBadRequest)
	}

	// Ensure the service exists
	if _, serviceExists := config[service]; !serviceExists {
		msg := "The specified service does not exist"
		return errorResponse(c, msg, http.StatusBadRequest)
	}

	// Check if the route exists and remove it if so
	tempServiceRoute := []string{}
	routes := reflect.ValueOf(config[service])
	routeExists := false

	for i := 0; i < routes.Len(); i++ {
		if fmt.Sprint(routes.Index(i)) == route {
			routeExists = true
			continue
		}

		tempServiceRoute = append(tempServiceRoute, fmt.Sprint(routes.Index(i)))
	}

	if !routeExists {
		msg := "The specified route does not exist"
		return errorResponse(c, msg, http.StatusBadRequest)
	}

	if len(tempServiceRoute) == 0 {
		delete(config, service)
	} else {
		config[service] = tempServiceRoute
	}

	// Forward this request to a proxy which hasn't completed it yet
	knownProxies := []string{}
	proxyRoutes := reflect.ValueOf(config["proxy"])
	for i := 0; i < proxyRoutes.Len(); i++ {
		knownProxies = append(knownProxies, fmt.Sprint(proxyRoutes.Index(i)))
	}

	// No other proxies have been visited
	if proxiesVisited == "" {
		for _, proxy := range knownProxies {
			if proxy != host {
				url := proxy + "/admin/routes?service=" + service + "&route=" + route + "&visited=" + host
				req, err := http.NewRequest("DELETE", url, nil)

				if err != nil {
					log.Println("Error:", err.Error())
					continue
				}

				req.Header.Add("Authorization", "Basic "+basicAuth(fmt.Sprint(config["username"]), fmt.Sprint(config["password"])))
				client := &http.Client{}
				resp, err := client.Do(req)

				if err != nil {
					log.Println("Error:", err.Error())
					continue
				}

				if resp.StatusCode >= 400 {
					continue
				}

				break
			}
		}
	}

	// Some proxies have already been visited
	if proxiesVisited != "" {
		proxyList := strings.Split(proxiesVisited, ",")

		for _, proxy := range knownProxies {
			if proxy != host && !contains(proxyList, proxy) {
				proxyList = append(proxyList, host)
				visited := strings.Join(proxyList, ",")

				url := proxy + "/admin/routes?service=" + service + "&route=" + route + "&visited=" + visited
				req, err := http.NewRequest("DELETE", url, nil)

				if err != nil {
					log.Println("Error:", err.Error())
					continue
				}

				req.Header.Add("Authorization", "Basic "+basicAuth(fmt.Sprint(config["username"]), fmt.Sprint(config["password"])))
				client := &http.Client{}
				resp, err := client.Do(req)

				if err != nil {
					log.Println("Error:", err.Error())
					continue
				}

				if resp.StatusCode >= 400 {
					continue
				}

				break
			}
		}
	}

	response := struct {
		Message string `json:"message"`
	}{
		Message: "Endpoint successfully deleted",
	}

	return c.JSON(http.StatusOK, response)
}

func login(c echo.Context) error {
	username := strings.TrimSpace(c.QueryParam("username"))
	password := strings.TrimSpace(c.QueryParam("password"))

	if username == "" || password == "" {
		msg := "A username and password must be provided"
		return errorResponse(c, msg, http.StatusUnauthorized)
	}

	actualUsername := fmt.Sprint(reflect.ValueOf(config["username"]))
	actualPassword := fmt.Sprint(reflect.ValueOf(config["password"]))

	if username == actualUsername && password == actualPassword {
		// Authenticate the client
		jsonBody := struct {
			Auth bool `json:"auth"`
		}{
			Auth: true,
		}

		return c.JSON(http.StatusOK, jsonBody)
	}

	msg := "An invalid username or password was provided"
	return errorResponse(c, msg, http.StatusUnauthorized)
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
	config["maxmin"] = []string{os.Getenv("max_min_url")}
	config["sortmodules"] = []string{os.Getenv("sort_modules_url")}
	config["totalmarks"] = []string{os.Getenv("total_marks_url")}
	config["averagemark"] = []string{os.Getenv("average_mark_url")}
	config["classify"] = []string{os.Getenv("classify_url")}
	config["classifymodules"] = []string{os.Getenv("classify_modules_url")}
	config["database"] = []string{os.Getenv("database_url")}
	config["monitor"] = []string{os.Getenv("monitor_url")}
	config["monitorlogs"] = []string{os.Getenv("monitor_logs_url")}
	config["proxy"] = []string{os.Getenv("proxy_url")}
	config["username"] = os.Getenv("proxy_username")
	config["password"] = os.Getenv("proxy_password")

	log.Println("Configuration loaded successfully")
}

func loadFromLiveProxy() {
	proxyRoutes := reflect.ValueOf(config["proxy"])
	username := config["username"]
	password := config["password"]

	for i := 0; i < proxyRoutes.Len(); i++ {
		route := fmt.Sprint(proxyRoutes.Index(i))

		url := route + "/admin/routes"
		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			log.Println("Error:", err.Error())
			continue
		}

		// Send an admin GET request to a proxy
		req.Header.Add("Authorization", "Basic "+basicAuth(fmt.Sprint(config["username"]), fmt.Sprint(config["password"])))
		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			log.Println("Error:", err.Error())
			continue
		}

		if resp.StatusCode >= 400 {
			continue
		}

		// Parse the response body
		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			log.Println("Error:", err.Error())
			continue
		}

		// Parse the body from string to json
		var jsonBody map[string]interface{}
		err = json.Unmarshal([]byte(body), &jsonBody)

		if err != nil {
			log.Println("Error:", err.Error())
			continue
		}

		// Map the response to config variable
		config = jsonBody
		config["username"] = username
		config["password"] = password

		break
	}
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
