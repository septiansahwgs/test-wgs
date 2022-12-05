package repository

var (
	CREATE_ORDER = `INSERT INTO public.orders
	(user_id, total_price)
	VALUES(:user_id, :total_price) RETURNING id`

	CREATE_ORDER_DETAIL = `INSERT INTO public.order_details
	(order_id, product_sku, products_qty)
	VALUES (:order_id, :product_sku, :products_qty)
	`
)
