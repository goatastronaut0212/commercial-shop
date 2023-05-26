# Commercial website backend (bản dịch tiếng việt)

Được viết bằng Go.

You need English version. Click [here](README.md)

## Coding style

Đọc tại [đây](https://google.github.io/styleguide/go/index)

## Cách để sử dụng backend?

### Lấy các phụ thuộc packages

```sh
go mod tidy
```

### Xây dựng và biên dịch chương trình

```sh
go build .
```

### Chạy chương trình

**Lưu ý**: tùy thuộc vào hệ điều hành sẽ tạo ra các files binary khác nhau.

+ Nếu là các hệ điều hành Unix-like

```sh
./commercial-shop.com
```

+ Nếu là Windows

```sh
./commercial-shop.com.exe
```

### Dọn dẹp các binary files

```sh
go clean
```