CREATE TABLE "users" (
  "id" serial PRIMARY KEY NOT NULL,
  "userName" varchar(255) NOT NULL,
  "password" varchar(255) NOT NULL,
  "email" varchar(255) NOT NULL UNIQUE,
  "createdAt" timestamptz NOT NULL DEFAULT (now()),
  "updatedAt" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "categories" (
  "id" serial PRIMARY KEY NOT NULL,
  "userId" int NOT NULL,
  "title" varchar(255) NOT NULL,
  "type" varchar(50) NOT NULL,
  "description" varchar(255) NOT NULL,
  "createdAt" timestamptz NOT NULL DEFAULT (now()),
  "updatedAt" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "categories" ADD FOREIGN KEY ("userId") REFERENCES "users" ("id");

CREATE TABLE "accounts" (
  "id" serial PRIMARY KEY NOT NULL,
  "userId" int NOT NULL,
  "categoryId" int NOT NULL,
  "title" varchar(255) NOT NULL,
  "description" varchar(255) NOT NULL,
  "value" INTEGER NOT NULL,
  "date" DATE NOT NULL,
  "createdAt" timestamptz NOT NULL DEFAULT (now()),
  "updatedAt" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("userId") REFERENCES "users" ("id");
ALTER TABLE "accounts" ADD FOREIGN KEY ("categoryId") REFERENCES "categories" ("id");