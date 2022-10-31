package handler

import (
	"errors"
	"fmt"
	loader "github.com/Antony201/CsvLoader"
	"github.com/Antony201/CsvLoader/pkg/service"
	mock_service "github.com/Antony201/CsvLoader/pkg/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http/httptest"
	"strconv"
	"testing"
)


func TestHandler_getTransactionById(t *testing.T) {
	type mockBehavior func(s *mock_service.MockTransactions, transaction loader.Transaction,
		id string)

	testTable := []struct {
		name                string
		id					string
		mockBehavior        mockBehavior
		transaction         loader.Transaction
		exceptedStatusCode  int
		exceptedRequestBody string
	} {
		{
			name: "OK",
			id: "1",
			mockBehavior: func(s *mock_service.MockTransactions, transaction loader.Transaction,
				id string) {
				transactionid, _ := strconv.Atoi(id)

				s.EXPECT().GetByTransactionId(transactionid).Return(transaction, nil)
			},
			transaction: loader.Transaction{
				Transactionid: 1,
				Requestid: 20020,
				Terminalid: 3506,
				PartnerObjectid: 1111,
				AmountTotal: 1,
				AmountOriginal: 1,
				CommissionPS: 0,
				CommissionClient: 0,
				CommissionProvider: 0,
				DateInput: "2022-08-12T11:25:27Z",
				DatePost: "2022-08-12T14:25:27Z",
				Status: "accepted",
				PaymentType: "cash",
				PaymentNumber: "PS16698205",
				Serviceid: 13980,
				Service: "Поповнення карток",
				Payeeid: 14232155,
				PayeeName: "pumb",
				PayeeBankMfo: 254751,
				PayeeBankAccount: "UA713451373919523",
				PaymentNarrative: "Перерахування коштів згідно договору про надання послуг А11/27122 від 19.11.2020 р.",
			},
			exceptedStatusCode: 200,
			exceptedRequestBody: `{"Transactionid":1,"Requestid":20020,"Terminalid":3506,"PartnerObjectid":1111,"AmountTotal":1,"AmountOriginal":1,"CommissionPS":0,"CommissionClient":0,"CommissionProvider":0,"DateInput":"2022-08-12T11:25:27Z","DatePost":"2022-08-12T14:25:27Z","Status":"accepted","PaymentType":"cash","PaymentNumber":"PS16698205","Serviceid":13980,"Service":"Поповнення карток","Payeeid":14232155,"PayeeName":"pumb","PayeeBankMfo":254751,"PayeeBankAccount":"UA713451373919523","PaymentNarrative":"Перерахування коштів згідно договору про надання послуг А11/27122 від 19.11.2020 р."}`,
		},
		{
			name: "Bad Transaction Id",
			id: "0",
			mockBehavior: func(s *mock_service.MockTransactions, transaction loader.Transaction,
				id string) {
				transactionid, _ := strconv.Atoi(id)

				s.EXPECT().GetByTransactionId(transactionid).Return(transaction, errors.New(`{"message":"sql: no rows in result set"}`))
			},
			transaction: loader.Transaction{},
			exceptedStatusCode: 500,
			exceptedRequestBody: `{"message":"{\"message\":\"sql: no rows in result set\"}"}`,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			transactions := mock_service.NewMockTransactions(c)
			test.mockBehavior(transactions, test.transaction, test.id)

			services := &service.Service{Transactions: transactions}
			handler := Handler{services}

			r := gin.New()
			r.GET("/transactions/", handler.getTransactions)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET",
				fmt.Sprintf("/transactions/?transaction_id=%s", test.id), nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, test.exceptedStatusCode, w.Code)
			assert.Equal(t, test.exceptedRequestBody, w.Body.String())

		})
	}
}

func TestHandler_getListOfTransactionsWithSeveralQueryParams(t *testing.T) {
	type mockBehavior func(s *mock_service.MockTransactions, transactions []loader.Transaction,
		queryvalue1 string, queryvalue2 string)

	testTable := []struct {
		name                string
		queryParam1         string
		queryParam2         string
		queryValue1         string
		queryValue2         string
		mockBehavior        mockBehavior
		transactions        []loader.Transaction
		exceptedStatusCode  int
		exceptedRequestBody string
	}{
		{
			name:        "Multiple Terminal Id Ok",
			queryParam1: "terminal_id",
			queryParam2: "terminal_id",
			queryValue1:  "3560",
			queryValue2: "3561",
			mockBehavior: func(s *mock_service.MockTransactions, transactions []loader.Transaction,
				queryvalue1 string, queryvalue2 string) {
				value1, _ := strconv.Atoi(queryvalue1)
				value2, _ := strconv.Atoi(queryvalue2)

				s.EXPECT().GetByTerminalIds([]int{value1, value2}).Return(transactions, nil)
			},
			transactions: []loader.Transaction{
				{
					Transactionid: 55,
					Requestid: 20560,
					Terminalid: 3560,
					PartnerObjectid: 1111,
					AmountTotal: 4983,
					AmountOriginal: 4983,
					CommissionPS: 3.49,
					CommissionClient: 0,
					CommissionProvider: -9.97,
					DateInput: "2022-08-23T09:10:01Z",
					DatePost: "2022-08-23T12:10:02Z",
					Status: "declined",
					PaymentType: "cash",
					PaymentNumber: "PS16698745",
					Serviceid: 14520,
					Service: "Поповнення карток",
					Payeeid: 19637555,
					PayeeName: "privat",
					PayeeBankMfo: 308805,
					PayeeBankAccount: "UA713989197718983",
					PaymentNarrative: "Перерахування коштів згідно договору про надання послуг А11/27122 від 19.11.2020 р.",
				},
				{
					Transactionid: 56,
					Requestid: 20570,
					Terminalid: 3561,
					PartnerObjectid: 1111,
					AmountTotal: 1999,
					AmountOriginal: 1999,
					CommissionPS: 1.4,
					CommissionClient: 0,
					CommissionProvider: -4,
					DateInput: "2022-08-23T08:19:21Z",
					DatePost: "2022-08-23T11:19:22Z",
					Status: "accepted",
					PaymentType: "cash",
					PaymentNumber: "PS16698755",
					Serviceid: 14530,
					Service: "Поповнення карток",
					Payeeid: 19737655,
					PayeeName: "pumb",
					PayeeBankMfo: 309806,
					PayeeBankAccount: "UA713999157418973",
					PaymentNarrative: "Перерахування коштів згідно договору про надання послуг А11/27123 від 19.11.2020 р.",
				},
			},
			exceptedStatusCode: 200,
			exceptedRequestBody: `[{"Transactionid":55,"Requestid":20560,"Terminalid":3560,"PartnerObjectid":1111,"AmountTotal":4983,"AmountOriginal":4983,"CommissionPS":3.49,"CommissionClient":0,"CommissionProvider":-9.97,"DateInput":"2022-08-23T09:10:01Z","DatePost":"2022-08-23T12:10:02Z","Status":"declined","PaymentType":"cash","PaymentNumber":"PS16698745","Serviceid":14520,"Service":"Поповнення карток","Payeeid":19637555,"PayeeName":"privat","PayeeBankMfo":308805,"PayeeBankAccount":"UA713989197718983","PaymentNarrative":"Перерахування коштів згідно договору про надання послуг А11/27122 від 19.11.2020 р."},{"Transactionid":56,"Requestid":20570,"Terminalid":3561,"PartnerObjectid":1111,"AmountTotal":1999,"AmountOriginal":1999,"CommissionPS":1.4,"CommissionClient":0,"CommissionProvider":-4,"DateInput":"2022-08-23T08:19:21Z","DatePost":"2022-08-23T11:19:22Z","Status":"accepted","PaymentType":"cash","PaymentNumber":"PS16698755","Serviceid":14530,"Service":"Поповнення карток","Payeeid":19737655,"PayeeName":"pumb","PayeeBankMfo":309806,"PayeeBankAccount":"UA713999157418973","PaymentNarrative":"Перерахування коштів згідно договору про надання послуг А11/27123 від 19.11.2020 р."}]`,
		},
		{
			name:        "From To Ok",
			queryParam1: "from",
			queryParam2: "to",
			queryValue1:  "2022-08-21",
			queryValue2: "2022-08-25",
			mockBehavior: func(s *mock_service.MockTransactions, transactions []loader.Transaction,
				queryvalue1 string, queryvalue2 string) {
				s.EXPECT().GetByDatePeriod(queryvalue1, queryvalue2).Return(transactions, nil)
			},
			transactions: []loader.Transaction{
				{
					Transactionid: 55,
					Requestid: 20560,
					Terminalid: 3560,
					PartnerObjectid: 1111,
					AmountTotal: 4983,
					AmountOriginal: 4983,
					CommissionPS: 3.49,
					CommissionClient: 0,
					CommissionProvider: -9.97,
					DateInput: "2022-08-23T09:10:01Z",
					DatePost: "2022-08-23T12:10:02Z",
					Status: "declined",
					PaymentType: "cash",
					PaymentNumber: "PS16698745",
					Serviceid: 14520,
					Service: "Поповнення карток",
					Payeeid: 19637555,
					PayeeName: "privat",
					PayeeBankMfo: 308805,
					PayeeBankAccount: "UA713989197718983",
					PaymentNarrative: "Перерахування коштів згідно договору про надання послуг А11/27122 від 19.11.2020 р.",
				},
				{
					Transactionid: 56,
					Requestid: 20570,
					Terminalid: 3561,
					PartnerObjectid: 1111,
					AmountTotal: 1999,
					AmountOriginal: 1999,
					CommissionPS: 1.4,
					CommissionClient: 0,
					CommissionProvider: -4,
					DateInput: "2022-08-23T08:19:21Z",
					DatePost: "2022-08-23T11:19:22Z",
					Status: "accepted",
					PaymentType: "cash",
					PaymentNumber: "PS16698755",
					Serviceid: 14530,
					Service: "Поповнення карток",
					Payeeid: 19737655,
					PayeeName: "pumb",
					PayeeBankMfo: 309806,
					PayeeBankAccount: "UA713999157418973",
					PaymentNarrative: "Перерахування коштів згідно договору про надання послуг А11/27123 від 19.11.2020 р.",
				},
			},
			exceptedStatusCode: 200,
			exceptedRequestBody: `[{"Transactionid":55,"Requestid":20560,"Terminalid":3560,"PartnerObjectid":1111,"AmountTotal":4983,"AmountOriginal":4983,"CommissionPS":3.49,"CommissionClient":0,"CommissionProvider":-9.97,"DateInput":"2022-08-23T09:10:01Z","DatePost":"2022-08-23T12:10:02Z","Status":"declined","PaymentType":"cash","PaymentNumber":"PS16698745","Serviceid":14520,"Service":"Поповнення карток","Payeeid":19637555,"PayeeName":"privat","PayeeBankMfo":308805,"PayeeBankAccount":"UA713989197718983","PaymentNarrative":"Перерахування коштів згідно договору про надання послуг А11/27122 від 19.11.2020 р."},{"Transactionid":56,"Requestid":20570,"Terminalid":3561,"PartnerObjectid":1111,"AmountTotal":1999,"AmountOriginal":1999,"CommissionPS":1.4,"CommissionClient":0,"CommissionProvider":-4,"DateInput":"2022-08-23T08:19:21Z","DatePost":"2022-08-23T11:19:22Z","Status":"accepted","PaymentType":"cash","PaymentNumber":"PS16698755","Serviceid":14530,"Service":"Поповнення карток","Payeeid":19737655,"PayeeName":"pumb","PayeeBankMfo":309806,"PayeeBankAccount":"UA713999157418973","PaymentNarrative":"Перерахування коштів згідно договору про надання послуг А11/27123 від 19.11.2020 р."}]`,
		},
	}


	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			transactions := mock_service.NewMockTransactions(c)
			test.mockBehavior(transactions, test.transactions, test.queryValue1, test.queryValue2)

			services := &service.Service{Transactions: transactions}
			handler := Handler{services}

			r := gin.New()
			r.GET("/transactions/", handler.getTransactions)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET",
				fmt.Sprintf("/transactions/?%s=%s&%s=%s",
					test.queryParam1, test.queryValue1,
					test.queryParam2, test.queryValue2), nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, test.exceptedStatusCode, w.Code)
			assert.Equal(t, test.exceptedRequestBody, w.Body.String())

		})
	}
}

func TestHandler_getListOfTransactionsWithOneQueryParam(t *testing.T) {
	type mockBehavior func(s *mock_service.MockTransactions, transactions []loader.Transaction, param string)

	testTable := []struct {
		name string
		queryParam string
		queryValue string
		mockBehavior mockBehavior
		transactions []loader.Transaction
		exceptedStatusCode int
		exceptedRequestBody string
	} {
		{
			name: "Status ok",
			queryParam: "status",
			queryValue: "accepted",
			mockBehavior: func(s *mock_service.MockTransactions, transactions []loader.Transaction, statusParam string ) {
				s.EXPECT().GetByStatus(statusParam).Return(transactions, nil)
			},
			transactions: []loader.Transaction{
				{
					Transactionid:      1,
					Requestid:          20020,
					Terminalid:         3506,
					PartnerObjectid:    1111,
					AmountTotal:        1,
					AmountOriginal:     1,
					CommissionPS:       0,
					CommissionClient:   0,
					CommissionProvider: 0,
					DateInput:          "2022-08-12T11:25:27Z",
					DatePost:           "2022-08-12T14:25:27Z",
					Status:             "accepted",
					PaymentType:        "cash",
					PaymentNumber:      "PS16698205",
					Serviceid:          13980,
					Service:            "Поповнення карток",
					Payeeid:            14232155,
					PayeeName:          "pumb",
					PayeeBankMfo:       254751,
					PayeeBankAccount:   "UA713451373919523",
					PaymentNarrative:   "Перерахування коштів згідно договору про надання послуг А11/27122 від 19.11.2020 р.",
				},
			},
			exceptedStatusCode: 200,
			exceptedRequestBody: `[{"Transactionid":1,"Requestid":20020,"Terminalid":3506,"PartnerObjectid":1111,"AmountTotal":1,"AmountOriginal":1,"CommissionPS":0,"CommissionClient":0,"CommissionProvider":0,"DateInput":"2022-08-12T11:25:27Z","DatePost":"2022-08-12T14:25:27Z","Status":"accepted","PaymentType":"cash","PaymentNumber":"PS16698205","Serviceid":13980,"Service":"Поповнення карток","Payeeid":14232155,"PayeeName":"pumb","PayeeBankMfo":254751,"PayeeBankAccount":"UA713451373919523","PaymentNarrative":"Перерахування коштів згідно договору про надання послуг А11/27122 від 19.11.2020 р."}]`,
		},
		{
			name: "Payment Type Ok",
			queryParam: "payment_type",
			queryValue: "cash",
			mockBehavior: func(s *mock_service.MockTransactions, transactions []loader.Transaction, paymentTypeParam string ) {
				s.EXPECT().GetByPaymentType(paymentTypeParam).Return(transactions, nil)
			},
			transactions: []loader.Transaction{
				{
					Transactionid:      1,
					Requestid:          20020,
					Terminalid:         3506,
					PartnerObjectid:    1111,
					AmountTotal:        1,
					AmountOriginal:     1,
					CommissionPS:       0,
					CommissionClient:   0,
					CommissionProvider: 0,
					DateInput:          "2022-08-12T11:25:27Z",
					DatePost:           "2022-08-12T14:25:27Z",
					Status:             "accepted",
					PaymentType:        "cash",
					PaymentNumber:      "PS16698205",
					Serviceid:          13980,
					Service:            "Поповнення карток",
					Payeeid:            14232155,
					PayeeName:          "pumb",
					PayeeBankMfo:       254751,
					PayeeBankAccount:   "UA713451373919523",
					PaymentNarrative:   "Перерахування коштів згідно договору про надання послуг А11/27122 від 19.11.2020 р.",
				},
			},
			exceptedStatusCode: 200,
			exceptedRequestBody: `[{"Transactionid":1,"Requestid":20020,"Terminalid":3506,"PartnerObjectid":1111,"AmountTotal":1,"AmountOriginal":1,"CommissionPS":0,"CommissionClient":0,"CommissionProvider":0,"DateInput":"2022-08-12T11:25:27Z","DatePost":"2022-08-12T14:25:27Z","Status":"accepted","PaymentType":"cash","PaymentNumber":"PS16698205","Serviceid":13980,"Service":"Поповнення карток","Payeeid":14232155,"PayeeName":"pumb","PayeeBankMfo":254751,"PayeeBankAccount":"UA713451373919523","PaymentNarrative":"Перерахування коштів згідно договору про надання послуг А11/27122 від 19.11.2020 р."}]`,
		},
		{
			name: "Payment Narrative Ok",
			queryParam: "payment_narrative",
			queryValue: "А11/27122",
			mockBehavior: func(s *mock_service.MockTransactions, transactions []loader.Transaction, paymentNarrativeParam string ) {
				s.EXPECT().GetByPaymentNarrative(paymentNarrativeParam).Return(transactions, nil)
			},
			transactions: []loader.Transaction{
				{
					Transactionid:      1,
					Requestid:          20020,
					Terminalid:         3506,
					PartnerObjectid:    1111,
					AmountTotal:        1,
					AmountOriginal:     1,
					CommissionPS:       0,
					CommissionClient:   0,
					CommissionProvider: 0,
					DateInput:          "2022-08-12T11:25:27Z",
					DatePost:           "2022-08-12T14:25:27Z",
					Status:             "accepted",
					PaymentType:        "cash",
					PaymentNumber:      "PS16698205",
					Serviceid:          13980,
					Service:            "Поповнення карток",
					Payeeid:            14232155,
					PayeeName:          "pumb",
					PayeeBankMfo:       254751,
					PayeeBankAccount:   "UA713451373919523",
					PaymentNarrative:   "Перерахування коштів згідно договору про надання послуг А11/27122 від 19.11.2020 р.",
				},
			},
			exceptedStatusCode: 200,
			exceptedRequestBody: `[{"Transactionid":1,"Requestid":20020,"Terminalid":3506,"PartnerObjectid":1111,"AmountTotal":1,"AmountOriginal":1,"CommissionPS":0,"CommissionClient":0,"CommissionProvider":0,"DateInput":"2022-08-12T11:25:27Z","DatePost":"2022-08-12T14:25:27Z","Status":"accepted","PaymentType":"cash","PaymentNumber":"PS16698205","Serviceid":13980,"Service":"Поповнення карток","Payeeid":14232155,"PayeeName":"pumb","PayeeBankMfo":254751,"PayeeBankAccount":"UA713451373919523","PaymentNarrative":"Перерахування коштів згідно договору про надання послуг А11/27122 від 19.11.2020 р."}]`,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			transactions := mock_service.NewMockTransactions(c)
			test.mockBehavior(transactions, test.transactions, test.queryValue)

			services := &service.Service{Transactions: transactions}
			handler := Handler{services}

			r := gin.New()
			r.GET("/transactions/", handler.getTransactions)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET",
				fmt.Sprintf("/transactions/?%s=%s",
					test.queryParam, test.queryValue), nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, test.exceptedStatusCode, w.Code)
			assert.Equal(t, test.exceptedRequestBody, w.Body.String())

		})
	}
}