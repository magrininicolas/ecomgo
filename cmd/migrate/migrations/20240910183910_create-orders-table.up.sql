create table if not exists orders(
    "id" uuid not null primary key,
    "user_id" uuid not null,
    "total" decimal(10, 2) not null,
    "status" status_enum not null default 'pending',
    "address" text not null,
    "created_at" timestamp not null,
    "updated_at" timestamp not null,

    foreign key ("user_id") references users("id")
);