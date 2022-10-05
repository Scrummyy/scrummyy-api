CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table if not EXISTS auth.users(
    id uuid DEFAULT uuid_generate_v4() not null primary key,
    name varchar(50) not null,
    email varchar(50) unique not null,
    username varchar(50) unique not null,
    password text not null,
    projects uuid[],
    workspace uuid,
    is_superadmin bool DEFAULT false not null
);
