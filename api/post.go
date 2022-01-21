package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/hamza-baazaoui/forum/db/sqlc"
)

type createPostRequest struct {
	Creator     string `json:"creator"`
	Title       string `json:"title" binding:"required,min=6"`
	Description string `json:"description" binding:"required"`
	Image       string `json:"image" binding:"required"`
}

func (server *Server) createPost(ctx *gin.Context) {
	var req createPostRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreatePostParams{
		Creator:     req.Creator,
		Title:       req.Title,
		Description: req.Description,
		Image:       req.Image,
	}

	post, err := server.store.CreatePostTx(ctx.Request.Context(), arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, post)

}

type getPostRquest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getPostById(ctx *gin.Context) {
	var req getPostRquest
	if err := ctx.BindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	post, err := server.store.GetPost(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, post)
}

type listPostRequest struct {
	PageId   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listPost(ctx *gin.Context) {
	var req listPostRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListPostsParams{
		Limit:  req.PageSize,
		Offset: (req.PageId - 1) * req.PageSize,
	}

	posts, err := server.store.ListPosts(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, posts)
}

type deletePostRquest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deletePost(ctx *gin.Context) {
	var req deletePostRquest
	if err := ctx.BindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeletePost(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "post deleted succesfully",
	})
}

type updatePostRequestBody struct {
	Title       string `json:"title" binding:"min=6"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type updatePostRequestUri struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) updatePost(ctx *gin.Context) {
	var reqBody updatePostRequestBody
	var reqUri updatePostRequestUri
	if err := ctx.BindUri(&reqUri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdatePostParams{
		ID:          reqUri.ID,
		Title:       reqBody.Title,
		Description: reqBody.Description,
		Image:       reqBody.Image,
	}
	post, err := server.store.UpdatePost(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, post)
}
