CREATE TABLE IF NOT EXISTS "zikr" (
    "guid"       UUID PRIMARY KEY,
    "user_guid"  UUID REFERENCES users(guid) ON DELETE CASCADE,
    "arabic"     TEXT NOT NULL,
    "uzbek"      TEXT NOT NULL,
    "pronounce"  TEXT NOT NULL,
    "is_favorite" BOOLEAN DEFAULT FALSE,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP
);
