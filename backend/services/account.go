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
		&sv.Items[0].RoleId,
		&sv.Items[0].Password,
		&sv.Items[0].DisplayName,
		&sv.Items[0].Email,
		&sv.Items[0].Active,
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
			&sv.Items[i].RoleId,
			&sv.Items[i].Password,
			&sv.Items[i].DisplayName,
			&sv.Items[i].Email,
			&sv.Items[i].Active,
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

	// SQL commamd check options input
	sql := "INSERT INTO Account (account_username, account_password, role_id, account_email) VALUES (@username, @password, @role_id, @email);"
	args := pgx.NamedArgs{
		"username": sv.Items[0].Username,
		"role_id":  sv.Items[0].RoleId,
		"password": sv.Items[0].Password,
		"email":    sv.Items[0].Email,
	}

	if sv.Items[0].DisplayName != "" {
		sql = "INSERT INTO Account (account_username, account_password, account_displayname, role_id, account_email) VALUES (@username, @password, @display_name, @role_id, @email);"
		args["display_name"] = sv.Items[0].DisplayName
	}

	// Execute sql command
	_, err = conn.Exec(database.CTX, sql, args)
	if err != nil {
		return err
	}

	return nil
}

func (sv *AccountService) Update(password *bool, displayname *bool, roleid *bool, email *bool) error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// SQL commamd check options input
	sql := "UPDATE Account SET "
	args := pgx.NamedArgs{
		"username": sv.Items[0].Username,
	}
	nextoption := true

	for i := 0; i < 3; i++ {
		switch {
		case *password == true:
			sql += "account_password=@password"
			args["password"] = sv.Items[0].Password
			*password = false

		case *displayname == true:
			sql += "account_displayname=@display_name"
			args["display_name"] = sv.Items[0].DisplayName
			*displayname = false

		case *roleid == true:
			sql += "role_id=@role_id"
			args["role_id"] = sv.Items[0].RoleId
			*roleid = false

		case *email == true:
			sql += "account_email=@email"
			args["email"] = sv.Items[0].Email
			*email = false

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
	sql += " WHERE account_username=@username;"

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
