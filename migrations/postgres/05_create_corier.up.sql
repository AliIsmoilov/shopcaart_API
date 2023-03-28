CREATE TABLE courier (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "phone_number"VARCHAR NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);