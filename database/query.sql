-- name: CreateUser :one
INSERT INTO
user (NAME, PHONE)
VALUES
    (?, ?)
RETURNING *;

-- name: GetUserByID :one
SELECT
    *
FROM
    user
WHERE
    ID = ?
LIMIT
    1;

-- name: ListUsers :many
SELECT
    *
FROM
    user;

-- name: DeleteUserByID :exec
DELETE FROM
    user
WHERE
    ID = ?;
