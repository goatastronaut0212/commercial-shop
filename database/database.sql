CREATE DATABASE shopdb;

\c shopdb;

CREATE TABLE Category (
    category_id    VARCHAR(20) NOT NULL,
    category_name  VARCHAR(100) DEFAULT 'Chưa có thông tin',
    PRIMARY KEY (category_id)
);

CREATE TABLE Product (
    product_id      VARCHAR(20)  NOT NULL,
    category_id     VARCHAR(20)  NOT NULL,
    product_name    VARCHAR(100)  DEFAULT 'Chưa có thông tin',
    PRIMARY KEY (product_id),
    CONSTRAINT fk_category_id_for_product
        FOREIGN KEY (category_id)
        REFERENCES Category(category_id)
);

CREATE TABLE ProductDetail (
    product_detail_id    VARCHAR(20)  NOT NUll,
    product_id           VARCHAR(20)  NOT NULL,
    product_color        VARCHAR(20)  DEFAULT 'Chưa có thông tin',
    product_fabric       VARCHAR(20)  DEFAULT 'Chưa có thông tin',
    product_size         VARCHAR(4)   DEFAULT 'None',
    product_price        INT          DEFAULT 0,
    product_amount       INT          DEFAULT 0 CHECK(product_amount >= 0),
    product_description  VARCHAR(400) DEFAULT 'Chưa có thông tin',
    PRIMARY KEY (product_detail_id),
    CONSTRAINT fk_product_id_for_product_detail
        FOREIGN KEY (product_id)
        REFERENCES Product(product_id)
);

CREATE TABLE ProductImage (
    product_image_id   VARCHAR(20) NOT NULL,
    product_detail_id  VARCHAR(20) NOT NULL,
    product_image      BYTEA,
    PRIMARY KEY (product_image_id),
    CONSTRAINT fk_product_detail_id_for_product_image
        FOREIGN KEY (product_detail_id)
        REFERENCES ProductDetail(product_detail_id)
);

CREATE TABLE Customer (
    customer_id      VARCHAR(20)  NOT NULL,
    customer_name    VARCHAR(50)  DEFAULT 'Chưa có thông tin',
    customer_phone   VARCHAR(20)  DEFAULT 'Chưa có thông tin',
    customer_email   VARCHAR(100) DEFAULT 'Chưa có thông tin', 
    customer_address VARCHAR(100) DEFAULT 'Chưa có thông tin',
    PRIMARY KEY (customer_id)
);

CREATE TABLE Account (
    account_username    VARCHAR(20) NOT NULL,
    customer_id         VARCHAR(20) NOT NULL,
    account_password    VARCHAR(20),
    account_displayname VARCHAR(20),
    role_id             INT CHECK (role_id = 0 OR role_id = 1),
    PRIMARY KEY (account_username),
    CONSTRAINT fk_customer_id_for_account
        FOREIGN KEY (customer_id)
        REFERENCES Customer(customer_id)
);

CREATE TABLE Discount (
    discount_id          VARCHAR(20)  NOT NULL,
    discount_description VARCHAR(200) DEFAULT 'Chưa có thông tin',
    discount_percent     REAL         DEFAULT 0.0,
    discount_date_start  DATE         DEFAULT NOW(),
    discount_date_end    DATE         DEFAULT NOW(),
    PRIMARY KEY (discount_id)
);

CREATE TABLE BillInfo (
    bill_id        VARCHAR(20) NOT NULL,
    customer_id    VARCHAR(20) NOT NULL,
    bill_date      DATE        DEFAULT NOW(),
    bill_status    INT         CHECK (bill_status=0 OR bill_status=1 OR bill_status=2 OR bill_status=3 OR bill_status=4) DEFAULT 0,
    bill_payment   INT         CHECK (bill_payment=0 OR bill_payment=1) DEFAULT 0,
    PRIMARY KEY (bill_id),
    CONSTRAINT fk_customer_id_for_bill_info
        FOREIGN KEY (customer_id)
        REFERENCES Customer(customer_id)
);

CREATE TABLE BillDetail (
    bill_detail_id     VARCHAR(20) NOT NULL,
    bill_id            VARCHAR(20) NOT NULL,
    product_detail_id  VARCHAR(20) NOT NULL,
    discount_id        VARCHAR(20) NOT NULL,
    bill_amount        INT         CHECK (bill_amount > 0),
    PRIMARY KEY (bill_detail_id),
    CONSTRAINT fk_bill_id_for_bill_detail
        FOREIGN KEY (bill_id)
        REFERENCES BillInfo(bill_id),
    CONSTRAINT fk_product_detail_id_for_bill_detail
        FOREIGN KEY (product_detail_id)
        REFERENCES ProductDetail(product_detail_id)
);

\i insert.sql