-- name: CreatePost :one
INSERT INTO posts (
  creator, 
  title,
  description,
  image
) VALUES (
  $1, $2,$3,$4
) RETURNING *;


-- name: GetPost :one
SELECT * FROM posts
WHERE id = $1 LIMIT 1;


-- name: ListPosts :many
SELECT * FROM posts
ORDER BY id
LIMIT $1
OFFSET $2;


-- name: UpdatePost :one
UPDATE posts 
SET title = $2,description=$3,image=$4
WHERE id = $1
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts WHERE id = $1;