CREATE TABLE IF NOT EXISTS "users" (
    "id" SERIAL PRIMARY KEY,
    "fio" VARCHAR(255),
    "phone" VARCHAR(255) NOT NULL,
    "uniqe_username"  VARCHAR(255) UNIQUE NOT NULL,
    "password" VARCHAR(255) NOT NULL
);