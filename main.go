package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type employee struct {
    ID     string  `json:"id"`
    Name  string  `json:"name"`
    LastName string  `json:"lastName"`
    Salary  float64 `json:"salary"`
}

var employees = []employee{
    {ID: "1", Name: "Natan", LastName: "Silva", Salary: 100.00},
    {ID: "2", Name: "Flavio", LastName: "De Moraes", Salary: 1700.00},
    {ID: "3", Name: "Marcelo", LastName: "Amaral", Salary: 1139.00},
}

var hello = "Deu tudo certo no deploy do heroku"

func main() {
    router := gin.Default()
    router.GET("/employees", getEmployees)
    router.GET("/employees/:id", getEmployeeById)
    router.POST("/employees", postEmployees)
	router.GET("/", getPath)

    router.Run("localhost:8080")
}

func getPath(c *gin.Context) {
	c.IndentedJSON(http.StatusOK,hello)
}

func getEmployees(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, employees)
}

func postEmployees(c *gin.Context) {
    var newEmployee employee

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newEmployee); err != nil {
        return
    }

    // Add the new album to the slice.
    employees = append(employees, newEmployee)
    c.IndentedJSON(http.StatusCreated, newEmployee)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getEmployeeById(c *gin.Context) {
    id := c.Param("id")

    // Loop through the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range employees {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Employee not found"})
}