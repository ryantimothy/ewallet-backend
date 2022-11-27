package handler

import (
	"errors"
	"ewallet/entity"
	"ewallet/usecase"
	"ewallet/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	usecase usecase.TransactionUsecase
}

func NewTransactionHandler(usecase usecase.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{
		usecase: usecase,
	}
}

func (th *TransactionHandler) GetAll(c *gin.Context) {
	tokenID, _ := utils.ExtractTokenID(c)
	id := int(tokenID)

	q := entity.Query{}
	q.SortBy = c.Query("sortBy")
	if q.SortBy == "" {
		q.SortBy = "created_at"
	}

	q.Sort = c.Query("sort")
	if q.Sort == "" {
		q.Sort = "desc"
	}

	q.Desc = "%" + c.Query("desc") + "%" //where description = ?

	trans, err := th.usecase.GetAll(id, q)

	if err != nil {

		c.JSON(http.StatusBadRequest, errors.New("Error: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, trans)

}

func (th *TransactionHandler) TopUp(c *gin.Context) {
	req := entity.TopUpRequest{}
	err := c.BindJSON(&req)

	if err != nil {

		c.JSON(http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	tokenID, _ := utils.ExtractTokenID(c)
	id := int(tokenID)

	err = th.usecase.TopUp(req, id)
	if err != nil {

		c.JSON(http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, req)
}

func (th *TransactionHandler) Transfer(c *gin.Context) {
	req := entity.TransferRequest{}
	err := c.BindJSON(&req)

	if err != nil {

		c.JSON(http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	tokenID, _ := utils.ExtractTokenID(c)
	id := int(tokenID)

	err = th.usecase.Transfer(req, id)
	if err != nil {

		c.JSON(http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, req)
}
