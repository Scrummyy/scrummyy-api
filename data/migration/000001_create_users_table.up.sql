CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE SCHEMA IF NOT EXISTS auth;

create table if not EXISTS auth.users(
    id uuid DEFAULT uuid_generate_v4() not null primary key,
    name varchar(50) not null,
    email varchar(50) unique not null,
    username varchar(50) unique not null,
    password text not null,
    projects uuid[],
    workspace uuid,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP
);
