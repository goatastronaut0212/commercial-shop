CREATE DATABASE shopdb;

\c shopdb;

CREATE TABLE category (
    category_id    VARCHAR(20) NOT NULL,
    category_name  VARCHAR(100),
    PRIMARY KEY (category_id)
);

CREATE TABLE product (
    product_id      VARCHAR(20)  NOT NULL,
    category_id     VARCHAR(20)  NOT NULL,
    product_name    VARCHAR(20)  DEFAULT 'Chưa có thông tin',
    PRIMARY KEY (product_id),
    CONSTRAINT fk_category_id_for_product
        FOREIGN KEY (category_id)
        REFERENCES category(category_id)
);

CREATE TABLE product_detail (
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
        REFERENCES product(product_id)
);

CREATE TABLE product_image (
    product_image_id   VARCHAR(20) NOT NULL,
    product_detail_id  VARCHAR(20) NOT NULL,
    product_image      BYTEA,
    PRIMARY KEY (product_image_id),
    CONSTRAINT fk_product_detail_id_for_product_image
        FOREIGN KEY (product_detail_id)
        REFERENCES product_detail(product_detail_id)
);

CREATE TABLE customer (
    customer_id      VARCHAR(20) NOT NULL,
    customer_name    VARCHAR(50),
    customer_phone   VARCHAR(20),
    customer_email   VARCHAR(20), 
    customer_address VARCHAR(100),
    PRIMARY KEY (customer_id)
);

CREATE TABLE account (
    account_username    VARCHAR(20) NOT NULL,
    customer_id         VARCHAR(20) NOT NULL,
    account_password    VARCHAR(20),
    account_displayname VARCHAR(20),
    role_id             INT CHECK (role_id = 0 OR role_id = 1),
    PRIMARY KEY (account_username),
    CONSTRAINT fk_customer_id_for_account
        FOREIGN KEY (customer_id)
        REFERENCES customer(customer_id)
);

CREATE TABLE discount (
    discount_id          VARCHAR(20) NOT NULL,
    discount_description VARCHAR(200) DEFAULT 'Chưa có thông tin',
    discount_percent     REAL DEFAULT 0.0,
    discount_date_start  DATE DEFAULT NOW(),
    discount_date_end    DATE DEFAULT NOW(),
    PRIMARY KEY (discount_id)
);

CREATE TABLE bill_info (
    bill_id        VARCHAR(20) NOT NULL,
    customer_id    VARCHAR(20) NOT NULL,
    bill_date      VARCHAR(20) DEFAULT NOW(),
    PRIMARY KEY (bill_id)
);

CREATE TABLE bill_detail (
    bill_detail_id VARCHAR(20) NOT NULL,
    bill_id        VARCHAR(20) NOT NULL,
    product_id     VARCHAR(20) NOT NULL,
    discount_id    VARCHAR(20) NOT NULL,
    PRIMARY KEY (bill_detail_id),
    CONSTRAINT fk_bill_id_for_bill_detail
        FOREIGN KEY (bill_id)
        REFERENCES bill_info(bill_id),
    CONSTRAINT fk_product_id_for_bill_detail
        FOREIGN KEY (product_id)
        REFERENCES product(product_id)
);

\i insert.sql
