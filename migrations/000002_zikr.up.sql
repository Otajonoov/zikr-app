CREATE TABLE IF NOT EXISTS "zikr" (
    "guid"       UUID PRIMARY KEY,
    "user_guid"  UUID REFERENCES users(guid) ON DELETE CASCADE,
    "arabic"     TEXT NOT NULL,
    "uzbek"      TEXT NOT NULL,
    "pronounce"  TEXT NOT NULL,
    "count"      BIGINT CHECK("count" >= 0),
    "is_favorite" BOOLEAN DEFAULT FALSE,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);
