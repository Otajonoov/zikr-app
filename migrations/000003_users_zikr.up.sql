CREATE TABLE IF NOT EXISTS users_zikr (
    user_guid       UUID  NOT NULL,
    zikr_guid       UUID  NOT NULL,
    zikr_count      BIGINT DEFAULT 0,
    isFavorite      BOOLEAN DEFAULT false
);

ALTER TABLE users_zikr
    ADD CONSTRAINT fk_user FOREIGN KEY (user_guid) REFERENCES users (guid) ON DELETE CASCADE;

ALTER TABLE users_zikr
    ADD CONSTRAINT fk_zikr FOREIGN KEY (zikr_guid) REFERENCES zikr (guid) ON DELETE CASCADE;