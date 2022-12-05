
-- +migrate Up
CREATE TABLE orders(
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(36),
    total_price DECIMAL,
    created_at timestamp(6) NOT NULL DEFAULT current_timestamp(6));

-- +migrate Down
DROP TABLE orders;