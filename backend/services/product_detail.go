package services

import (
	"fmt"
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
	sql := "INSERT INTO ProductDetail VALUES (@id, @productid, @color, @fabric, @size, @form, @price, @amount, @description);"
	args := pgx.NamedArgs{
		"id":          sv.Items[0].Id,
		"productid":   sv.Items[0].ProductId,
		"color":       sv.Items[0].Color,
		"fabric":      sv.Items[0].Fabric,
		"size":        sv.Items[0].Size,
		"form":        sv.Items[0].Form,
		"price":       sv.Items[0].Price,
		"amount":      sv.Items[0].Amount,
		"description": sv.Items[0].Description,
	}

	// Execute sql command
	_, err = conn.Exec(database.CTX, sql, args)
	if err != nil {
		fmt.Println(err)
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
		&sv.Items[0].ProductId,
		&sv.Items[0].Color,
		&sv.Items[0].Fabric,
		&sv.Items[0].Size,
		&sv.Items[0].Form,
		&sv.Items[0].Price,
		&sv.Items[0].Amount,
		&sv.Items[0].Description)
	if err != nil {
		return err
	}

	return nil
}

func (sv *ProductDetailService) GetAll(limit, page *int) error {
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
			&sv.Items[i].ProductId,
			&sv.Items[i].Color,
			&sv.Items[i].Fabric,
			&sv.Items[i].Size,
			&sv.Items[i].Form,
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

func (sv *ProductDetailService) Update(productid, color, fabric, size, form, amount, price, description *bool) error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// SQL commamd
	sql := "UPDATE ProductDetail SET "
	args := pgx.NamedArgs{"id": sv.Items[0].Id}
	nextoption := true

	for i := 0; i < 8; i++ {
		switch {
		case *productid == true:
			sql += "product_id=@product_id"
			args["product_id"] = sv.Items[0].ProductId
			*productid = false

		case *color == true:
			sql += "product_color=@color"
			args["percent"] = sv.Items[0].Color
			*color = false

		case *fabric == true:
			sql += "product_fabric=@fabric"
			args["fabric"] = sv.Items[0].Fabric
			*fabric = false

		case *size == true:
			sql += "product_size=@size"
			args["size"] = sv.Items[0].Size
			*size = false

		case *form == true:
			sql += "product_size=@form"
			args["form"] = sv.Items[0].Form
			*form = false

		case *price == true:
			sql += "product_price=@price"
			args["price"] = sv.Items[0].Price
			*price = false

		case *amount == true:
			sql += "product_amount=@amount"
			args["amount"] = sv.Items[0].Amount
			*amount = false

		case *description == true:
			sql += "product_description=@description"
			args["description"] = sv.Items[0].Description
			*description = false

		default:
			nextoption = false
		}

		// comma false
		if !nextoption {
			sql = sql[:len(sql)-1]
			break
		}
		sql += ","
	}
	sql += " WHERE product_detail_id=@id;"
	fmt.Println(sql)

	// Execute sql command
	_, err = conn.Exec(database.CTX, sql, args)
	if err != nil {
		return err
	}

	return nil
}
