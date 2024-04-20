CREATE TABLE "zikr_count"(
    "user_id" INT REFERENCES users(id),
    "zikr_id" INT REFERENCES zikr(id),
    "count" INT
);
