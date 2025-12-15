package repository

import (
	"database/sql"
	"fmt"
	"r_d/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByID(id int) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow(
		`SELECT id, name, age, phone, is_hidden, rating 
		 FROM users WHERE id = ?`, id,
	).Scan(
		&user.ID, &user.Name, &user.Age,
		&user.Phone, &user.IsHidden, &user.Rating,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}
	if err != nil {
		return nil, fmt.Errorf("error querying user: %w", err)
	}

	return &user, nil
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	rows, err := r.db.Query(
		`SELECT id, name, age, phone, is_hidden, rating FROM users`,
	)
	if err != nil {
		return nil, fmt.Errorf("error querying users: %w", err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID, &user.Name, &user.Age,
			&user.Phone, &user.IsHidden, &user.Rating,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning user: %w", err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return users, nil
}

func (r *UserRepository) Create(user models.User) (int64, error) {
	result, err := r.db.Exec(
		`INSERT INTO users (name, age, phone, is_hidden, rating) 
		 VALUES (?, ?, ?, ?, ?)`,
		user.Name, user.Age, user.Phone, user.IsHidden, user.Rating,
	)
	if err != nil {
		return 0, fmt.Errorf("error inserting user: %w", err)
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error getting last insert id: %w", err)
	}

	fmt.Printf("✅ User created! ID: %d, Name: %s\n", userID, user.Name)
	return userID, nil
}

func (r *UserRepository) Update(user models.User) (int64, error) {
	result, err := r.db.Exec(
		`UPDATE users 
		 SET name=?, age=?, phone=?, is_hidden=?, rating=? 
		 WHERE id=?`,
		user.Name, user.Age, user.Phone, user.IsHidden, user.Rating, user.ID,
	)
	if err != nil {
		return 0, fmt.Errorf("error updating user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("error getting rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return 0, fmt.Errorf("user not found")
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error getting last insert id: %w", err)
	}

	fmt.Printf("✅ User updated! ID: %d, Name: %s\n", user.ID, user.Name)
	return userID, nil
}

func (r *UserRepository) Delete(id int) error {
	result, err := r.db.Exec(`DELETE FROM users WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}
