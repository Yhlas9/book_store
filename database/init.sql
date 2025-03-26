CREATE TABLE IF NOT EXISTS books
(
    id     serial,
    title  varchar(255) not null,
    author varchar(255) not null
)