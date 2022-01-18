// Code generated by sqlc. DO NOT EDIT.
// source: comment.sql

package db

import (
	"context"
)

const createComment = `-- name: CreateComment :one
INSERT INTO comments (
  creator, 
  post_id,
  text
) VALUES (
  $1, $2,$3
) RETURNING id, creator, post_id, text, created_at
`

type CreateCommentParams struct {
	Creator string `json:"creator"`
	PostID  int64  `json:"post_id"`
	Text    string `json:"text"`
}

func (q *Queries) CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error) {
	row := q.db.QueryRowContext(ctx, createComment, arg.Creator, arg.PostID, arg.Text)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.Creator,
		&i.PostID,
		&i.Text,
		&i.CreatedAt,
	)
	return i, err
}

const deleteComment = `-- name: DeleteComment :exec
DELETE FROM comments WHERE id = $1
`

func (q *Queries) DeleteComment(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteComment, id)
	return err
}

const listCommentsByPost = `-- name: ListCommentsByPost :many
SELECT id, creator, post_id, text, created_at FROM comments
where post_id=$1
ORDER BY created_at
`

func (q *Queries) ListCommentsByPost(ctx context.Context, postID int64) ([]Comment, error) {
	rows, err := q.db.QueryContext(ctx, listCommentsByPost, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Comment{}
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.Creator,
			&i.PostID,
			&i.Text,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateComment = `-- name: UpdateComment :one
UPDATE comments 
SET text=$2
WHERE id = $1
RETURNING id, creator, post_id, text, created_at
`

type UpdateCommentParams struct {
	ID   int64  `json:"id"`
	Text string `json:"text"`
}

func (q *Queries) UpdateComment(ctx context.Context, arg UpdateCommentParams) (Comment, error) {
	row := q.db.QueryRowContext(ctx, updateComment, arg.ID, arg.Text)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.Creator,
		&i.PostID,
		&i.Text,
		&i.CreatedAt,
	)
	return i, err
}