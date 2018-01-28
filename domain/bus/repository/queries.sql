-- name: all
SELECT id, number, name, fare, last_update
FROM bus;

-- name: by-id
SELECT id, number, name, fare, last_update
FROM bus where id = ?;

-- name: by-company
SELECT id, number, name, fare, last_update
FROM bus where company_id = ?;