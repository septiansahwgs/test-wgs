
-- +migrate Up
CREATE TABLE products(
    id SERIAL PRIMARY KEY,
    sku VARCHAR(255),
    name VARCHAR (255),
    price DECIMAL,
    inventory_qty int,
    created_at timestamp(6) NOT NULL DEFAULT current_timestamp(6));

-- +migrate Down
DROP TABLE products;