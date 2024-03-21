package controller

import (
	"database/sql"
	"log"
	"martini/model"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/martini-contrib/render"
)

func GetAllUsers(r render.Render, db *sql.DB) {
	rows, err := db.Query("SELECT `ID`, `Name`, `Age`, `Address`, `Email` FROM `users`;")
	if err != nil {
		log.Println(err)
		sendErrorResponse(r, 500, "error while querying to database")
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.Email); err != nil {
			log.Println(err)
			sendErrorResponse(r, 500, "error whiel retrieving data")
			return
		} else {
			users = append(users, user)
		}
	}

	sendSuccessResponseWithData(r, users)
}
func NewUser(r render.Render, body model.User, db *sql.DB) {
	if body.Name == "" {
		sendErrorResponse(r, 400, "name required")
		return
	}
	if body.Age <= 0 {
		sendErrorResponse(r, 400, "invalid age")
		return
	}
	if body.Address == "" {
		sendErrorResponse(r, 400, "address required")
		return
	}
	if body.Email == "" {
		sendErrorResponse(r, 400, "email required")
		return
	}
	if body.Password == "" {
		sendErrorResponse(r, 400, "password required")
		return
	}

	_, err := db.Exec("INSERT INTO users (name, age, address, email, password, usertype) VALUES (?, ?, ?, ?, sha1(?), 0);",
		body.Name, body.Age, body.Address, body.Email, body.Password,
	)

	if err != nil {
		log.Println(err)
		sendErrorResponse(r, 510, "unable to add new user")
		return
	}

	sendSuccessResponse(r, "success")
}
func EditUser(r render.Render, body model.User, db *sql.DB) {
	if body.ID <= 0 {
		sendErrorResponse(r, 400, "id required")
		return
	}

	if body.Name == "" && body.Age <= 0 && body.Address == "" && body.Email == "" && body.Password == "" {
		sendSuccessResponse(r, "no changes")
		return
	}

	query := "UPDATE `users` SET"

	if body.Name != "" {
		query += " `name`='" + body.Name + "',"
	}
	if body.Age > 0 {
		query += " `age`='" + strconv.Itoa(body.Age) + "',"
	}
	if body.Address != "" {
		query += " `address`='" + body.Address + "',"
	}
	if body.Email != "" {
		query += " `email`='" + body.Email + "',"
	}
	if body.Password != "" {
		query += " `password`=sha1('" + body.Password + "'),"
	}

	query = query[:len(query)-1] + " WHERE `id`=" + strconv.Itoa(body.ID)

	_, err := db.Exec(query)

	if err != nil {
		log.Println(err)
		sendErrorResponse(r, 500, "unable to update user")
		return
	}

	sendSuccessResponse(r, "user updated successfully")
}
func DeleteUser(r render.Render, body model.User, db *sql.DB) {
	if body.ID <= 0 {
		sendErrorResponse(r, 400, "id required")
		return
	}

	_, err := db.Exec("DELETE FROM users WHERE id=?", body.ID)
	if err != nil {
		sendErrorResponse(r, 500, "unable to delete user")
		return
	}

	sendSuccessResponse(r, "user deleted successfully")
}

func sendSuccessResponseWithData(r render.Render, data interface{}) {
	var response model.BasicResponseWithData
	response.Status = 200
	response.Message = "success"
	response.Data = data
	r.JSON(http.StatusOK, response)
}

func sendSuccessResponse(r render.Render, message string) {
	var response model.BasicResponse
	response.Status = 200
	response.Message = message
	r.JSON(http.StatusOK, response)
}

func sendErrorResponse(r render.Render, status int, message string) {
	var response model.ErrorResponse
	response.Status = status
	response.Message = message
	r.JSON(status, response)
}
