package repositories

import (
    "github.com/metrey/hellogo/internal/models"
    "github.com/jmoiron/sqlx"
)

func GetAllUsers(db *sqlx.DB) ([]models.User, error) {
    users := []models.User{}
    err := db.Select(&users, "SELECT id, name, email FROM users")
    return users, err
}
