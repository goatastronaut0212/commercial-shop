package access

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"commercial-shop.com/database"
	"commercial-shop.com/models"
)

func FindAllBillDetail() ([]models.Bill_Detail , error) {
	dataSlice := []models.Bill_Detail{}
	data := &models.Bill_Detail{}

	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := "SELECT * FROM BILL_DETAIL;"

	// Get rows from conn with SQL command
	rows, err := conn.Query(database.CTX, sql)
	if err != nil {
		return nil, err
	}

	// convert each rows to struct and append to Slice to return
	for rows.Next() {
		err := rows.Scan( &data.ID , &data.BillID , &data.ProductID , &data.DiscountID )
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

func FindBillDetail(id *string) ([]models.Bill_Detail , error) {
	dataSlice := []models.Bill_Detail{}
	data := &models.Bill_Detail{}

	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := "SELECT * FROM BILL_DETAIL WHERE bill_detail_id='" + *id + "';"

	// Get rows from conn with SQL command
	rows, err := conn.Query(database.CTX, sql)
	if err != nil {
		return nil, err
	}

	// convert each rows to struct and append to Slice to return
	for rows.Next() {
		err := rows.Scan( &data.ID , &data.BillID , &data.ProductID , &data.DiscountID )
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

func CreateBillDetail(data *models.Bill_Detail) error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := "INSERT INTO BILL_DETAIL VALUES (@id , @bill_id , @product_id , @discount_id );"
	args := pgx.NamedArgs{
		"id":               data.ID,
		"bill_id":          data.BillID,
		"product_id":	    data.ProductID,
		"discount_id": 		data.DiscountID,
	}

	// Execute sql command
	_, err = conn.Exec(database.CTX, sql, args)
	if err != nil {
		return err
	}

	return nil
}

func UpdateBillDetail(data *models.Bill_Detail) error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := "UPDATE BILL_DETAIL SET bill_id=@bill_id , product_id=@product_id , discount_id=@discount_id WHERE bill_detail_id=@id;"
	args := pgx.NamedArgs{
		"id":               data.ID,
		"bill_id":          data.BillID,
		"product_id":	    data.ProductID,
		"discount_id": 		data.DiscountID,
	}
	// Execute sql command
	_, err = conn.Exec(database.CTX, sql, args)
	if err != nil {
		return err
	}

	return nil
}

func DeleteBillDetail(id *string) error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := "DELETE FROM BILL_DETAIL WHERE bill_detail_id='" + *id + "';"

	// Execute sql command
	_, err = conn.Exec(database.CTX, sql)
	if err != nil {
		return err
	}

	return nil
}
