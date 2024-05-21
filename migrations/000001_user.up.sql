CREATE TABLE IF NOT EXISTS "users" (
    "guid"              UUID PRIMARY KEY,
    "email"             VARCHAR(255) NOT NULL,
    "username"          VARCHAR(255) NOT NULL
);

