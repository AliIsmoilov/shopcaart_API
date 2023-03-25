
CREATE TABLE "book" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "price" NUMERIC NOT NULL,
    "count" INTEGER,
    "came_price" NUMERIC NOT NULL,
    "profit_status" VARCHAR NOT NULL,
    "profit" NUMERIC,
    "sell_price" NUMERIC NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);


CREATE TABLE "users" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "balance" NUMERIC DEFAULT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);


CREATE TABLE "author" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)


