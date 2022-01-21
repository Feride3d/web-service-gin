/* REST API provides access to the online store of tours. The client can
receive and add tours and their characteristics through endpoints.*/

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Tour structure represents information about tours.
type tour struct {
	ID      string  `json:"id"`
	Title   string  `json:"title"`
	Country string  `json:"country"`
	Price   float64 `json:"price"`
}

// Slice tours adds information about the tours.
var tours = []tour{
	{ID: "1", Title: "Feel Freedom", Country: "Portugal", Price: 600.00},
	{ID: "2", Title: "Unique nature", Country: "Azores", Price: 800.00},
	{ID: "3", Title: "Like Cosmos", Country: "Iceland", Price: 1000.00},
}

/* Endpoint path handler function, establishes a connection,
where getTours handles requests to the endpoint. */
func main() {
	router := gin.Default()
	router.GET("/tours", getTours)
	router.GET("/tours/:id", getTourByID)
	router.POST("/tours", postTours)

	router.Run("localhost:8080")
}

// getTours returns a list of all tours in JSON format.
func getTours(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tours)
}

// postTours adds a tour from the JSON, received in the request body.
func postTours(c *gin.Context) {
	var newTour tour

	// BindJSON to bind the resulting JSON to newTour.
	if err := c.BindJSON(&newTour); err != nil {
		return
	}

	// Add a new tour to the slice.
	tours = append(tours, newTour)
	c.IndentedJSON(http.StatusCreated, newTour)
}

/* getTourByID finds a tour whose id value matches the id parameter sent
by the client,then returns that tour as a response */
func getTourByID(c *gin.Context) {
	id := c.Param("id")

	/* Going through the list of tours in search of a tour,
	whose identifier value matches the parameter. */
	for _, a := range tours {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "tour is not found"})
}
