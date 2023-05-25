CREATE DATABASE shopdb;

\c shopdb;

CREATE TABLE category (
    category_id    VARCHAR(20) NOT NULL,
    category_name  VARCHAR(100),
    PRIMARY KEY(category_id)
);

CREATE TABLE product (
    product_id      VARCHAR(20) NOT NULL,
    category_id     VARCHAR(20) NOT NULL,
    product_name    VARCHAR(20),
    product_farbic  VARCHAR(50),
    product_price   INT DEFAULT 0,
    product_detail  VARCHAR(200),
    PRIMARY KEY(product_id),
    CONSTRAINT fk_category_id_for_product
        FOREIGN KEY (category_id)
        REFERENCES category(category_id)
);

CREATE TABLE product_detail (
    product_id     VARCHAR(20) NOT NULL,
    product_size   VARCHAR(20),
    product_amount INT CHECK(product_amount > 0),
    CONSTRAINT fk_product_id_for_product_detail
        FOREIGN KEY (product_id)
        REFERENCES product(product_id)
)

CREATE TABLE product_image (
    product_image_id VARCHAR(20) NOT NULL,
    product_id       VARCHAR(20),
    product_image    BYTEA,
    PRIMARY KEY(product_image_id),
    CONSTRAINT fk_product_id_for_product_image
        FOREIGN KEY (product_id)
        REFERENCES product(product_id)
);

CREATE TABLE customer (
    customer_id      VARCHAR(20) NOT NULL,
    customer_name    VARCHAR(50),
    customer_phone   VARCHAR(20),
    customer_email   VARCHAR(20), 
    customer_address VARCHAR(100),
    PRIMARY KEY(customer_id)
);

CREATE TABLE account (
    account_username    VARCHAR(20) NOT NULL,
    account_password    VARCHAR(20),
    account_displayname VARCHAR(20),
    role_id             INT CHECK (role_id = 0 OR role_id = 1),
    customer_id         VARCHAR(20),
    PRIMARY KEY(account_username),
    CONSTRAINT fk_customer_id_for_account
        FOREIGN KEY (customer_id)
        REFERENCES customer(customer_id)
);

CREATE TABLE discount (
    discount_id          VARCHAR(20) NOT NULL,
    discount_description VARCHAR(20),
    discount_date_start  DATE,
    discount_date_end    DATE,
    PRIMARY KEY(discount_id)
);

CREATE TABLE bill_detail (
    bill_detail_id VARCHAR(20) NOT NULL,
    product_id     VARCHAR(20),
    discount_id    VARCHAR(20),
    PRIMARY KEY (bill_detail_id),
    CONSTRAINT fk_product_id_for_bill_detail
        FOREIGN KEY (product_id)
        REFERENCES product(product_id)
);

CREATE TABLE bill_info (
    bill_id        VARCHAR(20) NOT NULL,
    bill_detail_id VARCHAR(20),
    customer_id    VARCHAR(20),
    bill_date      VARCHAR(20),
    PRIMARY KEY(bill_id),
    CONSTRAINT fk_bill_detail_id_for_bill_info
        FOREIGN KEY (bill_detail_id)
        REFERENCES bill_detail(bill_detail_id)
);

\i insert.sql
