-- name: CreateMember :one
INSERT INTO members.members_table(
    name
    , birthday
) VALUES(
    @name, @birthday
) RETURNING *
;

-- name: UpdateMember :one
UPDATE members.members_table
SET
    name = @name
    , birthday = @birthday
WHERE
    id = sqlc.arg(id)
RETURNING *
;

-- name: DeleteMember :one
DELETE FROM members.members_table
WHERE
    id = sqlc.arg(id)
RETURNING *
;

-- name: GetMember :one
SELECT *
FROM members.members_table
WHERE
    id = sqlc.arg(id)
LIMIT 1
;

-- name: GetMembersList :many
SELECT *
FROM members.members_table mt
WHERE
    sqlc.narg(name)::text is null
    OR
    mt.name ILIKE('%'||sqlc.narg(name)::text||'%')
ORDER BY
    mt.name
;

-- name: GetMembersList2 :many
SELECT *
FROM members.members_table mt
WHERE
    mt.name like sqlc.arg(conditions)
ORDER BY
    mt.name
;