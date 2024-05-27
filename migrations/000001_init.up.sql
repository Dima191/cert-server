create table cert
(
    id                integer generated always as identity primary key,
    domain            varchar not null unique,
    private_key       varchar not null,
    certificate_chain varchar not null
);

create index idx_domain on cert (domain);