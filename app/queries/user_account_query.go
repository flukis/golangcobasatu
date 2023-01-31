package queries

import (
	"github.com/fahmilukis/expense-tracker/app/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// struct query for expense category
type UserAccountQuery struct {
	*sqlx.DB
}

// get by id
func (q *UserAccountQuery) GetAccountById(id uuid.UUID) (models.User, error) {
	// define variable
	user := models.User{}

	// define query string
	query := `SELECT * FROM account WHERE id = $1`

	err := q.Get(&user, query, id)
	if err != nil {
		return user, err
	}

	return user, nil
}

// get by id
func (q *UserAccountQuery) GetAccountByEmail(email string) (models.User, error) {
	// define variable
	user := models.User{}

	// define query string
	query := `SELECT * FROM users WHERE email = $1`

	err := q.Get(&user, query, email)
	if err != nil {
		return user, err
	}

	return user, nil
}

// get all
func (q *UserAccountQuery) GetAccounts() ([]models.User, error) {
	// define variable
	user := []models.User{}

	// define query string
	query := `SELECT * from users`

	err := q.Get(&user, query)
	if err != nil {
		return user, err
	}

	return user, nil
}

// create account
func (q *UserAccountQuery) CreateAccount(a *models.User) error {
	query := `INSERT INTO users VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := q.Exec(
		query,
		a.ID, a.CreatedAt, a.UpdatedAt, a.Email, a.Name, a.Password,
	)

	if err != nil {
		return err
	}

	return nil
}

// update account
func (q *UserAccountQuery) UpdateAccount(id uuid.UUID, a *models.User) error {
	query := `UPDATE users SET
		id = $1,
		email = $2,
		password = $3,
		name = $4,
		created_at = $5,
		updated_at = $6,
	`

	_, err := q.Exec(
		query,
		id,
		a.Email,
		a.Password,
		a.Name,
		a.CreatedAt,
		a.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

// DeleteBook method for delete book by given ID.
func (q *UserAccountQuery) DeleteUser(id uuid.UUID) error {
	// Define query string.
	query := `DELETE FROM users WHERE id = $1`

	// Send query to database.
	_, err := q.Exec(query, id)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}
