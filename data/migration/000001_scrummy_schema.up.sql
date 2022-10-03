CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table if not EXISTS users(
    id uuid DEFAULT uuid_generate_v4(),
    
)   