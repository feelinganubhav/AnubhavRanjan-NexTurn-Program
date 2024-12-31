package repository

import (
	"database/sql"
	"go-blog-management-system/model"
	"log"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// ----------- User Repository Methods ----------- //

func (repo *UserRepository) CreateUser(user *model.User) (*model.User, error) {
	stmt, err := repo.DB.Prepare("INSERT INTO users (name, email, password) VALUES (?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(user.Name, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.ID = int(id)
	return user, nil
}

func (repo *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	row := repo.DB.QueryRow("SELECT id, name, email, password FROM users WHERE email = ?", email)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetUser(id int) (*model.User, error) {
	row := repo.DB.QueryRow("SELECT id, name, email, password FROM users WHERE id = ?", id)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetAllUsers() ([]model.User, error) {
	rows, err := repo.DB.Query("SELECT id, name, email, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (repo *UserRepository) UpdateUser(user *model.User) (*model.User, error) {
	stmt, err := repo.DB.Prepare("UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.Password, user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRepository) DeleteUser(id int) error {
	stmt, err := repo.DB.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
