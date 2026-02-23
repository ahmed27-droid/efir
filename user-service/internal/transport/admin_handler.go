package transport

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"user/internal/services"

	"github.com/gin-gonic/gin"
)

/*
админ
удалить пост (соотвественно удалять комментарии, реакции поста)
создать категорию
удалить категорию (удалять все посты внутри категории)
удалить юзера (ридер/автор) (не забывать удалять посты авторов, реакции ридеров)
*/

type AdminHandler struct{
	service services.AdminService
}

func NewAdminHandler(adminService services.AdminService) *AdminHandler {
	return &AdminHandler{
		service: adminService,
	}
}

func (h *AdminHandler) AdminDeletePost(c *gin.Context) {
	postID := c.Param("id")

	url := "http://localhost:8081/posts/" + postID

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create request",
		})
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to contact post service",
		})
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{
			"error": "post service returned error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "post deleted successfully",
	})
}

func (h *AdminHandler) AdminCreateCategory(c *gin.Context) {

	var req dto.CreateCategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}
	url := "http://localhost:8082/categories"

	jsonData, err := json.Marshal(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to encode request",
		})
		return
	}

	reqHTTP, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create request",
		})
		return
	}

	reqHTTP.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(reqHTTP)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "failed to contact category service",
		})
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		c.JSON(resp.StatusCode, gin.H{
			"error": "category service returned error",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "category created successfully",
	})

}

func (h *AdminHandler) AdminDeleteCategory(c *gin.Context) {
	categoryID := c.Param("id")

	if categoryID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid category id",
		})
		return
	}

	client := &http.Client{}

	postURL := "http://localhost:8081/posts/by-category/" + categoryID

	postReq, err := http.NewRequest(http.MethodDelete, postURL, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create post delete request",
		})
		return
	}

	postResp, err := client.Do(postReq)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "failed to contact post service",
		})
		return
	}

	defer postResp.Body.Close()

	if postResp.StatusCode != http.StatusOK {
		c.JSON(postResp.StatusCode, gin.H{
			"error": "post service returned error",
		})
		return
	}

	categoryURL := "http://localhost:8082/categories/" + categoryID

	categoryReq, err := http.NewRequest(http.MethodDelete, categoryURL, nil)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create category delete request",
		})
		return
	}

	categoryResp, err := client.Do(categoryReq)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "failed to contact category service",
		})
		return
	}
	defer categoryResp.Body.Close()

	if categoryResp.StatusCode != http.StatusOK {
		c.JSON(categoryResp.StatusCode, gin.H{
			"error": "category service returned error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "category and its posts deleted successfuly",
	})
}

func (h *AdminHandler) AdninDeleteUser(c *gin.Context) {

	userID := c.Param("id")

	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user id",
		})
		return
	}

	client := &http.Client{}

	postURL := "http://localhost:8081/posts/by-author/" + userID

	postReq, _ := http.NewRequest(http.MethodDelete, postURL, nil)
	postResp, err := client.Do(postReq)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "failed to contact post service",
		})
		return
	}

	defer postResp.Body.Close()

	if postResp.StatusCode != http.StatusOK {
		c.JSON(postResp.StatusCode, gin.H{
			"error": "post service returned error",
		})
		return
	}

	reactionURL := "http://localhost:8083/reactions/by-user/" + userID

	reactionReq, _ := http.NewRequest(http.MethodDelete, reactionURL, nil)
	reactionResp, err := client.Do(reactionReq)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "failed to contact reaction service",
		})
		return
	}
	defer reactionResp.Body.Close()

	if reactionResp.StatusCode != http.StatusOK {
		c.JSON(reactionResp.StatusCode, gin.H{
			"error": "reaction service returned error", 
		})
		return
	}

	uID, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "failed to execute internal work",
		})
		return
	}

	err = h.service.DeleteUser(uint(uID))
	if err != nil {
		c.JSON(500, gin.H{
			"error": "failed to delete user",
		})
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"message": "user and related data deleted successfully",
	})
	
}

// как делать http запросы в go
// мне нужно сделать delete запрос по определенному пути, объясни как это сделать
