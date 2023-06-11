package services

import (
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"commercial-shop.com/database"
	"commercial-shop.com/models"
)

type AccountService struct {
	Items []models.Account
}

func (sv *AccountService) Get() error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// SQL commamd
	sql := "SELECT * FROM Account WHERE account_username='" + sv.Items[0].Username + "';"

	// Get rows from conn with SQL command
	err = conn.QueryRow(database.CTX, sql).Scan(
		&sv.Items[0].Username,
		&sv.Items[0].CustomerId,
		&sv.Items[0].Password,
		&sv.Items[0].DisplayName,
		&sv.Items[0].RoleId,
	)
	if err != nil {
		return err
	}

	return nil
}

func (sv *AccountService) GetAll(limit *int, page *int) error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// SQL commamd
	sql := "SELECT * FROM Account LIMIT @limit OFFSET @offset;"
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
		sv.Items = append(sv.Items, models.Account{})

		err := rows.Scan(
			&sv.Items[i].Username,
			&sv.Items[i].CustomerId,
			&sv.Items[i].Password,
			&sv.Items[i].DisplayName,
			&sv.Items[i].RoleId,
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

func (sv *AccountService) Create() error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// SQL commamd
	sql := "INSERT INTO Account VALUES (@username, @customer_id, @password, @display_name, @role_id);"
	args := pgx.NamedArgs{
		"username":     sv.Items[0].Username,
		"customer_id":  sv.Items[0].CustomerId,
		"password":     sv.Items[0].Password,
		"display_name": sv.Items[0].DisplayName,
		"role_id":      sv.Items[0].RoleId,
	}

	// Execute sql command
	_, err = conn.Exec(database.CTX, sql, args)
	if err != nil {
		return err
	}

	return nil
}

func (sv *AccountService) Update() error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// SQL commamd
	sql := "UPDATE Account SET customer_id=@customer_id, account_password=@password, account_displayname=@display_name, role_id=@role_id WHERE account_username=@username;"
	args := pgx.NamedArgs{
		"username":     sv.Items[0].Username,
		"customer_id":  sv.Items[0].CustomerId,
		"password":     sv.Items[0].Password,
		"display_name": sv.Items[0].DisplayName,
		"role_id":      sv.Items[0].RoleId,
	}

	// Execute sql command
	_, err = conn.Exec(database.CTX, sql, args)
	if err != nil {
		return err
	}

	return nil
}

func (sv *AccountService) Delete() error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// SQL commamd
	sql := "DELETE FROM Account WHERE account_username='" + sv.Items[0].Username + "';"

	// Execute sql command
	_, err = conn.Exec(database.CTX, sql)
	if err != nil {
		return err
	}

	return nil
}
