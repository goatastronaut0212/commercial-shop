INSERT INTO Category VALUES
('1', 'Quần nam'),
('2', 'Áo nam');

INSERT INTO Product VALUES
('1', '1', 'Quần nam');

INSERT INTO ProductDetail (product_detail_id, product_id) VALUES
('XD32', '1'),
('D33', '1');

INSERT INTO ProductImage (product_image_id, product_detail_id) VALUES
('1', 'XD32'),
('2', 'XD32');

INSERT INTO Discount (discount_id, discount_percent) VALUES
('XV1', 0.05),
('HS1', 0.03);

INSERT INTO AccountRole VALUES
(1, 'Admin'),
(2, 'Khách hàng');

INSERT INTO Account (account_username, account_password, role_id, account_email) VALUES
('noob', '123', 1, 'noob@gmail.com');

INSERT INTO Customer VALUES
('1', 'noob', 'Noob', '123456', 'TP HCM');