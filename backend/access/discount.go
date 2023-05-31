package access

import (
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"commercial-shop.com/database"
	"commercial-shop.com/models"
)

func FindAllDiscount(limit *int, page *int) ([]models.Discount, error) {
	dataSlice := []models.Discount{}
	data := &models.Discount{}

	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := "SELECT * FROM discount ORDER BY discount_id LIMIT @limit OFFSET @offset;"
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
		err := rows.Scan(&data.Id, &data.Description, &data.Percent, &data.DateStart, &data.DateEnd)
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

func FindDiscount(id *string) (models.Discount, error) {
	data := &models.Discount{}

	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return *data, err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := "SELECT * FROM discount WHERE discount_id='" + *id + "';"

	// Get rows from conn with SQL command
	err = conn.QueryRow(database.CTX, sql).Scan(&data.Id, &data.Description, &data.Percent, &data.DateStart, &data.DateEnd)
	if err != nil {
		return *data, err
	}

	return *data, nil
}

func CreateDiscount(data *models.Discount) error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := ""
	args := pgx.NamedArgs{}
	if data.DateStart.IsZero() && data.DateEnd.IsZero() {
		sql = "INSERT INTO discount (discount_id, discount_description, discount_percent) VALUES (@id, @description, @percent);"
		args = pgx.NamedArgs{
			"id":          data.Id,
			"description": data.Description,
			"percent":     data.Percent,
		}
	} else {
		sql = "INSERT INTO discount VALUES (@id, @description, @percent, @dateStart, @dateEnd);"
		args = pgx.NamedArgs{
			"id":          data.Id,
			"description": data.Description,
			"percent":     data.Percent,
			"dateStart":   data.DateStart,
			"dateEnd":     data.DateEnd,
		}
	}

	// Execute sql command
	_, err = conn.Exec(database.CTX, sql, args)
	if err != nil {
		return err
	}

	return nil
}

func UpdateDiscount(data *models.Discount) error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := ""
	args := pgx.NamedArgs{}
	if data.DateStart.IsZero() && data.DateEnd.IsZero() {
		sql = "UPDATE discount SET discount_description=@description, discount_percent=@percent WHERE discount_id=@id;"
		args = pgx.NamedArgs{
			"id":          data.Id,
			"description": data.Description,
			"percent":     data.Percent,
		}
	} else {
		sql = "UPDATE discount SET discount_description=@description, discount_percent=@percent, discount_date_start=@dateStart, discount_date_end=@dateEnd WHERE discount_id=@id;"
		args = pgx.NamedArgs{
			"id":          data.Id,
			"description": data.Description,
			"percent":     data.Percent,
			"dateStart":   data.DateStart,
			"dateEnd":     data.DateEnd,
		}
	}

	// Execute sql command
	_, err = conn.Exec(database.CTX, sql, args)
	if err != nil {
		return err
	}

	return nil
}

func DeleteDiscount(id *string) error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := "DELETE FROM discount WHERE discount_id='" + *id + "';"

	// Execute sql command
	_, err = conn.Exec(database.CTX, sql)
	if err != nil {
		return err
	}

	return nil
}
