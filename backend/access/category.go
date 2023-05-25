package access 

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/jackc/pgx/v5"

    "commercial-shop.com/database"
    "commercial-shop.com/models"
)

var ctx = context.Background()

func FindCategoryAll() ([]models.Category, error) {
    categorySlice := []models.Category{}
    c := &models.Category{}

    // Connect to database and close after executing command
    conn, err := pgx.Connect(ctx, database.CONNECT_STR)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to connect to database: %v", err)
        return nil, err 
    }
    defer conn.Close(ctx)

    // sql as a basic SQL commamd
    sql := "SELECT * FROM category;"

    // Get rows from conn with SQL command
    rows, err := conn.Query(ctx, sql)
    if err != nil {
        log.Fatal(err)
        return nil, err
    }
  
    // convert each rows to struct and append to Slice to return
    for rows.Next() {
        err := rows.Scan(&c.Id, &c.Name)
        if err != nil {
            log.Fatal(err)
            return nil, err
        }
        categorySlice = append(categorySlice, *c)
    }
    if err := rows.Err(); err != nil {
        log.Fatal(err)
        return nil, err 
    }

    return categorySlice, nil
}

func FindCategoryById(id *string) (models.Category, error) {
    cg := &models.Category{}

    // Connect to database and close after executing command
    conn, err := pgx.Connect(ctx, database.CONNECT_STR)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to connect to database: %v", err)
        return *cg, err 
    }
    defer conn.Close(ctx)

    // sql as a basic SQL commamd
    sql := "SELECT * FROM category WHERE category_id='" + *id + "';"

    // Get rows from conn with SQL command
    err = conn.QueryRow(ctx, sql).Scan(&cg.Id, &cg.Name)
    if err != nil {
        log.Fatal(err)
        return *cg, err
    }
  
    return *cg, nil
}

func CreateCategory(cg *models.Category) error {
    // Connect to database and close after executing command
    conn, err := pgx.Connect(ctx, database.CONNECT_STR)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to connect to database: %v", err)
        return err 
    }
    defer conn.Close(ctx)

    // sql as a basic SQL commamd
    sql := "INSERT INTO category VALUES ('" + cg.Id + "', '"+ cg.Name +"');"

    // Get rows from conn with SQL command
    _, err = conn.Exec(ctx, sql)
    if err != nil {
        log.Fatal(err)
        return err
    }
  
    return nil
}

func UpdateCategory(cg *models.Category) error {
    // Connect to database and close after executing command
    conn, err := pgx.Connect(ctx, database.CONNECT_STR)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to connect to database: %v", err)
        return err 
    }
    defer conn.Close(ctx)

    // sql as a basic SQL commamd
    sql := "UPDATE category SET category_name='"+ cg.Name +"' WHERE category_id='"+ cg.Id +"';"

    // Get rows from conn with SQL command
    _, err = conn.Exec(ctx, sql)
    if err != nil {
        log.Fatal(err)
        return err
    }
  
    return nil
}

func DeleteCategory(id *string) error {
    // Connect to database and close after executing command
    conn, err := pgx.Connect(ctx, database.CONNECT_STR)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to connect to database: %v", err)
        return err 
    }
    defer conn.Close(ctx)

    // sql as a basic SQL commamd
    sql := "DELETE FROM category WHERE category_id='" + *id + "';"

    // Get rows from conn with SQL command
    _, err = conn.Exec(ctx, sql)
    if err != nil {
        log.Fatal(err)
        return err
    }
  
    return nil
}
