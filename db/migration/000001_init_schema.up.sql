CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "posts" (
  "id" bigserial PRIMARY KEY,
  "creator" varchar NOT NULL,
  "title" varchar NOT NULL,
  "description" varchar NOT NULL,
  "image" varchar NOT NULL,
  "vote_count" integer DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "comments" (
  "id" bigserial PRIMARY KEY,
  "creator" varchar NOT NULL,
  "post_id" bigint NOT NULL,
  "text" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "posts" ADD FOREIGN KEY ("creator") REFERENCES "users" ("username");

ALTER TABLE "comments" ADD FOREIGN KEY ("creator") REFERENCES "users" ("username");

ALTER TABLE "comments" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");
