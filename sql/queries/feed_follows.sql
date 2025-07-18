-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (
        id,
        created_at,
        updated_at,
        user_id,
        feed_id
    ) VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    ) RETURNING *
)
SELECT  inserted_feed_follow.*,
        users.name AS user_name,
        feeds.name AS feed_name
FROM inserted_feed_follow
INNER JOIN feeds ON feeds.id = inserted_feed_follow.feed_id
INNER JOIN users ON users.id = inserted_feed_follow.user_id;

-- name: GetFeedFollowsForUser :many
SELECT  feed_follows.*,
        users.name AS user_name,
        feeds.name AS feed_name
FROM feed_follows
INNER JOIN users ON users.id = feed_follows.user_id
INNER JOIN feeds ON feeds.id = feed_follows.feed_id
WHERE feed_follows.user_id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
WHERE feed_follows.user_id = $1
AND feed_follows.feed_id = (
    SELECT feeds.id
    FROM feeds
    WHERE feeds.url = $2
);
