
-- +migrate Up
CREATE TABLE order_details(
    id SERIAL PRIMARY KEY,
    order_id INT,
    product_sku VARCHAR(255),
    products_qty INT,
    created_at timestamp(6) NOT NULL DEFAULT current_timestamp(6));

-- +migrate Down
DROP TABLE order_details;