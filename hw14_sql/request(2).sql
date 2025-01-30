SELECT
    *
FROM
    Users;

SELECT
    *
FROM
    Products;

SELECT
    o.*,
    u.name
FROM
    Orders o
    JOIN Users u ON o.user_id = u.id
WHERE
    u.id = 1;