create table if not exists "users" (
    "id" UUID PRIMARY KEY,
    "full_name" VARCHAR(50),
    "login" VARCHAR(50),
    "password" VARCHAR(80)
);