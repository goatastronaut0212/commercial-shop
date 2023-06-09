package services

import (
	"strconv"

	"commercial-shop.com/database"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Product struct {
	Id         string `json:"id"         binding:"required"`
	IdCategory string `json:"idCategory" binding:"required"`
	Name       string `json:"name"       binding:"required"`
}

type ProductService struct {
	Items []Product
}

func (sv *ProductService) Get() error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := "SELECT * FROM Product WHERE product_id='" + sv.Items[0].Id + "';"

	// Get rows from conn with SQL command
	err = conn.QueryRow(database.CTX, sql).Scan(&sv.Items[0].Id, &sv.Items[0].IdCategory, &sv.Items[0].Name)
	if err != nil {
		return err
	}

	return nil
}

func (sv *ProductService) GetAll(limit *int, page *int) error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := "SELECT * FROM product ORDER BY product_id LIMIT @limit OFFSET @offset;"
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
		sv.Items = append(sv.Items, Product{})
		err := rows.Scan(&sv.Items[i].Id, &sv.Items[i].IdCategory, &sv.Items[i].Name)
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

func (sv *ProductService) Create() error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := "INSERT INTO Product VALUES (@id, @idCategory, @name);"
	args := pgx.NamedArgs{
		"id":         sv.Items[0].Id,
		"idCategory": sv.Items[0].IdCategory,
		"name":       sv.Items[0].Name,
	}

	// Execute sql command
	_, err = conn.Exec(database.CTX, sql, args)
	if err != nil {
		return err
	}

	return nil
}

func (sv *ProductService) Update() error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := ""
	args := pgx.NamedArgs{}
	if sv.Items[0].Name == "" {
		sql = "UPDATE Product SET category_id=@idCategory WHERE product_id=@id;"
		args = pgx.NamedArgs{
			"id":         sv.Items[0].Id,
			"idCategory": sv.Items[0].IdCategory,
		}
	} else {
		sql = "UPDATE Product SET category_id=@idCategory, product_name=@name WHERE product_id=@id;"
		args = pgx.NamedArgs{
			"id":         sv.Items[0].Id,
			"idCategory": sv.Items[0].IdCategory,
			"name":       sv.Items[0].Name,
		}
	}

	// Execute sql command
	_, err = conn.Exec(database.CTX, sql, args)
	if err != nil {
		return err
	}

	return nil
}

func (sv *ProductService) Delete() error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := "DELETE FROM Product WHERE product_id='" + sv.Items[0].Id + "';"

	// Execute sql command
	_, err = conn.Exec(database.CTX, sql)
	if err != nil {
		return err
	}

	return nil
}
