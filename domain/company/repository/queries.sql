-- name: all
SELECT id, name, image_url as imageUrl, description
FROM company;

-- name: by-id
SELECT id, name, image_url as imageUrl, description
FROM company
WHERE id = ?;