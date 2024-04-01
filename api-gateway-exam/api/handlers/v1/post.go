package v1

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"

	models "code-service-exam/api-gateway-exam/api/handlers/models"
	"code-service-exam/api-gateway-exam/api/tokens"
	pbp "code-service-exam/api-gateway-exam/genproto/post-proto"

	l "code-service-exam/api-gateway-exam/pkg/logger"
)

// POST CREATE POST
// @Summary CREATE A NEW POST
// @Description Api for creating new post
// @Tags post
// @Accept json
// @Produce json
// @Param Post body models.Post true "Post"
// @Success 200 {object} models.Post
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/createpost [post]
func (h *HandlerV1) CreatePost(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
		body        *models.Post
	)

	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	resPost, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error while marshalling body",
		})
		h.log.Error("marshalling body error")
	}

	log.Println(string(resPost))

	h.jwtHandler = tokens.JWTHandler{
		Sub:       body.Title,
		Role:      "admin",
		SignInKey: "uzbekcodewizard",
		Log:       h.log,
	}

	acces, refresh, err := h.jwtHandler.GenerateAuthJWT()
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "error while generating jwt",
		})
		h.log.Error("error generate new jwt tokens", l.Error(err))
		return
	}

	post_id := uuid.NewString()
	user_id := uuid.NewString()

	response, err := h.serviceManager.PostService().CreatePost(ctx, &pbp.Post{
		PostId:       post_id,
		UserId:       user_id,
		Content:      body.Content,
		Title:        body.Title,
		Likes:        body.Likes,
		Dislikes:     body.Dislikes,
		Views:        body.Views,
		MediaUrl:     body.MediaUrl,
		RefreshToken: refresh,
	})

	respBody := &models.Post{
		PostId:       response.PostId,
		UserId:       response.UserId,
		Content:      response.Content,
		Title:        response.Title,
		CreatedAt:    response.CreatedAt,
		Likes:        response.Likes,
		Dislikes:     response.Dislikes,
		Views:        response.Views,
		MediaUrl:     response.MediaUrl,
		AccesToken:   acces,
		RefreshToken: response.RefreshToken,
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error while getting request from Post Service",
		})
		h.log.Error(err.Error())
	}

	c.JSON(http.StatusOK, respBody)
}

// POST GET POST BY ID
// @Summary GET POST BY POST_ID WITH COMMENTS
// @Description Api for getting one post by Id
// @Tags post
// @Accept json
// @Produce json
// @Param post_id query string true "post_id" default(f47ac10b-58cc-4372-a567-0e02b2c3d479)
// @Success 200 {object} models.GetPostByIdResponse
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/getpostbyid [get]
func (h *HandlerV1) GetPostById(c *gin.Context) {
	var (
		jspbuMarshal protojson.MarshalOptions
	)
	jspbuMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	post_id := c.Query("post_id")

	response, err := h.serviceManager.PostService().GetPostById(ctx, &pbp.GetPostByIdReq{PostId: post_id})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while getting response from Post service",
		})
		h.log.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

// POST GET POSTS BY USER ID
// @Summary GET POSTS BY USER_ID WITH COMMENTS
// @Description Api for getting posts by user_id
// @Tags post
// @Accept json
// @Produce json
// @Param user_id query string true "user_id" default(647ac10b-58cc-4372-a567-0e02b2c3d479)
// @Success 200 {object} models.GetPostByIdResponse
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/getpostsbyuserid [get]
func (h *HandlerV1) GetPostsByUserId(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	user_id := c.Query("user_id")

	response, err := h.serviceManager.PostService().GetPostsByUserId(ctx, &pbp.GetPostsByUserIdReq{UserId: user_id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while getting products from Product service",
		})
		h.log.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

// POST UPDATE POST CONTENT
// @Summary UPDATE POST CONTENT BY POST_ID
// @Description Api for updating post content
// @Tags post
// @Accept json
// @Produce json
// @Param post_id query string true "post_id" default(aa5f7cb8-ae5f-4df3-907a-6e1c4389c8b1)
// @Param newcontent query string true "newcontent" default(updated content)
// @Success 200 {object} models.UpdateContentResponse
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/updatepostcontent [patch]
func (h *HandlerV1) UpdatePost(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	post_id := c.Query("post_id")
	content := c.Query("content")

	response, err := h.serviceManager.PostService().UpdatePost(ctx, &pbp.UpdatePostReq{
		PostId:  post_id,
		Content: content,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while updating post content",
		})
		h.log.Error(err.Error())
		return
	}

	if !response.Success {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Unfortunatelly, the content has not been changed",
			"message": "Sorry, nothing changed. Try again.",
		})
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Content successfuly updated",
	})
}

// POST DELETE POST BY POST_ID
// @Summary DELETE POST BY POST_ID
// @Description Api for deleting post
// @Tags post
// @Accept json
// @Produce json
// @Param post_id query string true "post_id" default(bb0a3189-3e58-4f57-b7b1-8d6ac4fb326f)
// @Success 200 {object} models.DeletePostResponse
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/deletepost [delete]
func (h *HandlerV1) DeletePost(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	post_id := c.Query("post_id")

	response, err := h.serviceManager.PostService().DeletePost(ctx, &pbp.DeletePostReq{
		PostId: post_id,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while deleting post",
		})
		h.log.Error(err.Error())
		return
	}

	if !response.Success {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Unfortunatelly, the post has not been deleted",
			"message": "Sorry, nothing deleted. Try again.",
		})
		h.log.Error("information not deleted")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post successfuly deleted",
	})
}
