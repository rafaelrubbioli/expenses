-- name: create
INSERT INTO users (name, password) VALUES (?, ?);

-- name: by-id
SELECT id, name FROM users WHERE id = ? LIMIT 1;

-- name: update
UPDATE users SET name = (?) WHERE id = ?;

-- name: delete
DELETE FROM users WHERE id = ?;
