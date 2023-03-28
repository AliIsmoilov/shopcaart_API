CREATE TABLE products (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "price"NUMERIC,
    "category_id" UUID NOT NULL REFERENCES categories (id),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);