package v1

import (
	pbc "code-service-exam/api-gateway-exam/genproto/comment-proto"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"
)

// COMMENT CREATE COMMENT
// @Summary CREATE A NEW COMMENT
// @Description Api for creating new comment
// @Tags comment
// @Accept json
// @Produce json
// @Param comment query string true "comment" default(I'm writing a comment)
// @Success 200 {object} models.CommentResponse
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/createcomment [post]
func (h *HandlerV1) CreateComment(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	content := c.Query("comment")
	comment_id := uuid.NewString()
	post_id := uuid.NewString()
	user_id := uuid.NewString()

	response, err := h.serviceManager.CommentService().CreateComment(ctx, &pbc.Comment{
		CommentId: comment_id,
		PostId:    post_id,
		UserId:    user_id,
		Content:   content,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error while getting request from Post Service",
		})
		h.log.Error(err.Error())
	}

	c.JSON(http.StatusOK, response)
}

// COMMENT GET COMMENT
// @Summary GET A COMMENT
// @Description Api for getting comment
// @Tags comment
// @Accept json
// @Produce json
// @Param comment_id query string true "comment_id" default(a5f1b10d-849b-40a8-bc18-6e1290e8bcac)
// @Success 200 {object} models.CommentResponse
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/getcomment [get]
func (h *HandlerV1) GetComment(c *gin.Context) {
	var (
		jspbuMarshal protojson.MarshalOptions
	)
	jspbuMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	comment_id := c.Query("comment_id")

	response, err := h.serviceManager.CommentService().GetComment(ctx, &pbc.GetCommentRequest{CommentId: comment_id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while getting response from Post service",
		})
		h.log.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

// COMMENT UPDATE COMMENT
// @Summary UPDATE A COMMENT
// @Description Api for updating comment
// @Tags comment
// @Accept json
// @Produce json
// @Param comment_id query string true "comment_id" default(b7c221f3-81dc-4cb9-779e-7d242212e02d)
// @Success 200 {object} models.UpdateCommentResponse
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/updatecomment [patch]
func (h *HandlerV1) UpdateComment(c *gin.Context) {
	var (
		jspbuMarshal protojson.MarshalOptions
	)
	jspbuMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	comment_id := c.Query("comment_id")

	response, err := h.serviceManager.CommentService().UpdateComment(ctx, &pbc.UpdateCommentRequest{CommentId: comment_id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while updating comment",
		})
		h.log.Error(err.Error())
		return
	}

	if !response.Success {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Unfortunatelly, the comment has not been changed",
			"message": "Sorry, nothing changed. Try again.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Comment successfuly updated",
	})
}

// COMMENT DELETE COMMENT
// @Summary DELETE A COMMENT
// @Description Api for deleting comment
// @Tags comment
// @Accept json
// @Produce json
// @Param comment_id query string true "comment_id" default(a8f3a70b-6e86-4a94-8d9d-7e674f33fe39)
// @Success 200 {object} models.UpdateCommentResponse
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/deletecomment [delete]
func (h *HandlerV1) DeleteComment(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	comment_id := c.Query("comment_id")

	response, err := h.serviceManager.CommentService().DeleteComment(ctx, &pbc.DeleteCommentRequest{CommentId: comment_id})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while deleting comment",
		})
		h.log.Error(err.Error())
		return
	}

	if !response.Success {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Unfortunatelly, the comment has not been deleted",
			"message": "Sorry, nothing deleted. Try again.",
		})
		h.log.Error("information not deleted")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Comment successfuly deleted",
	})
}

// COMMENT LIST COMMENTS
// @Summary GET A LIST OF COMMENTS
// @Description Api for getting a list of comments
// @Tags comment
// @Accept json
// @Produce json
// @Param post_id query string true "post_id" default(f47ac10b-58cc-4372-a567-0e02b2c3d479)
// @Success 200 {object} models.ListCommentsResponse
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/listcomments [get]
func (h *HandlerV1) ListComments(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	post_id := c.Query("post_id")

	response, err := h.serviceManager.CommentService().ListComments(ctx, &pbc.ListCommentsRequest{PostId: post_id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while getting comments from Comment service",
		})
		h.log.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}
