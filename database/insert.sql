INSERT INTO category VALUES
('1', 'Quần nam'),
('2', 'Áo nam');

INSERT INTO product VALUES
('1', '1', 'Quần nam');

INSERT INTO product_detail (product_detail_id, product_id) VALUES
('XD32', '1'),
('D33', '1');

INSERT INTO product_image (product_image_id, product_detail_id) VALUES
('1', 'XD32'),
('2', 'XD32');