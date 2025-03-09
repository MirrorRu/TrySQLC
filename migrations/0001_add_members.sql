-- +goose Up
CREATE EXTENSION IF NOT EXISTS citext;

CREATE SCHEMA members;

CREATE TABLE  members.members_table (
     id BIGINT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY NOT NULL,
     name CITEXT NOT NULL
);

create unique index members_table_name_lower  on members.members_table(lower(name))  ;

-- +goose Down
DROP TABLE members.members_table;

DROP SCHEMA members;
