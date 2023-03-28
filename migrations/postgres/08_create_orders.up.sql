CREATE TABLE "orders" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "price" NUMERIC ,
    "phone_number" VARCHAR NOT NULL,
    "latitude" NUMERIC NOT NULL,
    "longtitude" NUMERIC NOT NULL, 
    "user_id" UUID NOT NULL REFERENCES users (id),
    "customer_id" UUID NOT NULL REFERENCES customers (id),
    "courier_id" UUID NOT NULL REFERENCES courier (id),
    "product_id" UUID NOT NULL REFERENCES products (id),
    "quantity" NUMERIC,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);