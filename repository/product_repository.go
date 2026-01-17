package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
	"strings"
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
	sql := "SELECT id, product_name, price FROM product ORDER BY id ASC"
	rows, err := pr.connection.Query(sql)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
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
	stmt, err := pr.connection.Prepare("INSERT INTO product" +
		"(product_name, price)" +
		" VALUES ($1, $2) RETURNING ID;")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = stmt.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	stmt.Close()
	return id, nil
}

func (pr *ProductRepository) GetProductById(productId int) (*model.Product, error) {
	stmt, err := pr.connection.Prepare("SELECT id, product_name, price FROM product WHERE id=$1;")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var product model.Product

	err = stmt.QueryRow(productId).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		fmt.Println(err)
		return nil, err
	}

	stmt.Close()
	return &product, nil
}

func (pr *ProductRepository) UpdateProduct(product model.Product) (*model.Product, error) {
	var setClauses []string
	var args []any

	argId := 1

	if product.Name != "" {
		setClauses = append(setClauses, fmt.Sprintf("product_name=$%d", argId))
		args = append(args, product.Name)
		argId++
	}

	if product.Price != 0 {
		setClauses = append(setClauses, fmt.Sprintf("price=$%d", argId))
		args = append(args, product.Price)
		argId++
	}

	if len(setClauses) == 0 {
		return &product, nil
	}

	setQuery := strings.Join(setClauses, ", ")

	sql := fmt.Sprintf("UPDATE product SET %s WHERE id=$%d RETURNING *", setQuery, argId)

	args = append(args, product.ID)

	var updatedProduct model.Product

	stmt, err := pr.connection.Prepare(sql)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	updateErr := stmt.QueryRow(args...).Scan(&updatedProduct.ID, &updatedProduct.Name, &updatedProduct.Price)
	if updateErr != nil {
		fmt.Println(updateErr)
		return nil, updateErr
	}

	return &updatedProduct, nil
}

func (pr *ProductRepository) DeleteProduct(productId int) error {
	sql := fmt.Sprintf("DELETE FROM product WHERE id=%d RETURNING *", productId)

	_, err := pr.connection.Query(sql)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
