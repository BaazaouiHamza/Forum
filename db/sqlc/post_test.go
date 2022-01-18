package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/hamza-baazaoui/forum/util"
	"github.com/stretchr/testify/require"
)

func createRandomPost(t *testing.T) Post {
	user := createRandomUser(t)
	arg := CreatePostParams{
		Creator:     user.Username,
		Title:       util.RandomPostTitle(),
		Description: util.RandomPostDescriptionOrText(),
		Image:       util.RandomImage(),
	}

	post, err := testQueries.CreatePost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, post)

	require.Equal(t, arg.Creator, post.Creator)
	require.Equal(t, arg.Title, post.Title)
	require.Equal(t, arg.Description, post.Description)
	require.Equal(t, arg.Image, post.Image)

	require.NotZero(t, post.ID)
	require.NotZero(t, post.CreatedAt)

	return post

}

func TestCreatePost(t *testing.T) {
	createRandomPost(t)
}

func TestGetPost(t *testing.T) {
	post1 := createRandomPost(t)

	post2, err := testQueries.GetPost(context.Background(), post1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, post2)

	require.Equal(t, post1.ID, post2.ID)
	require.Equal(t, post1.Title, post2.Title)
	require.Equal(t, post1.Description, post2.Description)
	require.Equal(t, post1.Creator, post2.Creator)

	require.WithinDuration(t, post1.CreatedAt, post2.CreatedAt, time.Second)
}

func TestUpdatePost(t *testing.T) {
	post1 := createRandomPost(t)

	arg := UpdatePostParams{
		ID:          post1.ID,
		Title:       util.RandomPostTitle(),
		Description: util.RandomPostDescriptionOrText(),
		Image:       util.RandomImage(),
	}

	post2, err := testQueries.UpdatePost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, post2)

	require.Equal(t, arg.ID, post2.ID)
	require.Equal(t, arg.Title, post2.Title)
	require.Equal(t, arg.Description, post2.Description)
	require.Equal(t, post1.Creator, post2.Creator)
	require.WithinDuration(t, post1.CreatedAt, post2.CreatedAt, time.Second)
}
func TestDeletePost(t *testing.T) {
	post1 := createRandomPost(t)

	err := testQueries.DeletePost(context.Background(), post1.ID)
	require.NoError(t, err)

	post2, err := testQueries.GetPost(context.Background(), post1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, post2)

}

func TestListPost(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomPost(t)
	}
	arg := ListPostsParams{
		Limit:  5,
		Offset: 5,
	}
	posts, err := testQueries.ListPosts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, posts, 5)

	for _, post := range posts {
		require.NotEmpty(t, post)
	}
}
