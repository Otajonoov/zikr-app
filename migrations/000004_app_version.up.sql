CREATE TABLE IF NOT EXISTS app_version (
    guid UUID PRIMARY KEY,
    android_version VARCHAR(33) NOT NULL,
    ios_version VARCHAR(33) NOT NULL,
    force_update BOOLEAN DEFAULT true
);