CREATE TABLE "users" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "balance" NUMERIC DEFAULT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);