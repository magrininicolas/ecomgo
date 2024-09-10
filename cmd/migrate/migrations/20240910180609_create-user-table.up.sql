create table if not exists users(
    "id" UUID primary key,
    "first_name" varchar(255) not null,
    "last_name" varchar(255) not null,
    "email" varchar(255) not null unique,
    "password" varchar(255) not null,
    "created_at" timestamp not null,
    "updated_at" timestamp not null
);