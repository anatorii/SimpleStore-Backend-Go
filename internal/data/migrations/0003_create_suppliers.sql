create extension if not exists "uuid-ossp";

create table if not exists suppliers (
    id                  uuid primary key default uuid_generate_v4(),
    name                varchar(255) not null,
    address_id          uuid,
    phone_number        varchar(255) not null,
    created_at          timestamptz default current_timestamp,
    updated_at          timestamptz default current_timestamp
);
