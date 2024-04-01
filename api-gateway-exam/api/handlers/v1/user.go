package v1

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"

	models "code-service-exam/api-gateway-exam/api/handlers/models"
	"code-service-exam/api-gateway-exam/api/tokens"
	pbu "code-service-exam/api-gateway-exam/genproto/user-proto"
	l "code-service-exam/api-gateway-exam/pkg/logger"
)

// CREATE USER
// @Summary CREATE USER
// @Description Api for creating user
// @Tags user
// @Accept json
// @Produce json
// @Param User body models.CreateUserRequest true "User"
// @Success 200 {object} models.CreateUserResponse
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/createuser [post]
func (h *HandlerV1) CreateUser(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
		body        *models.CreateUserResponse
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

	// user ro'yhatdan o'tayotganda u tanlagan pozitsiyasiga ko'ra unga token generatsiya qilinadi
	h.jwtHandler = tokens.JWTHandler{
		Sub:       body.FirstName,
		Role:      "user",
		SignInKey: "uzbekcodewizard",
		Log:       h.log,
	}

	id := uuid.NewString()

	response, err := h.serviceManager.UserService().CreateUser(ctx, &pbu.User{
		Id:        id,
		Username:  body.Username,
		Email:     body.Email,
		Password:  body.Password,
		CreatedAt: body.CreatedAt.Format("02-01-2006 15:04:05"),
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Bio:       body.Bio,
		Website:   body.Website,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error while getting request from Post Service",
		})
		h.log.Error("Creating new post error")
	}

	c.JSON(http.StatusOK, response)
}

// GET USER
// @Summary GET USER BY EMAIL
// @Description Api for getting user
// @Tags user
// @Accept json
// @Produce json
// @Param email query string true "email" default(user1@example.com)
// @Success 200 {object} models.CreateUserResponse
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/getuser [get]
func (h *HandlerV1) GetUser(c *gin.Context) {
	var (
		jspbuMarshal protojson.MarshalOptions
	)
	jspbuMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	// converting string to int
	email := c.Query("email")

	response, err := h.serviceManager.UserService().GetUser(ctx, &pbu.GetUserReq{Email: email})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while getting response from User service",
		})
		h.log.Error("error while getting response from User service")
		return
	}

	c.JSON(http.StatusOK, response)
}

// LIST USERS
// @Summary LIST USERS BY PAGE AND LIMIT
// @Description Api for list of  user
// @Tags user
// @Accept json
// @Produce json
// @Param page query string true "page" default(1)
// @Param limit query string true "limit" default(10)
// @Success 200 {object} models.ListUsersResponse
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/listusers [get]
func (h *HandlerV1) ListUsers(c *gin.Context) {
	var (
		jspbuMarshal protojson.MarshalOptions
	)
	jspbuMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	page := c.Query("page")
	limit := c.Query("limit")

	pageS, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while converting page to int",
		})
		h.log.Error("error while converting page to int")
		return
	}

	limitS, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while converting limit to int",
		})
		h.log.Error("error while converting limit to int")
		return
	}

	response, err := h.serviceManager.UserService().ListUsers(ctx, &pbu.ListUsersRequest{Page: int64(pageS), Limit: int64(limitS)})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while getting response from User service",
		})
		h.log.Error("error while getting response from User service")
		return
	}

	c.JSON(http.StatusOK, response)
}

// UPDATE USER
// @Summary UPDATE USER BY EMAIL
// @Description Api for updating user
// @Tags user
// @Accept json
// @Produce json
// @Param email query string true "email" default(user2@example.com)
// @Success 200 {object} models.UpdateUserResponse
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/updateuser [patch]
func (h *HandlerV1) UpdateUser(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	email := c.Query("email")

	response, err := h.serviceManager.UserService().UpdateUser(ctx, &pbu.UpdateUserReq{Email: email})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while updating user info",
		})
		h.log.Error("user info updating error")
		return
	}

	if !response.Success {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Unfortunatelly, the info not updated",
			"message": "Sorry, nothing updated. Try again.",
		})
		h.log.Error("information not updated")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User successfuly updated",
	})
}

// DELETE USER
// @Summary DELETE USER BY EMAIL
// @Description Api for deleting user
// @Tags user
// @Accept json
// @Produce json
// @Param email query string true "email" default(user3@example.com)
// @Success 200 {object} models.DeleteUserResponse
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/deleteuser [delete]
func (h *HandlerV1) DeleteUser(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	email := c.Query("email")

	response, err := h.serviceManager.UserService().DeleteUser(ctx, &pbu.DeleteUserReq{Email: email})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while deleting user",
		})
		h.log.Error("user deleting error")
		return
	}

	if !response.Success {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Unfortunatelly, the user not deleted",
			"message": "Sorry, nothing deleted. Try again.",
		})
		h.log.Error("information not deleted")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User successfuly deleted from the Database",
	})
}
