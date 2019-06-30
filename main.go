package main

import (
	"database/sql"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()
	//r.Use(authMiddleware)
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
	url := "postgres://auriwacs:ZHFxnZyO99adFwMurc3w0JxaQcaAmc3P@satao.db.elephantsql.com:5432/auriwacs"
	db, err := sql.Open("postgres", url)
	//db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT id, name, email, status FROM customer")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	rows, _ := stmt.Query()
	customers := []Customer{}
	for rows.Next() {
		cus := Customer{}
		err := rows.Scan(&cus.ID, &cus.Name, &cus.Email, &cus.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		customers = append(customers, cus)
	}
	c.JSON(http.StatusOK, customers)
}

func postCustomers(c *gin.Context) {
	url := "postgres://auriwacs:ZHFxnZyO99adFwMurc3w0JxaQcaAmc3P@satao.db.elephantsql.com:5432/auriwacs"
	db, err := sql.Open("postgres", url)
	//db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	defer db.Close()

	var customer Customer
	c.BindJSON(&customer)
	query := `INSERT INTO customer (name, email, status) VALUES ($1, $2, $3) RETURNING ID, name, email, status`
	row := db.QueryRow(query, customer.Name, customer.Email, customer.Status)

	err = row.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, customer)
}

func getCustomersByID(c *gin.Context) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT id, name, email, status FROM customer WHERE id=$1;")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var customer Customer
	rows := stmt.QueryRow(c.Param("id"))
	err = rows.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, customer)
}

func updateCustomersByID(c *gin.Context) {
	url := "postgres://auriwacs:ZHFxnZyO99adFwMurc3w0JxaQcaAmc3P@satao.db.elephantsql.com:5432/auriwacs"
	db, err := sql.Open("postgres", url)
	//db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	defer db.Close()

	var customer Customer
	c.BindJSON(&customer)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	customer.ID = id
	err = updateTodo(db, id, customer.Name, customer.Email, customer.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func updateTodo(db *sql.DB, id int, name string, email string, status string) error {
	stmt, err := db.Prepare("UPDATE customer SET name=$2, email=$3, status=$4 WHERE id=$1;")
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(id, name, email, status); err != nil {
		return err
	}

	return err
}

func deleteCustomersByID(c *gin.Context) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	defer db.Close()

	query := `DELETE FROM customer WHERE id=$1 RETURNING ID, name, email, status;`
	stmt, err := db.Prepare(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var customer Customer
	rows := stmt.QueryRow(c.Param("id"))
	err = rows.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "customer deleted"})
}

func authMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != "token123" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		c.Abort()
		return
	}
	c.Next()
}
