create extension if not exists "uuid-ossp";

create table if not exists clients (
    id                  uuid primary key default uuid_generate_v4(),
    client_name         varchar(100) not null,
    client_surname      varchar(100) not null,
    birthday            date not null,
    gender              varchar(1) not null,
    registration_date   date not null,
    address_id          uuid,
    created_at          timestamptz default current_timestamp,
    updated_at          timestamptz default current_timestamp
);

ALTER TABLE clients
ADD CONSTRAINT clients_unique_key
UNIQUE (client_name, client_surname, birthday, gender);
