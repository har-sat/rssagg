-- name: CreateUserFeed :one
INSERT INTO feed_follows(id, created_at, updated_at, user_id, feed_id)
VALUES
    ($1, $2, $3, $4, $5) 
RETURNING *;


-- name: GetUserFeedFollows :many
SELECT * FROM feed_follows WHERE user_id = $1;

-- name: DeleteUserFeedFollow :exec
DELETE FROM feed_follows WHERE user_id = $1 and feed_id = $2;