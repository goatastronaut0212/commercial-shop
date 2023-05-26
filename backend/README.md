# Commercial website backend

Written in Go.

Bạn cần bản tiếng Việt. Vào [đây](README-vn.md)

## Coding style

See in [here](https://google.github.io/styleguide/go/index)

## How to use backend?

### Grab all package dependencies

```
go mod tidy
```

### Build program

```
go build .
```

### Run program

**NOTE**: Depend on your operating system it will generate different binary files.

+ If Unix-like

```
./commercial-shop.com
```

+ If Windows

```
./commercial-shop.com.exe
```

### Clean binary file

```
go clean
```

## Testing

Since we only using `testing` standard library.

If you want to test and see if it pass

```
go test
```

If you want to check all functions lists all of the tests and their results. 

```
go test -v
```