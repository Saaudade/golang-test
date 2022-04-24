package database

import "database/sql"

type Repository interface {
	FindCurrentDiscount() int
}

type DiscountRepository struct {
	db *sql.DB
}

func NewDiscountRepository(db *sql.DB) *DiscountRepository {
	return &DiscountRepository{
		db: db,
	}
}

func (dr *DiscountRepository) FindCurrentDiscount() (discount int) {
	sql := "SELECT value FROM discount ORDER BY RAND() LIMIT 1"
	row := dr.db.QueryRow(sql)
	row.Scan(&discount)
	return discount
}
