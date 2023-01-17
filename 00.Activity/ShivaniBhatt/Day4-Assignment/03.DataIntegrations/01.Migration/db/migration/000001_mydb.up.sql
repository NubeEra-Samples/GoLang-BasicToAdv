CREATE TABLE "Users" (
  "UserID" integer PRIMARY KEY
);

CREATE TABLE "Posts" (
  "PostId" Integer PRIMARY KEY,
  "UserID" integer
);

ALTER TABLE "Posts" ADD FOREIGN KEY ("UserID") REFERENCES "Users" ("UserID");
