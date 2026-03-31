create extension if not exists "uuid-ossp";

create table if not exists addresses (
    id                  uuid primary key default uuid_generate_v4(),
    country             varchar(100) not null,
    city                varchar(100) not null,
    street              varchar(100) not null,
    created_at          timestamptz default current_timestamp,
    updated_at          timestamptz default current_timestamp
);

ALTER TABLE addresses
ADD CONSTRAINT addresses_unique_key
UNIQUE (country, city, street);
