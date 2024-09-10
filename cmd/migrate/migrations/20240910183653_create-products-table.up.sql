create table if not exists products(
    "id" uuid not null primary key,
    "name" varchar(255) not null,
    "description" text not null,
    "image" varchar(255) not null,
    "quantity" integer not null,
    "created_at" timestamp not null,
    "updated_at" timestamp not null
)