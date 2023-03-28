CREATE TABLE customers (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "phone"VARCHAR NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);