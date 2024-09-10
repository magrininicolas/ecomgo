create table if not exists orders_items(
    "id" uuid primary key,
    "order_id" uuid not null,
    "product_id" uuid not null,
    "quantity" integer not null,
    "price" decimal(10, 2) not null,

    foreign key ("order_id") references orders("id"),
    foreign key ("product_id") references products("id")
);