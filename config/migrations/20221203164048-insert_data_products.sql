
-- +migrate Up
INSERT INTO public.products
(sku, "name", price, inventory_qty)
VALUES
('120P90', 'Google Home', 49.99, 10),
('43N23P', 'MacBook Pro', 5399.99, 5),
('A304SD', 'Alexa Speaker', 109.50, 10),
('234324', 'Raspberry Pi B', 30.00, 10);

-- +migrate Down
