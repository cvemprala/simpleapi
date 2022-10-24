package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type SimpleRepository struct {
	db *sql.DB
}

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "PickTrace"
	dbname   = "simple"
)

func NewSimpleRepository() (SimpleRepository, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		return SimpleRepository{}, err
	}
	return SimpleRepository{db: db}, nil
}

func (sr SimpleRepository) Create(simple Simple) (string, error) {
	sqlStatement := `INSERT INTO simple(name, birthday, phone, email) VALUES ($1, $2, $3, $4) RETURNING id`
	id := ""
	err := sr.db.QueryRow(sqlStatement, simple.Name, simple.Birthday, simple.Phone, simple.Email).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}
