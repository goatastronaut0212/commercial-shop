package access

import (
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"commercial-shop.com/database"
	"commercial-shop.com/models"
)

func FindAllProductDetail(limit *int, page *int) ([]models.ProductDetail, error) {
	dataSlice := []models.ProductDetail{}
	data := &models.ProductDetail{}

	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := "SELECT product_detail_id, product_id, product_color, product_fabric, product_size, product_price, product_amount, product_description FROM product_detail ORDER BY product_detail_id LIMIT @limit OFFSET @offset;"
	args := pgx.NamedArgs{
		"limit":  strconv.Itoa(*limit),
		"offset": strconv.Itoa(*limit * (*page - 1)),
	}

	// Get rows from conn with SQL command
	rows, err := conn.Query(database.CTX, sql, args)
	if err != nil {
		return nil, err
	}

	// convert each rows to struct and append to Slice to return
	for rows.Next() {
		err := rows.Scan(&data.Id, &data.IdProduct, &data.Color, &data.Fabric, &data.Size, &data.Price, &data.Amount, &data.Description)
		if err != nil {
			return nil, err
		}
		dataSlice = append(dataSlice, *data)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return dataSlice, nil
}

func FindProductDetail(id *string) (models.ProductDetail, error) {
	data := &models.ProductDetail{}

	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return *data, err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := "SELECT * FROM product_detail WHERE product_detail_id='" + *id + "';"

	// Get rows from conn with SQL command
	err = conn.QueryRow(database.CTX, sql).Scan(&data.Id, &data.IdProduct, &data.Color, &data.Fabric, &data.Size, &data.Price, &data.Amount, &data.Description)
	if err != nil {
		return *data, err
	}

	return *data, nil
}

func CreateProductDetail(data *models.ProductDetail) error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := "INSERT INTO product_detail VALUES (@id, @idProduct, @color, @fabric, @size, @price, @amount, @description);"
	args := pgx.NamedArgs{
		"id":          data.Id,
		"idProduct":   data.IdProduct,
		"color":       data.Color,
		"fabric":      data.Fabric,
		"size":        data.Size,
		"price":       data.Price,
		"amount":      data.Amount,
		"description": data.Description,
	}

	// Execute sql command
	_, err = conn.Exec(database.CTX, sql, args)
	if err != nil {
		return err
	}

	return nil
}

func UpdateProductDetail(data *models.ProductDetail) error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := "UPDATE product_detail SET product_id=@idProduct, product_color=@color, product_fabric=@fabric, product_size=@size, product_price=@price, product_amount=@amount, product_description=@description WHERE product_detail_id=@id;"
	args := pgx.NamedArgs{
		"id":          data.Id,
		"idProduct":   data.IdProduct,
		"color":       data.Color,
		"fabric":      data.Fabric,
		"size":        data.Size,
		"price":       data.Price,
		"amount":      data.Amount,
		"description": data.Description,
	}

	// Execute sql command
	_, err = conn.Exec(database.CTX, sql, args)
	if err != nil {
		return err
	}

	return nil
}

func DeleteProductDetail(id *string) error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := "DELETE FROM product_detail WHERE product_detail_id='" + *id + "';"

	// Execute sql command
	_, err = conn.Exec(database.CTX, sql)
	if err != nil {
		return err
	}

	return nil
}
