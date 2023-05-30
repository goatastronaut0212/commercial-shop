package access

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"commercial-shop.com/database"
	"commercial-shop.com/models"
)

func FindAllAccount() ([]models.Account, error) {
	dataSlice := []models.Account{}
	data := &models.Account{}

	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := "SELECT * FROM ACCOUNT;"

	// Get rows from conn with SQL command
	rows, err := conn.Query(database.CTX, sql)
	if err != nil {
		return nil, err
	}

	// convert each rows to struct and append to Slice to return
	for rows.Next() {
		err := rows.Scan(&data.Username, &data.CustomerId, &data.Password, &data.DisplayName, &data.RoleID)
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

func FindAccount(username *string) (models.Account, error) {
	data := &models.Account{}

	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return *data, err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := "SELECT * FROM ACCOUNT WHERE account_username='" + *username + "';"

	// Get rows from conn with SQL command
	err = conn.QueryRow(database.CTX, sql).Scan(&data.Username, &data.CustomerId, &data.Password, &data.DisplayName, &data.RoleID)
	if err != nil {
		return *data, err
	}

	return *data, nil
}

func CreateAccount(data *models.Account) error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := "INSERT INTO ACCOUNT VALUES (@username, @customer_id, @password, @display_name, @role_id);"
	args := pgx.NamedArgs{
		"username":     data.Username,
		"customer_id":  data.CustomerId,
		"password":     data.Password,
		"display_name": data.DisplayName,
		"role_id":      data.RoleID,
	}

	// Execute sql command
	_, err = conn.Exec(database.CTX, sql, args)
	if err != nil {
		return err
	}

	return nil
}

func UpdateAccount(data *models.Account) error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := "UPDATE ACCOUNT SET customer_id=@customer_id , account_password=@password , account_displayname=@displayname ,  role_id=@roleID WHERE account_username=@username;"
	args := pgx.NamedArgs{
		"username":     data.Username,
		"customer_id":  data.CustomerId,
		"password":     data.Password,
		"display_name": data.DisplayName,
		"role_id":      data.RoleID,
	}
	// Execute sql command
	_, err = conn.Exec(database.CTX, sql, args)
	if err != nil {
		return err
	}

	return nil
}

func DeleteAccount(username *string) error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := "DELETE FROM ACCOUNT WHERE account_username='" + *username + "';"

	// Execute sql command
	_, err = conn.Exec(database.CTX, sql)
	if err != nil {
		return err
	}

	return nil
}
