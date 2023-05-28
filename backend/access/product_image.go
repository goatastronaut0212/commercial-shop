package access

import (
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"commercial-shop.com/database"
	"commercial-shop.com/models"
)

func FindAllProductImage(limit *int, page *int) ([]models.ProductImage, error) {
	dataSlice := []models.ProductImage{}
	data := &models.ProductImage{}

	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := "SELECT product_image_id, product_detail_id, product_image FROM product_image ORDER BY product_image_id LIMIT @limit OFFSET @offset;"
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
		err := rows.Scan(&data.Id, &data.IdProductDetail, &data.Image)
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

func FindProductImage(id *string) (models.ProductImage, error) {
	data := &models.ProductImage{}

	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return *data, err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := "SELECT * FROM product_image WHERE product_image_id='" + *id + "';"

	// Get rows from conn with SQL command
	err = conn.QueryRow(database.CTX, sql).Scan(&data.Id, &data.IdProductDetail, &data.Image)
	if err != nil {
		return *data, err
	}

	return *data, nil
}

func CreateProductImage(data *models.ProductImage) error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := ""
	args := pgx.NamedArgs{}
	if data.Image == nil {
		sql = "INSERT INTO product_image (product_image_id, product_detail_id) VALUES (@id, @idProductDetail);"
		args = pgx.NamedArgs{
			"id":              data.Id,
			"idProductDetail": data.IdProductDetail,
		}
	} else {
		sql = "INSERT INTO product_image VALUES (@id, @idProductDetail, @image);"
		args = pgx.NamedArgs{
			"id":              data.Id,
			"idProductDetail": data.IdProductDetail,
			"image":           data.Image,
		}
	}

	// Execute sql command
	_, err = conn.Exec(database.CTX, sql, args)
	if err != nil {
		return err
	}

	return nil
}

func UpdateProductImage(data *models.ProductImage) error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := ""
	args := pgx.NamedArgs{}
	if data.Image == nil {
		sql = "UPDATE product_image SET product_detail_id=@idProductDetail, product_image=@image WHERE product_image_id=@id;"
		args = pgx.NamedArgs{
			"id":              data.Id,
			"idProductDetail": data.IdProductDetail,
		}
	} else {
		sql = "UPDATE product_image SET product_detail_id=@idProductDetail, product_image=@image WHERE product_image_id=@id;"
		args = pgx.NamedArgs{
			"id":              data.Id,
			"idProductDetail": data.IdProductDetail,
			"image":           data.Image,
		}
	}

	// Execute sql command
	_, err = conn.Exec(database.CTX, sql, args)
	if err != nil {
		return err
	}

	return nil
}

func DeleteProductImage(id *string) error {
	// Connect to database and close after executing command
	conn, err := pgxpool.New(database.CTX, database.CONNECT_STR)
	if err != nil {
		return err
	}
	defer conn.Close()

	// sql as a basic SQL commamd
	sql := "DELETE FROM product_image WHERE product_image_id='" + *id + "';"

	// Execute sql command
	_, err = conn.Exec(database.CTX, sql)
	if err != nil {
		return err
	}

	return nil
}
