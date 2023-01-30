package queries

import (
	"github.com/fahmilukis/expense-tracker/app/models"
	"github.com/jmoiron/sqlx"
)

// struct query for expense category
type ExpenseCategoryQuery struct {
	*sqlx.DB
}

// get by id
func (q *ExpenseCategoryQuery) GetCategoryById(id string) (models.ExpenseCategory, error) {
	// define variable
	category := models.ExpenseCategory{}

	// define query string
	query := `SELECT * FROM expense_categories WHERE id = $1`

	err := q.Get(&category, query, id)
	if err != nil {
		return category, err
	}

	return category, nil
}

// get all
func (q *ExpenseCategoryQuery) GetCategories() ([]models.ExpenseCategory, error) {
	// define variable
	categories := []models.ExpenseCategory{}

	// define query string
	query := `SELECT * from expense_categories`

	err := q.Get(&categories, query)
	if err != nil {
		return categories, err
	}

	return categories, nil
}
