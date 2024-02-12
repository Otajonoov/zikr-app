CREATE TABLE IF NOT EXISTS "zikr" (
                                      "id" UUID NOT NULL,
                                      "arabic" TEXT NOT NULL,
                                      "uzbek" TEXT NOT NULL,
                                      "pronounce" TEXT NOT NULL,
                                      PRIMARY KEY ("id")
    );