INSERT INTO
    Users (name, email, password)
VALUES
    ('Райн Гослинг', 'rayn@mail.ru', '12345');

UPDATE
    Users
SET
    email = 'gosling@mail.ru'
WHERE
    id = 1;

INSERT INTO
    Products (name, price)
VALUES
    ('Ноутбук', 99999.99);

UPDATE
    Products
SET
    price = 89999.99
WHERE
    id = 1;