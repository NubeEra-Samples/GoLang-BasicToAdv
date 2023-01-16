CREATE TABLE "Users" (
  "Id" serial PRIMARY KEY
);

CREATE TABLE "Posts" (
  "Id" serial PRIMARY KEY,
  "UserId" serial
);

ALTER TABLE "Posts" ADD FOREIGN KEY ("UserId") REFERENCES "Users" ("Id");
