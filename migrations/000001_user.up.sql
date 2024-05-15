CREATE TABLE IF NOT EXISTS "users" (
    "guid"              UUID PRIMARY KEY,
    "email"             VARCHAR(255) NOT NULL,
    "unique_username"   VARCHAR(255) NOT NULL
);

