CREATE TABLE IF NOT EXISTS "product" (
    id          SERIAL PRIMARY KEY,
    created_at  TIMESTAMP NOT NULL DEFAULT (now()),
    updated_at  TIMESTAMP NOT NULL DEFAULT (now()),
    deleted_at  TIMESTAMP,
    code        VARCHAR(255) NOT NULL,
    name        VARCHAR(255) NOT NULL,
    description TEXT,
    price       NUMERIC(10, 2) NOT NULL,
    active      BOOLEAN NOT NULL
);

CREATE TYPE user_role AS ENUM ('admin', 'customer', 'guest');

CREATE TABLE IF NOT EXISTS "user" (
    id          SERIAL PRIMARY KEY,
    created_at  TIMESTAMP NOT NULL DEFAULT (now()),
    updated_at  TIMESTAMP NOT NULL DEFAULT (now()),
    deleted_at  TIMESTAMP,
    email       VARCHAR(255) NOT NULL,
    password    VARCHAR(255) NOT NULL,
    role        user_role NOT NULL
);

CREATE TYPE order_status AS ENUM ('new', 'in-progress', 'done', 'cancelled');

CREATE TABLE IF NOT EXISTS "order" (
    id          SERIAL PRIMARY KEY,
    created_at  TIMESTAMP NOT NULL DEFAULT (now()),
    updated_at  TIMESTAMP NOT NULL DEFAULT (now()),
    deleted_at  TIMESTAMP,
    code        VARCHAR(255) NOT NULL,
    user_id     SERIAL REFERENCES "user"(id),
    total_price NUMERIC(10, 2) NOT NULL,
    status      order_status NOT NULL
);

CREATE TABLE IF NOT EXISTS order_line (
    id        SERIAL PRIMARY KEY,
    created_at  TIMESTAMP NOT NULL DEFAULT (now()),
    updated_at  TIMESTAMP NOT NULL DEFAULT (now()),
    deleted_at  TIMESTAMP,
    order_id   SERIAL REFERENCES "order"(id),
    product_id SERIAL REFERENCES "product"(id),
    quantity   INT NOT NULL,
    price      NUMERIC(10, 2) NOT NULL
);