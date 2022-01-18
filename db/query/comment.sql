-- name: CreateComment :one
INSERT INTO comments (
  creator, 
  post_id,
  text
) VALUES (
  $1, $2,$3
) RETURNING *;


-- name: ListCommentsByPost :many
SELECT * FROM comments
where post_id=$1
ORDER BY created_at;


-- name: UpdateComment :one
UPDATE comments 
SET text=$2
WHERE id = $1
RETURNING *;

-- name: DeleteComment :exec
DELETE FROM comments WHERE id = $1;