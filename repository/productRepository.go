package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(conn *sql.DB) ProductRepository {
	return ProductRepository{
		connection: conn,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	querry := `SELECT id,
	 				  product_name,
					  price
				 FROM product`

	rows, err := pr.connection.Query(querry)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.Id,
			&productObj.Name,
			&productObj.Price,
		)
		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	var id int

	query, err := pr.connection.Prepare(
		`INSERT into product (product_name, price)
		 VALUES ($1, $2)
	  RETURNING id`,
	)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
	}

	query.Close()

	return id, nil
}

func (pr *ProductRepository) GetProductById(idProduct int) (*model.Product, error) {
	var product model.Product

	query, err := pr.connection.Prepare(
		`SELECT id,
				product_name,
				price
		   FROM product
		  WHERE id = $1`,
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = query.QueryRow(idProduct).Scan(
		&product.Id,
		&product.Name,
		&product.Price,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()

	return &product, nil
}
