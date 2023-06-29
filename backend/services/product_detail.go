package services

import (
	"strconv"

	"commercial-shop.com/database"
	"commercial-shop.com/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductDetailService struct {
	Items []models.ProductDetail
}

func (sv *ProductDetailService) Create() error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// SQL commamd
	sql := "INSERT INTO ProductDetail VALUES (@id, @idProduct, @color, @fabric, @size, @price, @amount, @description);"
	args := pgx.NamedArgs{
		"id":          sv.Items[0].Id,
		"idProduct":   sv.Items[0].IdProduct,
		"color":       sv.Items[0].Color,
		"fabric":      sv.Items[0].Fabric,
		"size":        sv.Items[0].Size,
		"price":       sv.Items[0].Price,
		"amount":      sv.Items[0].Amount,
		"description": sv.Items[0].Description,
	}

	// Execute sql command
	_, err = conn.Exec(database.CTX, sql, args)
	if err != nil {
		return err
	}

	return nil
}

func (sv *ProductDetailService) Delete() error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// SQL commamd
	sql := "DELETE FROM ProductDetail WHERE product_detail_id='" + sv.Items[0].Id + "';"

	// Execute sql command
	_, err = conn.Exec(database.CTX, sql)
	if err != nil {
		return err
	}

	return nil
}

func (sv *ProductDetailService) Get() error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// SQL commamd
	sql := "SELECT * FROM ProductDetail WHERE product_detail_id='" + sv.Items[0].Id + "';"

	// Get rows from conn with SQL command
	err = conn.QueryRow(database.CTX, sql).Scan(
		&sv.Items[0].Id,
		&sv.Items[0].IdProduct,
		&sv.Items[0].Color,
		&sv.Items[0].Fabric,
		&sv.Items[0].Size,
		&sv.Items[0].Price,
		&sv.Items[0].Amount,
		&sv.Items[0].Description)
	if err != nil {
		return err
	}

	return nil
}

func (sv *ProductDetailService) GetAll(limit *int, page *int) error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// SQL commamd
	sql := "SELECT * FROM ProductDetail ORDER BY product_detail_id LIMIT @limit OFFSET @offset;"
	args := pgx.NamedArgs{
		"limit":  strconv.Itoa(*limit),
		"offset": strconv.Itoa(*limit * (*page - 1)),
	}

	// Get rows from conn with SQL command
	rows, err := conn.Query(database.CTX, sql, args)
	if err != nil {
		return err
	}

	// convert each rows to struct and append to Slice to return
	i := 0
	for rows.Next() {
		sv.Items = append(sv.Items, models.ProductDetail{})

		err := rows.Scan(
			&sv.Items[i].Id,
			&sv.Items[i].IdProduct,
			&sv.Items[i].Color,
			&sv.Items[i].Fabric,
			&sv.Items[i].Size,
			&sv.Items[i].Price,
			&sv.Items[i].Amount,
			&sv.Items[i].Description,
		)
		if err != nil {
			return err
		}

		i++
	}
	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func (sv *ProductDetailService) Update() error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// SQL commamd
	sql := "UPDATE ProductDetail SET product_id=@idProduct, product_color=@color, product_fabric=@fabric, product_size=@size, product_price=@price, product_amount=@amount, product_description=@description WHERE product_detail_id=@id;"
	args := pgx.NamedArgs{
		"id":          sv.Items[0].Id,
		"idProduct":   sv.Items[0].IdProduct,
		"color":       sv.Items[0].Color,
		"fabric":      sv.Items[0].Fabric,
		"size":        sv.Items[0].Size,
		"price":       sv.Items[0].Price,
		"amount":      sv.Items[0].Amount,
		"description": sv.Items[0].Description,
	}

	// Execute sql command
	_, err = conn.Exec(database.CTX, sql, args)
	if err != nil {
		return err
	}

	return nil
}
