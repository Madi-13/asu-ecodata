CREATE TABLE IF NOT EXISTS mdm_auth.resources(
    code varchar(100) primary key,
    title varchar(100) not null,
    description text,
    created_at  timestamptz default now(),
    updated_at  timestamptz default now()
)