package access

import (
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"commercial-shop.com/database"
	"commercial-shop.com/models"
)

func FindAllBillInfo(limit *int, page *int) ([]models.BillInfo, error) {
	dataSlice := []models.BillInfo{}
	data := &models.BillInfo{}

	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := "SELECT * FROM bill_info ORDER BY bill_id LIMIT @limit OFFSET @offset;"
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
		err := rows.Scan(&data.Id, &data.CustomerId, &data.Date)
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

func FindBillInfo(id *string) (models.BillInfo, error) {
	data := &models.BillInfo{}

	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return *data, err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := "SELECT * FROM bill_info WHERE bill_id='" + *id + "';"

	// Get rows from conn with SQL command
	err = conn.QueryRow(database.CTX, sql).Scan(&data.Id, &data.CustomerId, &data.Date, &data.Status, &data.Payment)
	if err != nil {
		return *data, err
	}

	return *data, nil
}

func CreateBillInfo(data *models.BillInfo) error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := "INSERT INTO BILL_INFO VALUES (@id , @customer_id , @billdate, @status, @payment);"
	args := pgx.NamedArgs{
		"id":          data.Id,
		"customer_id": data.CustomerId,
		"date":        data.Date,
		"status":      data.Status,
		"payment":     data.Payment,
	}

	// Execute sql command
	_, err = conn.Exec(database.CTX, sql, args)
	if err != nil {
		return err
	}

	return nil
}

func UpdateBillInfo(data *models.BillInfo) error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := "UPDATE BILL_INFO SET customer_id=@customer_id, bill_date=@billdate, bill_status=@status, bill_payment=@payment WHERE bill_id=@id;"
	args := pgx.NamedArgs{
		"id":          data.Id,
		"customer_id": data.CustomerId,
		"date":        data.Date,
		"status":      data.Status,
		"payment":     data.Payment,
	}
	// Execute sql command
	_, err = conn.Exec(database.CTX, sql, args)
	if err != nil {
		return err
	}

	return nil
}

func DeleteBillInfo(id *string) error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := "DELETE FROM BILL_INFO WHERE bill_id='" + *id + "';"

	// Execute sql command
	_, err = conn.Exec(database.CTX, sql)
	if err != nil {
		return err
	}

	return nil
}
