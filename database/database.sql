CREATE DATABASE shopdb;

\c shopdb;

CREATE TABLE category (
    category_id    VARCHAR(20),
    category_name  VARCHAR(100),
    PRIMARY KEY(category_id)
);

CREATE TABLE product (
    product_id     VARCHAR(20),
    category_id    VARCHAR(20),
    product_name   VARCHAR(20),
    product_size   VARCHAR(20),
    product_farbic VARCHAR(40),
    product_price  INT,
    product_detail VARCHAR(200),
    PRIMARY KEY(product_id),
    CONSTRAINT fk_category_id_for_product
      FOREIGN KEY (category_id)
      REFERENCES category(category_id)
);

CREATE TABLE product_image (
    product_image_id VARCHAR(20),
    product_id       VARCHAR(20),
    product_image    BYTEA,
    PRIMARY KEY(product_image_id),
    CONSTRAINT fk_product_id_for_product_image
        FOREIGN KEY (product_id)
        REFERENCES product(product_id)
);
