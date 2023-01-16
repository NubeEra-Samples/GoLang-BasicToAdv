CREATE TABLE "Student" (
  "ID" integer PRIMARY KEY
);

CREATE TABLE "Teacher" (
  "TID" integer PRIMARY KEY,
  "ID" integer
);

ALTER TABLE "Teacher" ADD FOREIGN KEY ("ID") REFERENCES "Student" ("ID");
