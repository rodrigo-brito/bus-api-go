-- name: all
SELECT id, number, name, fare
FROM bus;

-- name: by-id
SELECT id, number, name, fare
FROM bus where id = ?;