CREATE TABLE "Users" (
  "UserID" integer PRIMARY KEY
);

CREATE TABLE "Posts" (
  "PostId" integer PRIMARY KEY,
  "address" varchar(100) UNIQUE NOT NULL
);

COMMENT ON COLUMN "Posts"."address" IS 'client address';
