create table if not exists mdm_auth.permissions
(
    id            int primary key generated always as identity,
    code          varchar(100)                           not null check ( code ~ '^[a-zA-Z0-9_\-]+$' ),
    title         varchar(200),
    role_name     varchar(200)                           not null,
    resource_code varchar(200)                           not null
        references mdm_auth.resources,
    scopes        text[]                                 not null
        constraint permissions_scopes_check
            check (scopes <@ ARRAY ['read'::text, 'write'::text, 'delete'::text, 'update'::text]),
    created_at    timestamp with time zone default now() not null,
    updated_at    timestamp with time zone default now() not null
);

CREATE UNIQUE INDEX  permissions_code_role_name_uidx on mdm_auth.permissions(code, role_name);
