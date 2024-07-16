package controllers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"user-management-api/models"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
)

var db *pgxpool.Pool

// InitDB initializes the database connection
func InitDB(dataSourceName string) {
	var err error
	db, err = pgxpool.Connect(context.Background(), dataSourceName)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	log.Println("Connected to the database successfully")
}

// GetUsers godoc
// @Summary Get list of users
// @Description Get list of users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User
// @Router /users [get]
func GetUsers(c echo.Context) error {
	rows, err := db.Query(context.Background(), "SELECT user_id, user_name, first_name, last_name, email, user_status, department FROM users")
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		return c.JSON(http.StatusInternalServerError, "Error fetching users")
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.UserID, &user.UserName, &user.FirstName, &user.LastName, &user.Email, &user.UserStatus, &user.Department)
		if err != nil {
			log.Printf("Error scanning user: %v", err)
			return c.JSON(http.StatusInternalServerError, "Error fetching users")
		}
		users = append(users, user)
	}
	return c.JSON(http.StatusOK, users)
}

// GetUserById godoc
// @Summary Get a user by ID
// @Description Get a user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param   user_id  path   int     true  "User ID"
// @Success 200 {object} models.User
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /users/{user_id} [get]
func GetUserById(c echo.Context) error {
	id := c.Param("user_id")

	// Parse the user ID from the route parameter
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	// Query to fetch user details by ID
	query := `
		SELECT user_id, user_name, first_name, last_name, email, user_status, department
		FROM users
		WHERE user_id = $1
	`

	// Execute the query and retrieve the user details
	var user models.User
	row := db.QueryRow(context.Background(), query, userID)
	err = row.Scan(&user.UserID, &user.UserName, &user.FirstName, &user.LastName, &user.Email, &user.UserStatus, &user.Department)
	if err != nil {
		if err.Error() == "pgx: no rows in result set" {
			return c.JSON(http.StatusNotFound, "User not found")
		}
		log.Printf("Error fetching user: %v", err)
		return c.JSON(http.StatusInternalServerError, "Error fetching user")
	}

	return c.JSON(http.StatusOK, user)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param   user  body   models.User  true  "User"
// @Success 201 {object} models.User
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /users [post]
func CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	query := "INSERT INTO users (user_name, first_name, last_name, email, user_status, department) VALUES ($1, $2, $3, $4, $5, $6) RETURNING user_id"
	err := db.QueryRow(context.Background(), query, user.UserName, user.FirstName, user.LastName, user.Email, user.UserStatus, user.Department).Scan(&user.UserID)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return c.JSON(http.StatusInternalServerError, "Error creating user")
	}

	return c.JSON(http.StatusCreated, user)
}

// UpdateUser godoc
// @Summary Update an existing user
// @Description Update an existing user
// @Tags users
// @Accept  json
// @Produce  json
// @Param   user_id  path   int           true  "User ID"
// @Param   user     body   models.User   true  "User"
// @Success 200 {object} models.User
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /users/{user_id} [put]
func UpdateUser(c echo.Context) error {
	id := c.Param("user_id")
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}
	query := "UPDATE users SET user_name=$1, first_name=$2, last_name=$3, email=$4, user_status=$5, department=$6 WHERE user_id = $7"
	_, err := db.Exec(context.Background(), query, user.UserName, user.FirstName, user.LastName, user.Email, user.UserStatus, user.Department, id)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return c.JSON(http.StatusInternalServerError, "Error updating user")
	}

	return c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user
// @Tags users
// @Accept  json
// @Produce  json
// @Param   user_id  path   int     true  "User ID"
// @Success 204
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /users/{user_id} [delete]
func DeleteUser(c echo.Context) error {
	id := c.Param("user_id")

	query := "DELETE FROM users WHERE user_id = $1"
	_, err := db.Exec(context.Background(), query, id)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return c.JSON(http.StatusInternalServerError, "Error deleting user")
	}

	return c.NoContent(http.StatusNoContent)
}
