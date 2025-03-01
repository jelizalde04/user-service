package services

import (
    "database/sql"
    "errors"
)

type UserService struct {
    DB *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
    return &UserService{DB: db}
}

func (s *UserService) DeleteUser(userID int) error {
    result, err := s.DB.Exec("DELETE FROM users WHERE id = $1", userID)
    if err != nil {
        return errors.New("error al eliminar usuario")
    }

    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        return errors.New("usuario no encontrado")
    }

    return nil
}
