package repository

import (
	"go-api/model"
	"database/sql"
	"fmt"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {

	query := "SELECT id, product_name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return []model.Product{}, err
	}

	var productsList []model.Product
	var productbj model.Product

	for rows.Next() {
		err := rows.Scan(
			&productbj.ID,
			&productbj.Name,
			&productbj.Price,
		)

		if err != nil {
			fmt.Println("Error scanning row:", err)
			return []model.Product{}, err
		}
		
		productsList = append(productsList, productbj)
	}
	
	rows.Close()

	return productsList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product)  (int, error) {
	
	var id int

	query, err := pr.connection.Prepare("INSERT INTO product" +
	"(product_name, price)" +
	"VALUES ($1, $2) Returning id")
	
	if err != nil {
		fmt.Println("Error preparing query:", err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	
	if err != nil {
		fmt.Println("Error preparing query:", err)
		return 0, err
	}
	query.Close()
	
	return id, nil
}

func (pr *ProductRepository) GetProductByID(id int) (*model.Product, error) {
	var product model.Product
	query, err := pr.connection.Prepare("SELECT * FROM product WHERE id = $1")
	if err != nil {
		return nil, err
	}
	
	err = query.QueryRow(id).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No product found with the given ID
		}
		return nil, err
	}
	
	query.Close()
	return &product, nil
}

func (pr *ProductRepository) UpdateProduct(product model.Product) (*model.Product, error) {
	query, err := pr.connection.Prepare("UPDATE product SET product_name = $1, price = $2 WHERE id = $3")
	if err != nil {
		return nil, err
	}


	_, err = query.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return nil, err
	}

	defer query.Close()

	return &product, nil
}