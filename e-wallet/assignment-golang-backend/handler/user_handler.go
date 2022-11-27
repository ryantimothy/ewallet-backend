package handler

import (
	"ewallet/entity"
	"ewallet/usecase"
	"ewallet/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(usecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		usecase: usecase,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	user := entity.User{}
	err := c.BindJSON(&user)

	if err != nil {

		c.JSON(http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	resp, err := h.usecase.Register(&user)

	if err != nil {

		c.JSON(http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	c.JSON(http.StatusCreated, resp)

}

func (h *UserHandler) Login(c *gin.Context) {
	user := entity.User{}
	err := c.BindJSON(&user)

	if err != nil {

		c.JSON(http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	if user.Email == "" {

		c.JSON(http.StatusBadRequest, "Error: empty email ")
		return
	}

	if user.Password == "" {

		c.JSON(http.StatusBadRequest, "Error: empty password ")
		return
	}

	token, err := h.usecase.Login(user.Email, user.Password)

	if err != nil {

		c.JSON(http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	responseJson := struct {
		Token string `response:"token"`
	}{
		Token: token,
	}

	c.JSON(http.StatusOK, responseJson)
}

func (h *UserHandler) GetDetail(c *gin.Context) {
	tokenID, _ := utils.ExtractTokenID(c)
	id := int(tokenID)

	resp, err := h.usecase.GetDetail(id)

	if err != nil {

		c.JSON(http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}
