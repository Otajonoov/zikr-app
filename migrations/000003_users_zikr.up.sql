CREATE TABLE IF NOT EXISTS users_zikr (
    guid UUID       PRIMARY KEY,
    user_guid       UUID  NOT NULL,
    zikr_guid       UUID  NOT NULL,
    count           BIGINT DEFAULT 0,
    isFavorite      BOOLEAN DEFAULT false,
    CONSTRAINT users_zikr_user_email_foreign FOREIGN KEY (user_guid) REFERENCES users (guid),
    CONSTRAINT users_zikr_zikr_guid_foreign FOREIGN KEY (zikr_guid) REFERENCES zikr (guid)
);
