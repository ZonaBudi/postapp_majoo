package reportservice_test

import (
	"context"
	"net/http"
	"postapp/internal/domain"
	"postapp/internal/port/mock"
	"postapp/internal/service/reportservice"
	"postapp/pkg/model"
	"postapp/pkg/paginator"
	"postapp/pkg/response"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type ReportTestSuite struct {
	suite.Suite
	Log                       *zap.Logger
	User                      *domain.User
	Merchant                  *domain.Merchant
	Outlet                    *domain.Outlet
	MockTransactionRepository *mock.MockTransactionRepository
	MockMerchantRepository    *mock.MockMerchantRepository
	MockOutletRepository      *mock.MockOutletRepository
}

func (suite *ReportTestSuite) SetupTest() {
	mockCtrl := gomock.NewController(suite.T())
	suite.Log, _ = zap.NewProduction()
	var id uint64 = 1
	suite.User = &domain.User{
		Base: model.Base{
			ID: &id,
		},
		Name:     "test",
		UserName: "test",
		Password: "e00cf25ad42683b3df678c61f42c6bda",
	}

	suite.Merchant = &domain.Merchant{
		Base: model.Base{
			ID: &id,
		},
		UserID:       1,
		MerchantName: "test",
	}

	suite.Outlet = &domain.Outlet{
		Base: model.Base{
			ID: &id,
		},
		MerchantID: 1,
		OutletName: "test",
	}

	suite.MockTransactionRepository = mock.NewMockTransactionRepository(mockCtrl)
	suite.MockMerchantRepository = mock.NewMockMerchantRepository(mockCtrl)
	suite.MockOutletRepository = mock.NewMockOutletRepository(mockCtrl)
}

func (suite *ReportTestSuite) TestReportMerchant() {
	var time = time.Now()
	cases := []struct {
		name       string
		wantResult []*domain.TransactionMerchant
		err        error
		params     *domain.TransactionMerchantFilter
	}{
		{
			name:       "failed_get_merchant",
			wantResult: nil,
			err:        response.NewError(http.StatusInternalServerError, response.WithMessageError(response.ErrDefault)),
			params:     &domain.TransactionMerchantFilter{},
		},
		{
			name:       "merchant_not_found",
			wantResult: nil,
			err:        response.NewError(http.StatusNotFound, response.WithMessageError(reportservice.ErrMerchantNotFound)),
			params:     &domain.TransactionMerchantFilter{},
		},
		{
			name:       "failed_get_transaction",
			wantResult: nil,
			err:        response.NewError(http.StatusInternalServerError, response.WithMessageError(response.ErrDefault)),
			params: &domain.TransactionMerchantFilter{
				MerchantID: 1,
			},
		},
		{
			name: "success_get_transaction",
			wantResult: []*domain.TransactionMerchant{
				{
					TransactionDate: "2020-11-01",
					Summary:         9500,
				},
			},
			err: nil,
			params: &domain.TransactionMerchantFilter{
				MerchantID: 1,
				StartDate:  &time,
				EndDate:    &time,
				Pagination: &paginator.Pagination{
					Limit:        10,
					Page:         1,
					Offset:       0,
					NextPage:     1,
					PreviousPage: 1,
					Count:        1,
					TotalPage:    1,
				},
			},
		},
	}
	for _, tCase := range cases {
		switch tCase.name {
		case "failed_get_merchant":
			suite.Run(tCase.name, func() {
				suite.MockMerchantRepository.EXPECT().
					FindOneByUserID(gomock.Any(), gomock.Any()).
					Return(nil, tCase.err).
					Times(1)
				reportService := reportservice.NewReportService(suite.Log, suite.MockTransactionRepository, suite.MockMerchantRepository, suite.MockOutletRepository)
				_, _, err := reportService.ReportMerchant(context.Background(), *suite.User.ID, tCase.params)
				assert.Error(suite.T(), err)
			})
		case "merchant_not_found":
			suite.Run(tCase.name, func() {
				suite.MockMerchantRepository.EXPECT().
					FindOneByUserID(gomock.Any(), gomock.Any()).
					Return(nil, nil).
					Times(1)
				reportService := reportservice.NewReportService(suite.Log, suite.MockTransactionRepository, suite.MockMerchantRepository, suite.MockOutletRepository)
				_, _, err := reportService.ReportMerchant(context.Background(), *suite.User.ID, tCase.params)
				assert.Error(suite.T(), err)
			})
		case "failed_get_transaction":
			suite.Run(tCase.name, func() {
				suite.MockMerchantRepository.EXPECT().
					FindOneByUserID(gomock.Any(), gomock.Any()).
					Return(suite.Merchant, nil).
					Times(1)
				suite.MockTransactionRepository.EXPECT().
					TransactionByMerchant(gomock.Any(), gomock.Any()).
					Return(nil, nil, tCase.err).
					Times(1)
				reportService := reportservice.NewReportService(suite.Log, suite.MockTransactionRepository, suite.MockMerchantRepository, suite.MockOutletRepository)
				_, _, err := reportService.ReportMerchant(context.Background(), *suite.User.ID, tCase.params)
				assert.Error(suite.T(), err)
			})
		case "success_get_transaction":
			suite.Run(tCase.name, func() {
				suite.MockMerchantRepository.EXPECT().
					FindOneByUserID(gomock.Any(), gomock.Any()).
					Return(suite.Merchant, nil).
					Times(1)
				suite.MockTransactionRepository.EXPECT().
					TransactionByMerchant(gomock.Any(), gomock.Any()).
					Return(tCase.wantResult, tCase.params, nil).
					Times(1)
				reportService := reportservice.NewReportService(suite.Log, suite.MockTransactionRepository, suite.MockMerchantRepository, suite.MockOutletRepository)
				_, _, err := reportService.ReportMerchant(context.Background(), *suite.User.ID, tCase.params)
				assert.NoError(suite.T(), err)
			})
		}
	}
}

func (suite *ReportTestSuite) TestReportOutlet() {
	var time = time.Now()
	cases := []struct {
		name       string
		wantResult []*domain.TransactionOutlet
		err        error
		params     *domain.TransactionOutletFilter
	}{
		{
			name:       "failed_get_merchant",
			wantResult: nil,
			err:        response.NewError(http.StatusInternalServerError, response.WithMessageError(response.ErrDefault)),
			params:     &domain.TransactionOutletFilter{},
		},
		{
			name:       "merchant_not_found",
			wantResult: nil,
			err:        response.NewError(http.StatusNotFound, response.WithMessageError(reportservice.ErrMerchantNotFound)),
			params:     &domain.TransactionOutletFilter{},
		},
		{
			name:       "failed_get_outlet",
			wantResult: nil,
			err:        response.NewError(http.StatusInternalServerError, response.WithMessageError(response.ErrDefault)),
			params:     &domain.TransactionOutletFilter{},
		},
		{
			name:       "outlet_not_found",
			wantResult: nil,
			err:        response.NewError(http.StatusNotFound, response.WithMessageError(reportservice.ErrOutletNotFound)),
			params:     &domain.TransactionOutletFilter{},
		},
		{
			name:       "failed_not_my_outlet",
			wantResult: nil,
			err:        response.NewError(http.StatusForbidden, response.WithMessageError(reportservice.ErrOutletNotForYou)),
			params:     &domain.TransactionOutletFilter{},
		},
		{
			name:       "failed_get_transaction",
			wantResult: nil,
			err:        response.NewError(http.StatusInternalServerError, response.WithMessageError(response.ErrDefault)),
			params:     &domain.TransactionOutletFilter{},
		},
		{
			name: "success_get_transaction",
			wantResult: []*domain.TransactionOutlet{
				{
					TransactionDate: "2020-11-01",
					Summary:         9500,
				},
			},
			err: nil,
			params: &domain.TransactionOutletFilter{
				OutletID:  1,
				StartDate: &time,
				EndDate:   &time,
				Pagination: &paginator.Pagination{
					Limit:        10,
					Page:         1,
					Offset:       0,
					NextPage:     1,
					PreviousPage: 1,
					Count:        1,
					TotalPage:    1,
				},
			},
		},
	}
	for _, tCase := range cases {
		switch tCase.name {
		case "failed_get_merchant":
			suite.Run(tCase.name, func() {
				suite.MockMerchantRepository.EXPECT().
					FindOneByUserID(gomock.Any(), gomock.Any()).
					Return(nil, tCase.err).
					Times(1)
				reportService := reportservice.NewReportService(suite.Log, suite.MockTransactionRepository, suite.MockMerchantRepository, suite.MockOutletRepository)
				_, _, err := reportService.ReportOutlet(context.Background(), *suite.User.ID, tCase.params)
				assert.Error(suite.T(), err)
			})
		case "merchant_not_found":
			suite.Run(tCase.name, func() {
				suite.MockMerchantRepository.EXPECT().
					FindOneByUserID(gomock.Any(), gomock.Any()).
					Return(nil, nil).
					Times(1)
				reportService := reportservice.NewReportService(suite.Log, suite.MockTransactionRepository, suite.MockMerchantRepository, suite.MockOutletRepository)
				_, _, err := reportService.ReportOutlet(context.Background(), *suite.User.ID, tCase.params)
				assert.Error(suite.T(), err)
			})
		case "failed_get_outlet":
			suite.Run(tCase.name, func() {
				suite.MockMerchantRepository.EXPECT().
					FindOneByUserID(gomock.Any(), gomock.Any()).
					Return(suite.Merchant, nil).
					Times(1)
				suite.MockOutletRepository.EXPECT().
					FindOneByOutletID(gomock.Any(), gomock.Any()).
					Return(nil, tCase.err).
					Times(1)
				reportService := reportservice.NewReportService(suite.Log, suite.MockTransactionRepository, suite.MockMerchantRepository, suite.MockOutletRepository)
				_, _, err := reportService.ReportOutlet(context.Background(), *suite.User.ID, tCase.params)
				assert.Error(suite.T(), err)
			})
		case "outlet_not_found":
			suite.Run(tCase.name, func() {
				suite.MockMerchantRepository.EXPECT().
					FindOneByUserID(gomock.Any(), gomock.Any()).
					Return(suite.Merchant, nil).
					Times(1)
				suite.MockOutletRepository.EXPECT().
					FindOneByOutletID(gomock.Any(), gomock.Any()).
					Return(nil, nil).
					Times(1)
				reportService := reportservice.NewReportService(suite.Log, suite.MockTransactionRepository, suite.MockMerchantRepository, suite.MockOutletRepository)
				_, _, err := reportService.ReportOutlet(context.Background(), *suite.User.ID, tCase.params)
				assert.Error(suite.T(), err)
			})
		case "failed_not_my_outlet":
			suite.Run(tCase.name, func() {
				merchant := suite.Merchant
				var merchantID uint64 = 2
				merchant.ID = &merchantID
				outlet := suite.Outlet
				suite.MockMerchantRepository.EXPECT().
					FindOneByUserID(gomock.Any(), gomock.Any()).
					Return(merchant, nil).
					Times(1)
				suite.MockOutletRepository.EXPECT().
					FindOneByOutletID(gomock.Any(), gomock.Any()).
					Return(outlet, nil).
					Times(1)
				reportService := reportservice.NewReportService(suite.Log, suite.MockTransactionRepository, suite.MockMerchantRepository, suite.MockOutletRepository)
				_, _, err := reportService.ReportOutlet(context.Background(), *suite.User.ID, tCase.params)
				assert.Error(suite.T(), err)
			})
		case "failed_get_transaction":
			suite.Run(tCase.name, func() {
				var merchantID uint64 = 1
				suite.Merchant.ID = &merchantID
				suite.MockMerchantRepository.EXPECT().
					FindOneByUserID(gomock.Any(), gomock.Any()).
					Return(suite.Merchant, nil).
					Times(1)
				suite.MockOutletRepository.EXPECT().
					FindOneByOutletID(gomock.Any(), gomock.Any()).
					Return(suite.Outlet, nil).
					Times(1)
				suite.MockTransactionRepository.EXPECT().
					TransactionByOutlet(gomock.Any(), gomock.Any()).
					Return(nil, nil, tCase.err).
					Times(1)
				reportService := reportservice.NewReportService(suite.Log, suite.MockTransactionRepository, suite.MockMerchantRepository, suite.MockOutletRepository)
				_, _, err := reportService.ReportOutlet(context.Background(), *suite.User.ID, tCase.params)
				assert.Error(suite.T(), err)
			})
		case "success_get_transaction":
			suite.Run(tCase.name, func() {
				suite.MockMerchantRepository.EXPECT().
					FindOneByUserID(gomock.Any(), gomock.Any()).
					Return(suite.Merchant, nil).
					Times(1)
				suite.MockOutletRepository.EXPECT().
					FindOneByOutletID(gomock.Any(), gomock.Any()).
					Return(suite.Outlet, nil).
					Times(1)
				suite.MockTransactionRepository.EXPECT().
					TransactionByOutlet(gomock.Any(), gomock.Any()).
					Return(tCase.wantResult, tCase.params, nil).
					Times(1)
				reportService := reportservice.NewReportService(suite.Log, suite.MockTransactionRepository, suite.MockMerchantRepository, suite.MockOutletRepository)
				_, _, err := reportService.ReportOutlet(context.Background(), *suite.User.ID, tCase.params)
				assert.NoError(suite.T(), err)
			})
		}
	}
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(ReportTestSuite))
}
