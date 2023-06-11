# Commercial-shop

Dùng cho mục đích thương mại điện tử

Hiện tại chưa có mục lục. Sẽ bổ sung sau

## Tại sao?

**Quan trọng** Vì không có giải pháp hiện tại không hề đơn giản và chuyên
nghiệp. Tự phát minh cho bản thân

### Mục tiêu

**Quan trọng:** 1 web server đầy đủ tính năng cơ bản cho phục vụ thương mại
điện tử trong việc bán quần áo.

**Lưu ý:** Khách hàng sử dụng phần mềm không phải là 1 người mù tịt hoàn
toàn về công nghệ thông tin cũng như quá cổ hủ chỉ muốn vấn đề giải quyết đơn
giản gọn lẹ. Nếu muốn đơn giản dễ sử dụng và quản lý.

**Đề xuất mục tiêu khác:**
+ Bảo mật web server
+ Dễ setup và sử dụng?
+ Bảo trì?
+ Các thương mại điện tử khác ngoài bán quần áo?

1 số thứ sẽ không đúng hoàn toàn trong mô tả này. Hãy đợi để chúng được hoàn
chỉnh theo thời gian.

## Về công nghệ?

### Database

+ Sẽ sử dụng Postgresql
  + Vì nó là SQL
  + Hỗ trợ nhiề u kiểu dữ liệu.

#### Bán quần áo

##### Mô tả

- Website bán quần áo với người sử dụng là khách hàng truy cập website và người
quản lý cửa hàng sẽ quản lý thông tin cửa hàng trên chính website đó. Với sản
phẩm là quần áo gồm có nhiều thuộc tính như size, màu, chất liệu,vv,. Website
sẽ có 1 số chức năng như sau :
  - Khách hàng khi mua hàng có thể không cần đăng nhập vào website và có thể 
  thêm sản phẩm vào giỏ hàng và thanh toán trực tiếp sau khi đã điền đầy đủ
  các thông tin cần thiết ( Họ tên , SĐT , gmail , địa chỉ giao hàng , ..)
  - Với khách hàng đã đăng nhập và mua hàng thì họ sẽ xem được các thông tin
  của đơn hàng như là trạng thái : chờ xác nhận, đang giao hàng, đã giao thành
  công , vv .
  - Khách hàng có thể lọc sản phẩm theo giá cao thấp , tìm kiếm sản phẩm dựa
  theo mục lục hoặc tên , vv
  - Khách hàng có thể chỉnh sửa giỏ hàng như là số lượng sản phẩm mua, chọn
  sản phẩm nào được mua , xóa sản phẩm khỏi giỏ , cập nhật thông tin sản phẩm
  trong giỏ hàng
  - Khách hàng có thể tự chỉnh sửa thông tin tài khoản của chính mình (các
  thông tin cá nhân được hiển thị trên website )
  - Người quản lý cửa hàng có thể xem, thêm, sửa, xóa (CRUD) dữ liệu trên
  website như quần áo , thông tin bill (trong trường hợp nhầm lẫn hoặc sai
  thông tin), như là account khách hàng , tạo tài khoản admin quản lý website
  khác , quản lý các sự kiện giảm giá , phần trăm giảm giá

##### Mô tả cơ sở dữ liệu ####

- Table ***Category*** 

    |            | category_id | category_name |
    | :--------- |:------------|:------------- |
    |  Constraint| Primary key | Default       |
    |  Data type | VARCHAR(20) | VARCHAR(100)  |

- Table ***Product***

    |            | product_id | category_id             | product_name |
    | :--------- |:------------|:-----------            |:------------ |
    |  Constraint| Primary key | FOREIGN KEY , NOT NULL | DEFAULT      |
    |  Data type | VARCHAR(20) | VARCHAR(20)            | VARCHAR(100) |  

- Table ***ProductDetail***

    |            | product_detail_id | product_id   | product_color   | product_fabric|  product_size | product_price   | product_amount | product_description   |
    | :--------- |:------------      |:-----------  |:------------    | :------------ | :------------ | :------------   | :------------  |:------------          | 
    |  Constraint| Primary key       | FOREIGN KEY , NOT NULL| DEFAULT| DEFAULT       | DEFAULT       | DEFAULT         |  DEFAULT       |  DEFAULT              | 
    |  Data type | VARCHAR(20)       | VARCHAR(20)  | VARCHAR(20)     | VARCHAR(20)   | VARCHAR(4)    | INT             | INT            | VARCHAR(400)          |

- Table ***ProductImage***

    |            |   product_image_id |  product_detail_id     | product_image |
    | :--------- |:------------       |:-----------            |:------------  |
    |  Constraint| Primary key        | FOREIGN KEY , NOT NULL |               |
    |  Data type | VARCHAR(20)        | VARCHAR(20)            | BYTEA         | 

- Table ***Account***

    |            | account_username   |  account_password   | account_displayname |    role_id    | 
    | :--------- |:------------       |:------------        | :------------       | :------------ |
    |  Constraint| Primary key        |                     |                     | CHECK         | 
    |  Data type | VARCHAR(20)        | VARCHAR(20)         | VARCHAR(20)         |INT            |

- Table ***Customer*** 

    |            | customer_id   |  account_username  |  customer_name     | customer_phone      | customer_email | customer_address   | 
    | :--------- |:------------  | :-----------       |:-----------        |:------------        | :------------  | :------------      |
    |  Constraint| Primary key   |  FOREIGN KEY       | DEFAULT            | DEFAULT             | DEFAULT        |  DEFAULT           |
    |  Data type | VARCHAR(20)   |  VARCHAR(20)       | VARCHAR(50)        | VARCHAR(20)         | VARCHAR(100)   |  VARCHAR(100)      |


- Table ***Discount***

    |            | discount_id       |   discount_description  |  discount_percent   | discount_date_start   | discount_date_end   | 
    | :--------- |:------------      |:-----------             |:------------        | :------------         | :----------         |
    |  Constraint| Primary key       | DEFAULT                 | DEFAULT             | DEFAULT               | DEFAULT             |
    |  Data type | VARCHAR(20)       | VARCHAR(200)            | REAL                | DATE                  | DATE                |

- Table ***BillInfo***

    |            |  bill_id          |    customer_id          |   bill_date             |  bill_status           | bill_payment | 
    | :--------- |:------------      |:-----------             |:------------            | :------------          | :----------  |
    |  Constraint| Primary key       | FOREIGN KEY , NOT NULL  | DEFAULT                 | CHECK                  | CHECK        |
    |  Data type | VARCHAR(20)       | VARCHAR(20)             | VARCHAR(20)             | INT                    | INT          |

- Table ***BillDetail***

    |            |  bill_detail_id   | bill_id                 |   product_detail_id     | discount_id            | bill_amount  | 
    | :--------- |:------------      |:-----------             |:------------            | :------------          | :----------  |
    |  Constraint| Primary key       | FOREIGN KEY , NOT NULL  |  FOREIGN KEY , NOT NULL |  FOREIGN KEY , NOT NULL| CHECK        |
    |  Data type | VARCHAR(20)       | VARCHAR(20)             |  VARCHAR(20)            |  VARCHAR(20)           | INT          |

##### Use case diagram


##### Mã nguồn

Xem tại [đây](database/)

### Web API

Sử dụng Go Web Framework tên là [Gin](https://github.com/gin-gonic/gin)

Sao lại là Go?
+ Go
  + Nhanh (compiled language, Want to work like scripting? It can too.)
  + Memory safe (Garbage collector)
  + Hõ trợ kiểu dữ liệu 
  + Đơn giản?
    + Các syntax đa phần là sạch vì ít khi phải nhớ
    + Không class nhưng có struct và methods
    + Có embedded struct để thừa kế các struct khác
    + Không Override? Bạn chỉ cần copy paste và edit code lại. Get the job done.
  + Package Management?
    + Không hoàn hảo nhưng phù hợp (Because Go packages use links which is bad
    idea)
  + Weird but fun?

Mã nguồn xem tại [đây](backend/)

Dự tính phát triển thêm để code có thể tự tạo code backend tùy thuộc vào bảng.
Hãy bàn bạc sau khi đã đạt được các mục tiêu chính.

### Client

#### Web

Chưa có quyết định rõ ràng hoàn toàn

Phác thảo 1 số công nghệ quyết định sử dụng
+ TypeScript
+ React + React-router
+ TailwindCSS

#### Mobile

Là 1 tính năng tốt để có tuy nhiên vẫn chưa có quyết định rõ ràng cũng chưa
chăc chắn làm thành công dự án hay không nhưng hãy bàn bạc sau khi chúng ta
đạt các cột mốc lớn ở phía trên.

### Server

1 phần cực kì quan trọng công nghệ sử dụng sẽ là [NixOS](https://nixos.org)

Tại sao? Để đánh bật những hoài nghi về giải pháp khác của bạn 

+ Docker? Nix package manager có thể làm nhiều hơn Docker
+ Operating System? NixOS xây trên Nix package manager.
+ Ansible? Bạn có NixOS có thể 

Đúng vậy chỉ cần học Nix bạn có 3 thứ gộp lại với nhau

Phần dễ nhất: Học Nix

Phần khó nhất: Học Nix

Phần vui nhất: Học Nix.

Hãy thừa nhận nếu bạn muốn chạm vô nó, bạn phải học nó.

Hiện tại vẫn chưa có những giải pháp áp dụng vô thực tế để thuyết phục bạn.
Cũng như các khuyết điểm của công nghệ. Tất cả những miêu tả này về server này
còn khá mơ hồ và cần được thử nghiệm trước khi đưa ra dẫn chứng thuyết phục bạn.

**Stay tune for more?**

## Tổng kết

Các công nghệ trên có thể không hoàn hảo tuy nhiên phù hợp với vấn đề cần phải
giải quyết.

## Đóng góp 

goatastronaut0212 (Nguyễn Nguyên Khang)

Khoadz299 (Nguyễn Hiền Anh Khoa)
