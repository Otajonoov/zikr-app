CREATE TABLE IF NOT EXISTS "zikr" (
    "guid"       UUID PRIMARY KEY,
    "arabic"     TEXT NOT NULL,
    "uzbek"      TEXT NOT NULL,
    "pronounce"  TEXT NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP,
    "user_id"    UUID,
    CONSTRAINT zikr_user_id_foreign FOREIGN KEY(user_id) REFERENCES users(guid)
);
