package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()
	r.Use(authMiddleware)
	r.GET("/customers", getCustomers)
	r.GET("/customers/:id", getCustomersByID)
	r.POST("/customers/", postCustomers)
	r.PUT("/customers/:id", updateCustomersByID)
	r.DELETE("/customers/:id", deleteCustomersByID)
	r.Run(":2019")
}

type Customer struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

func getCustomers(c *gin.Context) {
	url := "postgres://fqwnvlfk:lv3nDmkzmBXgk6dup77dO6CbsjcJa2-L@satao.db.elephantsql.com:5432/fqwnvlfk"
	db, err := sql.Open("postgres", url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "not connect"})
	}
	stmt, err := db.Prepare("SELECT id, name, email, status FROM customer")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "not query data"})
	}
	rows, _ := stmt.Query()
	customers := []Customer{}
	for rows.Next() {
		cus := Customer{}
		err := rows.Scan(&cus.ID, &cus.Name, &cus.Email, &cus.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "not data"})
		}
		customers = append(customers, cus)
	}
	c.JSON(200, customers)
}

func postCustomers(c *gin.Context) {
	url := "postgres://fqwnvlfk:lv3nDmkzmBXgk6dup77dO6CbsjcJa2-L@satao.db.elephantsql.com:5432/fqwnvlfk"
	//db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	db, err := sql.Open("postgres", url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "not connect"})
	}

	name := "top"
	email := "suttipong.kull@gmail.com"
	status := "inactive"

	var customer Customer
	c.BindJSON(&customer)
	query := `INSERT INTO customer (name, email, status) VALUES ($1, $2, $3) RETURNING ID, name, email, status`
	row := db.QueryRow(query, name, email, status)

	err = row.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, customer)
}

func getCustomersByID(c *gin.Context) {
	url := "postgres://fqwnvlfk:lv3nDmkzmBXgk6dup77dO6CbsjcJa2-L@satao.db.elephantsql.com:5432/fqwnvlfk"
	//db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	db, err := sql.Open("postgres", url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "not connect"})
	}

	stmt, err := db.Prepare("SELECT id, name, email, status FROM customer WHERE id=$1;")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var customer Customer
	rows := stmt.QueryRow(1)
	err = rows.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(200, customer)
}

func updateCustomersByID(c *gin.Context) {
	url := "postgres://fqwnvlfk:lv3nDmkzmBXgk6dup77dO6CbsjcJa2-L@satao.db.elephantsql.com:5432/fqwnvlfk"
	//db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	db, err := sql.Open("postgres", url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "not connect"})
	}

	var customer Customer

	c.BindJSON(&customer)

	name := "suttipong"
	email := "suttipong.kull@gmail.com"
	status := "active"

	query := `UPDATE customer SET name=$2, email=$3, status=$4 WHERE id=$1;`
	stmt, err := db.Prepare(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	rows := stmt.QueryRow(query, name, email, status)
	err = rows.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	//customer.ID = id
	c.JSON(200, customer)
}

func deleteCustomersByID(c *gin.Context) {
	url := "postgres://fqwnvlfk:lv3nDmkzmBXgk6dup77dO6CbsjcJa2-L@satao.db.elephantsql.com:5432/fqwnvlfk"
	//db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	db, err := sql.Open("postgres", url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "not connect"})
	}

	query := `DELETE FROM customer WHERE id=$1 RETURNING ID, name, email, status;`
	stmt, err := db.Prepare(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var customer Customer
	rows := stmt.QueryRow(1)
	err = rows.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{"message": "customer deleted"})
}

func authMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != "Bearer token123" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		c.Abort()
		return
	}
	c.Next()
}
