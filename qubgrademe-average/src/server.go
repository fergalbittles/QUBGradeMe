package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Response
type Response struct {
	Error   bool     `json:"error"`
	Average *float64 `json:"average"`
	String  string   `json:"string"`
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())

	e.GET("/", calculateAverage)

	e.Logger.Fatal(e.Start(":1323"))
}

func calculateAverage(c echo.Context) error {
	marksString := map[string]interface{}{}
	marksInt := map[string]int{}

	marksString["Mark 1"] = c.QueryParam("mark_1")
	marksString["Mark 2"] = c.QueryParam("mark_2")
	marksString["Mark 3"] = c.QueryParam("mark_3")
	marksString["Mark 4"] = c.QueryParam("mark_4")
	marksString["Mark 5"] = c.QueryParam("mark_5")

	// Error handling
	for key, value := range marksString {
		// Missing parameter error
		if value == "" {
			msg := key + " value is missing"
			return errorResponse(c, msg, http.StatusBadRequest)
		}

		// Invalid input error
		mark, err := strconv.Atoi(fmt.Sprint(value))
		if err != nil {
			msg := "You must provide a valid integer for " + key
			return errorResponse(c, msg, http.StatusBadRequest)
		}
		if mark < 0 {
			msg := "You must provide a non-negative integer for " + key
			return errorResponse(c, msg, http.StatusBadRequest)
		}
		if mark > 100 {
			msg := "You cannot exceed 100 marks for " + key
			return errorResponse(c, msg, http.StatusBadRequest)
		}

		marksInt[key] = mark
	}

	// Perform calculation
	average, err := averageMarks(marksInt)

	// Return error if calculation was unsuccessful
	if err != nil {
		return errorResponse(c, err.Error(), http.StatusBadRequest)
	}

	// Return response
	response := &Response{
		Error:   false,
		Average: &average,
		String:  "Your average mark is " + fmt.Sprint(average),
	}
	return c.JSON(http.StatusOK, response)
}

func errorResponse(c echo.Context, msg string, statusCode int) error {
	err := &Response{
		Error:  true,
		String: msg,
	}
	return c.JSON(statusCode, err)
}
