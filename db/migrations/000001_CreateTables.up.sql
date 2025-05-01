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

CREATE TABLE IF NOT EXISTS "dossier" (
    "id" UUID NOT NULL DEFAULT uuid_generate_v4(),
    "owner_id" UUID NOT NULL UNIQUE,
    FOREIGN KEY ("owner_id") REFERENCES "profile"("id") ON DELETE CASCADE,
    PRIMARY KEY ("id","owner_id"),
    "face_type" TEXT NOT NULL,
    "skin_tone" TEXT NOT NULL,
    "body_type" TEXT NOT NULL,
    "gender" TEXT NOT NULL,
    "preferred_colors" TEXT[],  
    "disliked_colors" TEXT[],   
    "height" TEXT,              
    "weight" TEXT               
);

CREATE TABLE IF NOT EXISTS "dossier_access" (
    "dossier_id" UUID NOT NULL,
    "profile_id" UUID NOT NULL,
    PRIMARY KEY ("dossier_id", "profile_id"),
    FOREIGN KEY ("dossier_id") REFERENCES "dossier"("id") ON DELETE CASCADE,
    FOREIGN KEY ("profile_id") REFERENCES "profile"("id") ON DELETE CASCADE
);

COMMIT;
