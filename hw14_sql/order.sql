WITH new_order AS (
    INSERT INTO
        Orders (user_id, total_amount)
    VALUES
        (2, 199999.98) RETURNING id
)
INSERT INTO
    OrderProducts (order_id, product_id, quantity)
SELECT
    no.id,
    2,
    2
FROM
    new_order no;

DELETE FROM
    Orders
WHERE
    id = 1;