CREATE TABLE IF NOT EXISTS user_zikr (
    user_id UUID NOT NULL,
    zikr_id UUID NOT NULL,
    count INT,
    isfavorite BOOLEAN,
    PRIMARY KEY(user_id, zikr_id),
    CONSTRAINT user_zikr_user_id_foreign FOREIGN KEY(user_id) REFERENCES users(guid),
    CONSTRAINT user_zikr_zikr_id_foreign FOREIGN KEY(zikr_id) REFERENCES zikr(guid)
);
