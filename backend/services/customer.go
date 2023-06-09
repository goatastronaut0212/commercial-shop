package services

import (
	"strconv"

	"commercial-shop.com/database"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Customer struct {
	Id      string `json:"id"         binding:"required"`
	Name    string `json:"name"       binding:"required"`
	Phone   string `json:"phone"      binding:"required"`
	Email   string `json:"email"      binding:"required"`
	Address string `json:"address"    binding:"required"`
}

type CustomerService struct {
	Items []Customer
}

func (sv *CustomerService) Get() error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := "SELECT * FROM Customer WHERE customer_id='" + sv.Items[0].Id + "';"

	// Get rows from conn with SQL command
	err = conn.QueryRow(database.CTX, sql).Scan(&sv.Items[0].Id, &sv.Items[0].Name, &sv.Items[0].Phone, &sv.Items[0].Email, &sv.Items[0].Address)
	if err != nil {
		return err
	}

	return nil
}

func (sv *CustomerService) GetAll(limit *int, page *int) error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := "SELECT * FROM Customer ORDER BY customer_id LIMIT @limit OFFSET @offset;"
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
		sv.Items = append(sv.Items, Customer{})
		err := rows.Scan(&sv.Items[i].Id, &sv.Items[i].Name, &sv.Items[i].Phone, &sv.Items[i].Email, &sv.Items[i].Address)
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

func (sv *CustomerService) Create() error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Execute sql command
	sql := "INSERT INTO Customer VALUES (@id, @name, @phone, @email , @address);"
	args := pgx.NamedArgs{
		"id":      sv.Items[0].Id,
		"name":    sv.Items[0].Name,
		"phone":   sv.Items[0].Phone,
		"email":   sv.Items[0].Email,
		"address": sv.Items[0].Address,
	}

	// Execute sql command
	_, err = conn.Exec(database.CTX, sql, args)
	if err != nil {
		return err
	}

	return nil
}

func (sv *CustomerService) Update() error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := "UPDATE Customer SET customer_name=@name, customer_phone=@phone, customer_email=@email, customer_address=@address WHERE customer_id=@id;"
	args := pgx.NamedArgs{
		"id":      sv.Items[0].Id,
		"name":    sv.Items[0].Name,
		"phone":   sv.Items[0].Phone,
		"email":   sv.Items[0].Email,
		"address": sv.Items[0].Address,
	}
	// Execute sql command
	_, err = conn.Exec(database.CTX, sql, args)
	if err != nil {
		return err
	}

	return nil
}

func (sv *CustomerService) Delete(id *string) error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := "DELETE FROM Customer WHERE customer_id='" + *id + "';"

	// Execute sql command
	_, err = conn.Exec(database.CTX, sql)
	if err != nil {
		return err
	}

	return nil
}
