-- +goose Up
-- +goose StatementBegin
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    price NUMERIC(11, 2) NOT NULL,
    stock INT NOT NULL DEFAULT 0
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE products;

-- +goose StatementEnd