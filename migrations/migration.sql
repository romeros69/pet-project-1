CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table if not exists book (
    id uuid DEFAULT uuid_generate_v4(),
    tittle varchar(255) not null ,
    author varchar(255)
);