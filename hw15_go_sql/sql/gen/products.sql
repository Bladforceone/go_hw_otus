-- name: ProductCreate :one
INSERT INTO products (name, price, stock)
VALUES ($1, $2, $3)
RETURNING *;
-- name: ProductGet :one
SELECT *
FROM products
WHERE id = $1;
-- name: ProductUpdateStock :exec
UPDATE products
SET stock = $2
WHERE id = $1;
-- name: ProductDelete :exec
DELETE FROM products
WHERE id = $1;