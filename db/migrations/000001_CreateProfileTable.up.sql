BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "profile" (
    "id" UUID NOT NULL UNIQUE PRIMARY KEY,
    "email" VARCHAR (100) NOT NULL UNIQUE,
    "full_name" VARCHAR (150) NOT NULL,
    "phone_number" VARCHAR (15) NOT NULL UNIQUE,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP(3) NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS "style_profile" (
    "id" UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    "owner_id" UUID NOT NULL UNIQUE,
    FOREIGN KEY ("owner_id") REFERENCES "profile"("id") ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "style_profile_access" (
    "style_profile_id" UUID NOT NULL,
    "profile_id" UUID NOT NULL,
    PRIMARY KEY ("style_profile_id", "profile_id"),
    FOREIGN KEY ("style_profile_id") REFERENCES "style_profile"("id") ON DELETE CASCADE,
    FOREIGN KEY ("profile_id") REFERENCES "profile"("id") ON DELETE CASCADE
);

COMMIT;
