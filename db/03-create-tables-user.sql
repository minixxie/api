DROP TABLE IF EXISTS "User";
CREATE TABLE if not exists "User" (
    "id" SERIAL PRIMARY KEY,
    "phone" VARCHAR(40) NOT NULL,
    "password" VARCHAR(40) NOT NULL,

    "createdAtMicroseconds" BIGINT NOT NULL,
    "updatedAtMicroseconds" BIGINT NOT NULL
);
ALTER TABLE "User" ADD CONSTRAINT "phoneIsUnique" UNIQUE ("phone");

