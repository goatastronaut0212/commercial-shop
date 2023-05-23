package category 

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/jackc/pgx/v5/pgxpool"

    "commercial-shop.com/database"
)

var ctx = context.Background()

type Category struct {
    Id   string   `json:"id" binding:"required"`
    Name string   `json:"name" binding:"required"`
}

func FindAll() ([]Category, error) {
    categorySlice := []Category{}
    c := &Category{}

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

    // Get connection to database and check error
    // Close after function end with defer
    conn, err := pgxpool.New(ctx, database.CONNECT_STR)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to connect to database: %v", err)
        return *c, err 
    }
    defer conn.Close()

    // sql as a basic SQL commamd
    sql := "SELECT * FROM category WHERE category_id='" + id + "';"

    // Get rows from conn with SQL command
    err = conn.QueryRow(ctx, sql).Scan(&c.Id, &c.Name)
    if err != nil {
        log.Fatal(err)
        return *c, err
    }
  
    return *c, nil
}

func Create(id string, name string) error {
    // Get connection to database and check error
    // Close after function end with defer
    conn, err := pgxpool.New(ctx, database.CONNECT_STR)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to connect to database: %v", err)
        return err 
    }
    defer conn.Close()

    // sql as a basic SQL commamd
    sql := "INSERT INTO category VALUES ('" + id + "', '"+ name +"');"

    // Get rows from conn with SQL command
    _, err = conn.Exec(ctx, sql)
    if err != nil {
        log.Fatal(err)
        return err
    }
  
    return nil
}

func Update(id string, name string) error {
    // Get connection to database and check error
    // Close after function end with defer
    conn, err := pgxpool.New(ctx, database.CONNECT_STR)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to connect to database: %v", err)
        return err 
    }
    defer conn.Close()

    // sql as a basic SQL commamd
    sql := "UPDATE category SET category_name='"+ name +"' WHERE category_id='"+ id +"';"

    // Get rows from conn with SQL command
    _, err = conn.Exec(ctx, sql)
    if err != nil {
        log.Fatal(err)
        return err
    }
  
    return nil
}

func Delete(id string) error {
    // Get cpgxpoolonnection to database and check error
    // Close after function end with defer
    conn, err := pgxpool.New(ctx, database.CONNECT_STR)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to connect to database: %v", err)
        return err 
    }
    defer conn.Close()

    // sql as a basic SQL commamd
    sql := "DELETE FROM category WHERE category_id='" + id + "';"

    // Get rows from conn with SQL command
    _, err = conn.Exec(ctx, sql)
    if err != nil {
        log.Fatal(err)
        return err
    }
  
    return nil
}
