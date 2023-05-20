package category 

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/jackc/pgx/v5/pgxpool"

    "commercial-shop.com/database"
)

type Category struct {
    Id   string
    Name string
}

func FindAll() ([]Category, error) {
    categorySlice := []Category{}
    c := &Category{}
    ctx := context.Background()

    // Get connection to database and check error
    // Close after function end with defer
    conn, err := pgxpool.New(ctx, database.CONNECT_STR)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to connect to database: %v", err)
        return nil, err 
    }
    defer conn.Close()

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

func FindById(id string) (Category, error) {
    c := &Category{}
    ctx := context.Background()

    // Get connection to database and check error
    // Close after function end with defer
    conn, err := pgxpool.New(ctx, database.CONNECT_STR)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to connect to database: %v", err)
        return *c, err 
    }
    defer conn.Close()

    // sql as a basic SQL commamd
    sql := "SELECT * FROM category where category_id='" + id + "';"

    // Get rows from conn with SQL command
    err = conn.QueryRow(ctx, sql).Scan(&c.Id, &c.Name)
    if err != nil {
        log.Fatal(err)
        return *c, err
    }
  
    return *c, nil
}
