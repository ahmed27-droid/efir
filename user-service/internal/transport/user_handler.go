package transport

import (
	"errors"
	"net/http"
	"strconv"
	"user/internal/dto"
	"user/internal/errs"
	"user/internal/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{Service: userService}
}

func (h *UserHandler) GetByID(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	user, err := h.Service.GetByID(uint(id))

	if err != nil {
		if errors.Is(err, errs.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "user not found",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve user",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":        user.ID,
		"email":     user.Email,
		"username":  user.Username,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"role":      user.Role,
	})

}

func (h *UserHandler) UpdateProfile(ctx *gin.Context) {

	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	if id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	var req dto.UpdateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	user, err := h.Service.UpdateProfile(uint(id), req)

	if err != nil {
		if errors.Is(err, errs.ErrUsernameAlreadyExists) {
			ctx.JSON(http.StatusConflict, gin.H{
				"error": "Username already taken",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update profile",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Profile updated successfully",
		"id":       user.ID,
		"username": user.Username,
	})
}

func (h *UserHandler) GetMe(ctx *gin.Context) {
	userHeader := ctx.GetHeader("X-User-ID")

	id, err := strconv.Atoi(userHeader)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	user, err := h.Service.GetByID(uint(id))

	if err != nil {
		if errors.Is(err, errs.ErrUserNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"error": "Failed to retrieve user",
	})
	return
}
	
	ctx.JSON(http.StatusOK, gin.H{
		"id":        user.ID,
		"email":     user.Email,
		"username":  user.Username,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"role":      user.Role,
	})

}

/*
админ
удалить пост (соотвественно удалять комментарии, реакции поста)
создать категорию
удалить категорию (удалять все посты внутри категории)
удалить юзера (ридер/автор) (не забывать удалять посты авторов, реакции ридеров)
*/

/*
автор
создавать посты в категориях
удалять свои посты
*/

/*
читатель
ставит реакции
убирает реакции
пишет коментарии
удаляет комментарии 
*/




