package repos

import (
	"database/sql"
	"errors"
	"invsvc/types"
)

type ProductRepoInterface interface {
	GetProduct() ([]*types.Product, error)
	GetById(string) (*types.Product, error)
}

type ProductRepo struct {
	db *sql.DB
}

func NewProducRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (p *ProductRepo) GetProduct() ([]*types.Product, error) {
	var products = []*types.Product{}
	query := "SELECT * FROM Product"

	prep, err := p.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := prep.Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		item := &types.Product{}

		err = rows.Scan(
			&item.Id,
			&item.ProductId,
			&item.ProductName,
			&item.Description,
			&item.Tag,
			&item.Image,
			&item.Price,
			&item.InStock,
			&item.Ingredients,
			&item.CreateAt,
			&item.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, item)

	}

	if len(products) < 0 {
		return nil, errors.New("Products not found")
	}

	return products, nil
}

func (p *ProductRepo) GetById(id string) (*types.Product, error) {
	var product = &types.Product{}
	query := "SELECT * FROM Product WHERE id = ?"

	prep, err := p.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := prep.Query(id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		item := &types.Product{}

		err = rows.Scan(
			&item.Id,
			&item.ProductId,
			&item.ProductName,
			&item.Description,
			&item.Tag,
			&item.Image,
			&item.Price,
			&item.InStock,
			&item.Ingredients,
			&item.CreateAt,
			&item.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		product = item

	}

	if product.Id == 0 {
		return nil, errors.New("Product not found ")
	}

	return product, nil
}

func (p *ProductRepo) CreateProduct(product types.Product) (bool, error) {
	insertQuery := `INSERT INTO Product (product_name, description, tag, image, price, in_stock, ingredients) 
	VALUES ($1, $2, $3, $4, $5, $6, $7)`

	rows, err := p.db.Exec(insertQuery, product)
	if err != nil {
		return false, err
	}

	affected, err := rows.RowsAffected()
	if err != nil {
		return false, err
	}

	if affected == 0 {
		return false, errors.New("Nothing was created")
	}

	return true, nil
}
