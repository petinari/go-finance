create table "users"(
    "id" serial primary key not null,
    "username" VARCHAR NOT NULL,
    "password" VARCHAR NOT NULL,
    "email" VARCHAR UNIQUE NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (NOW())
);

create table "categories"(
    "id" serial primary key not null,
    "user_id" int NOT NULL,
    "title" VARCHAR NOT NULL,
    "type" VARCHAR NOT NULL,
    "description" VARCHAR NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (NOW())
);
ALTER TABLE "categories" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");


create table "accounts"(
    "id" serial primary key not null,
    "user_id" int NOT NULL,
    "category_id" INT NOT NULL,
    "title" VARCHAR NOT NULL,
    "type" VARCHAR NOT NULL,
    "description" VARCHAR NOT NULL,
    "value" INTEGER NOT NULL,
    "date" DATE not NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (NOW())
);
ALTER TABLE "accounts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "accounts" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

