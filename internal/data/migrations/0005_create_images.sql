create extension if not exists "uuid-ossp";

create table if not exists images (
    id                  uuid primary key default uuid_generate_v4(),
    image               bytea not null,
    description         text,
    created_at          timestamptz default current_timestamp,
    updated_at          timestamptz default current_timestamp
);
