package handler

import (
	"fmt"
	loader "github.com/Antony201/CsvLoader"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


// @Summary Upload Transactions
// @Tags Transactions
// @Description Upload Transactions file to save to db.
// @ID uploadTransactions
// @Accept mpfd
// @Produce json
// @Param transactions_file formData string true "The transaction file with extension .csv"
// @Success 200 {string} string "message"
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/transactions/upload [post]
func (h *Handler) uploadTransactions(c *gin.Context) {
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

	go h.services.Transactions.LoadFileToDb(transactions_file)

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("Thanks, i am saving the data now."),
	})

}

// @Summary Get transactions by filter
// @Tags Transactions
// @Description Filter Transactions by some parameter.
// @ID getTransactions
// @Accept mpfd
// @Produce json
// @Param transaction_id 	query int    false "Search transaction by transaction_id"
// @Param terminal_id    	query int    false "Search transactions by terminal_id (you can use it multiple)"
// @Param status         	query string false "Search transactions by status"
// @Param payment_type   	query string false "Search transactions by payment_type"
// @Param from   		 	query string false "Search transactions by period (from date to date), use only with to"
// @Param to  		     	query string false "Search transactions by period (from date to date) use only with from"
// @Param payment_narrative query string false "Search transactions by payment_narrative"
// @Success 200 {object} loader.Transaction
// @Success 200 {array} loader.Transaction
// @Failure 500 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/transactions/ [get]
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
		return
	}

	terminalIdQueryParams, ok := c.GetQueryArray("terminal_id")
	if ok {
		terminalIdParams:= make([]int, len(terminalIdQueryParams))

		for index, terminalId := range terminalIdQueryParams {
			terminalId, err := strconv.Atoi(terminalId)
			if err != nil {
				newErrorResponse(c, http.StatusInternalServerError,
					"TerminalId should be the number.")
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
		return
	}

	statusParam, ok := c.GetQuery("status")
	if ok {
		transactions, err := h.services.Transactions.GetByStatus(statusParam)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, transactions)
		return
	}

	paymentTypeParam, ok := c.GetQuery("payment_type")
	if ok {
		transactions, err := h.services.Transactions.GetByPaymentType(paymentTypeParam)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, transactions)
		return
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
		return
	}

	paymentNarrativeParam, ok := c.GetQuery("payment_narrative")
	if ok {
		transactions, err := h.services.GetByPaymentNarrative(paymentNarrativeParam)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, transactions)
		return
	}

	c.JSON(http.StatusOK, loader.Transaction{})
}