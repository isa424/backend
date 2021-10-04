package handlers

import (
	"fmt"
	model "github.com/isa424/backend/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

func (h *Handler) GetUser(c echo.Context) (err error) {
	userId := c.Param("user_id")
	user := &model.User{}

	query := "SELECT user_id, username, email FROM users WHERE user_id=?"
	stmt, err := h.DB.Prepare(query)
	defer stmt.Close()

	if err != nil {
		log.Error(err)
		return err
	}

	err = stmt.QueryRow(userId).Scan(&user.ID, &user.Username, &user.Email)

	if err != nil {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, user)
}

func (h *Handler) GetUsers(c echo.Context) (err error) {
	query := "SELECT user_id, username, email FROM users"
	rows, err := h.DB.Query(query)
	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var user model.User

		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			return err
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func (h *Handler) CreateUser(c echo.Context) (err error) {
	var user model.User

	if err = c.Bind(&user); err != nil {
		return err
	}

	query := "INSERT INTO users (username, email, password) VALUES (?, ?, 'secret')"
	stmt, err := h.DB.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	exec, err := stmt.Exec(user.Username, user.Email)
	if err != nil {
		return err
	}

	id, err := exec.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = uint64(id)
	fmt.Println(user)

	return c.JSON(http.StatusCreated, user)
}

func (h *Handler) UpdateUser(c echo.Context) (err error) {
	var user model.User
	userId := c.Param("user_id")

	fmt.Println(userId)

	if err = c.Bind(&user); err != nil {
		return err
	}

	fmt.Println(user)
	query := "UPDATE users SET username=?, email=? WHERE user_id=?"
	stmt, err := h.DB.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	exec, err := stmt.Exec(&user.Username, &user.Email, userId)
	if err != nil {
		return err
	}

	count, err := exec.RowsAffected()
	if err != nil {
		return err
	}

	if count < 1 {
		return c.JSON(http.StatusNotFound, Response{Message: "User not found!", Code: 404})
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) DeleteUser(c echo.Context) (err error) {
	userId := c.Param("user_id")

	stmt, err := h.DB.Prepare("DELETE FROM users WHERE user_id=?")
	defer stmt.Close()
	if err != nil {
		return err
	}

	exec, err := stmt.Exec(userId)
	if err != nil {
		return err
	}

	count, err := exec.RowsAffected()
	if err != nil {
		return err
	}

	if count < 1 {
		return c.JSON(http.StatusNotFound, Response{Message: "User not found!", Code: 404})
	}

	return c.NoContent(http.StatusNoContent)
}
