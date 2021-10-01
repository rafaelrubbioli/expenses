-- name: create
INSERT INTO tags (name) VALUES (?);

-- name: by-id
SELECT id, name FROM tags WHERE id = ? LIMIT 1;

-- name: update
UPDATE tags SET name = (?) WHERE id = ?;

-- name: delete
DELETE FROM tags WHERE id = ?;
