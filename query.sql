-- name: GetUser :one
SELECT * FROM users WHERE id = ? LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (username, password) VALUES (?, ?) RETURNING *;

-- name: UpdateUser :exec
UPDATE users SET password = ? WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;

-- name: CreateSong :one
INSERT INTO songs (created_by, list_id, link, comment) VALUES (?, ?, ?, ?) RETURNING *;

-- name: UpdateSong :exec
UPDATE songs SET link = ?, comment = ? WHERE id = ?;

-- name: GetSongsByUser :many
SELECT * FROM songs WHERE created_by = ?;

-- name: CountSongsByUser :one
SELECT COUNT(*) FROM songs WHERE created_by = ?;

-- name: CountSongsInList :one
SELECT COUNT(*) FROM songs WHERE list_id = ?;

-- name: CreateList :one
INSERT INTO lists (created_by, name, description, hidden) VALUES (?, ?, ?, ?) RETURNING *;

-- name: UpdateList :exec
UPDATE lists SET name = ?, description = ?, hidden = ? WHERE id = ?;

-- name: DeleteList :exec
DELETE FROM lists WHERE id = ?;

-- name: GetSongsByList :many
SELECT * FROM songs WHERE list_id = ?;

-- name: CreateComment :one
INSERT INTO song_comments (created_by, song_id, comment) VALUES (?, ?, ?) RETURNING *;

-- name: CountCommentsBySong :one
SELECT COUNT(*) FROM song_comments WHERE song_id = ?;

-- name: GetCommentsByUser :many
SELECT * FROM song_comments WHERE created_by = ?;

-- name: GetCommentsBySong :many
SELECT * FROM song_comments WHERE song_id = ?;

-- name: DeleteComment :exec
DELETE FROM song_comments WHERE id = ?;

-- name: CreateSongVote :one
INSERT INTO song_votes (user_id, song_id, value) VALUES (?, ?, ?) RETURNING *;

-- name: UpdateSongVote :exec
UPDATE song_votes SET value = ? WHERE user_id = ? AND song_id = ?;

-- name: DeleteSongVote :exec
DELETE FROM song_votes WHERE user_id = ? AND song_id = ?;

-- name: GetSongVoteScore :one
SELECT SUM(value) FROM song_votes WHERE song_id = ?;