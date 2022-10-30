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

func (h *Handler) getTransactions(c *gin.Context) { // filtering handler
	transactionIdParam, ok := c.GetQuery("transaction_id")
	if ok {
		transactionId, err := strconv.Atoi(transactionIdParam)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "TransactionId should be the number.")
			return
		}

		transaction, err := h.services.Transactions.GetByTransactionId(transactionId)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, transaction)
	}

	terminalIdQueryParams, ok := c.GetQueryArray("terminal_id")
	if ok {
		terminalIdParams:= make([]int, len(terminalIdQueryParams))

		for index, terminalId := range terminalIdQueryParams {
			terminalId, err := strconv.Atoi(terminalId)
			if err != nil {
				newErrorResponse(c, http.StatusInternalServerError,
					"TransactionId should be the number.")
				return
			}

			terminalIdParams[index] = terminalId
		}

		transactions, err := h.services.Transactions.GetByTerminalIds(terminalIdParams)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, transactions)
	}

	statusParam, ok := c.GetQuery("status")
	if ok {
		transactions, err := h.services.Transactions.GetByStatus(statusParam)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, transactions)
	}

	paymentTypeParam, ok := c.GetQuery("payment_type")
	if ok {
		transactions, err := h.services.Transactions.GetByPaymentType(paymentTypeParam)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, transactions)
	}


	fromDateParam, fromIs := c.GetQuery("from")
	toDateParam, toIs := c.GetQuery("to")

	if fromIs && toIs {
		transactions, err := h.services.GetByDatePeriod(fromDateParam, toDateParam)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, transactions)
	}

	paymentNarrativeParam, ok := c.GetQuery("payment_narrative")
	if ok {
		transactions, err := h.services.GetByPaymentNarrative(paymentNarrativeParam)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, transactions)
	}
}