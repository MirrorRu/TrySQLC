-- +goose Up
ALTER TABLE  members.members_table
    ADD COLUMN birthday DATE
;

-- +goose Down
ALTER TABLE  members.members_table
    DROP COLUMN birthday
;
