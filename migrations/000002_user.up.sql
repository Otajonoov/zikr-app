CREATE TABLE IF NOT EXISTS "users" (
    "guid" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "fio" VARCHAR(255),
    "phone" VARCHAR(255) NOT NULL,
    "uniqe_username"  VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL
);