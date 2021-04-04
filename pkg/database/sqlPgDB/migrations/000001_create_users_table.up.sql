CREATE TABLE IF NOT EXISTS "users"
(
    "user_id"    serial              NOT NULL,
    "user_name"  VARCHAR(150) UNIQUE NOT NULL,
    "first_name" varchar(255),
    "last_name"  varchar(255),
    "email"      varchar(512)        NOT NULL,
--     "password" varchar(1024) NOT NULL,
    "password"   bytea               NOT NULL,
    "created"    TIMESTAMP           NOT NULL,
    "changed"    TIMESTAMP           NOT NULL,
    "is_admin"   bool                NOT NULL default false,
    CONSTRAINT "users_pk" PRIMARY KEY ("user_id")
) WITH (
      OIDS= FALSE
    );