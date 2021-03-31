CREATE TABLE IF NOT EXISTS "users" (
    "user_id" serial NOT NULL,
    "user_name" VARCHAR (150) UNIQUE NOT NULL,
    "first_name" varchar(255) NOT NULL,
    "last_name" varchar(255) NOT NULL,
    "email" varchar(512) NOT NULL,
    "password" varchar(1024) NOT NULL,
    "created" TIMESTAMP NOT NULL,
    "changed" TIMESTAMP NOT NULL,
    "is_admin" bool NOT NULL,
CONSTRAINT "users_pk" PRIMARY KEY ("user_id")
) WITH (
OIDS=FALSE
);