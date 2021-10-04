package handlers

import (
	"database/sql"
)

type (
	Handler struct {
		DB *sql.DB
	}

	Response struct {
		Message string
		Code    uint64
	}
)

const (
	//Host   = "localhost"
	//Port   = "3306"
	User   = "root"
	Pass   = "oxford24"
	DBName = "project"
)
