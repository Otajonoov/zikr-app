CREATE TABLE IF NOT EXISTS "zikr" (
    "id" SERIAL PRIMARY KEY,
    "user_id" INT REFERENCES users(id), -- relationship between user and zikr
    "arabic" TEXT NOT NULL,
    "uzbek" TEXT NOT NULL,
    "pronounce" TEXT NOT NULL,
    "is_favorite" BOOLEAN DEFAULT FALSE,
    "created_at" TIMESTAMP(0) WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(0) WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);