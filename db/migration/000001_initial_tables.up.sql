CREATE TABLE "users" (
  "id" serial PRIMARY KEY NOT NULL,
  "username" varchar(255) NOT NULL,
  "password" varchar(255) NOT NULL,
  "email" varchar(255) NOT NULL UNIQUE,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "categories" (
  "id" serial PRIMARY KEY NOT NULL,
  "user_id" int NOT NULL,
  "title" varchar(255) NOT NULL,
  "type" varchar(50) NOT NULL,
  "description" varchar(255) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "categories" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE TABLE "accounts" (
  "id" serial PRIMARY KEY NOT NULL,
  "user_id" int NOT NULL,
  "category_id" int NOT NULL,
  "title" varchar(255) NOT NULL,
  "type" varchar NOT NULL,
  "description" varchar(255) NOT NULL,
  "value" INTEGER NOT NULL,
  "date" DATE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "accounts" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");