-- name: CreatePost :one
INSERT INTO
    posts(
        id,
        feed_id,
        created_at,
        updated_at,
        name,
        description,
        published_at,
        url
    )
VALUES
    ($1, $2, $3, $4, $5 , $6, $7, $8) RETURNING *;

-- name: GetUserPosts :many
SELECT
    posts.*
FROM
    posts
    JOIN feed_follows
    ON posts.feed_id = feed_follows.feed_id
where
    user_id = $1
LIMIT
    $2;