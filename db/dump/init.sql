CREATE TABLE IF NOT EXISTS orders (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    status VARCHAR(100) NOT NULL,
    quantity INT UNSIGNED NOT NULL,
    product_id BIGINT UNSIGNED NOT NULL
);


INSERT INTO orders (status, quantity, product_id) VALUES
('Pending', 2, 1),
('Shipped', 1, 2),
('Delivered', 5, 3);
