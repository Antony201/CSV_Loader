package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
	"net/http"
	"strconv"
	"test_task"
	//"os"
	//"path/filepath"
)


func (h *Handler) uploadTransactions(c *gin.Context) {
	transactions := []test_task.Transaction{}

	uploadedFile, err := c.FormFile("transactions_file")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	transactions_file, err := uploadedFile.Open()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := gocsv.Unmarshal(transactions_file, &transactions); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	defer transactions_file.Close()


	rows, err := h.services.Transactions.Create(transactions)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("File was uploaded and data is saved in db! (%d rows)", rows),
	})
}

func (h *Handler) getTransactions(c *gin.Context) {
	var resultTransaction test_task.Transaction

	transactionId, ok := c.GetQuery("transaction_id")
	if ok {
		transactionId, err := strconv.Atoi(transactionId)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		transaction, err := h.services.Transactions.GetByTransactionId(transactionId)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, transaction)

	} else {
		fmt.Println(resultTransaction)

		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "NICE",
		})
	}
}