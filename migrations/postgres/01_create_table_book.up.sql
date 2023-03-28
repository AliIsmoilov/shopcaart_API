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