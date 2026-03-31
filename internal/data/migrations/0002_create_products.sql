create extension if not exists "uuid-ossp";

create table if not exists products (
    id                  uuid primary key default uuid_generate_v4(),
    name                varchar(255) not null,
    category            varchar(255) not null,
    price               decimal(10,2) not null,
    available_stock     integer not null check (available_stock >= 0),
    last_update_date    date not null,
    supplier_id         uuid not null,
    image_id            uuid,
    created_at          timestamptz default current_timestamp,
    updated_at          timestamptz default current_timestamp
);

ALTER TABLE products
ADD CONSTRAINT products_unique_key
UNIQUE (name, category, supplier_id);
