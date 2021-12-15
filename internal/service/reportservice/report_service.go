package reportservice

import (
	"context"
	"errors"
	"net/http"
	"postapp/internal/domain"
	"postapp/internal/port"
	"postapp/pkg/response"
	"time"

	"go.uber.org/zap"
)

var (
	ErrMerchantNotFound = errors.New("merchant not found")
	ErrOutletNotFound   = errors.New("outlet not found")
	ErrOutletNotForYou  = errors.New("forbidden, you are not the owner of this outlet")
)

type reportService struct {
	log              *zap.Logger
	transactionMysql port.TransactionRepository
	merchantMysql    port.MerchantRepository
	outletMysql      port.OutletRepository
}

func NewReportService(log *zap.Logger,
	transactionMysql port.TransactionRepository,
	merchantMysql port.MerchantRepository,
	outletMysql port.OutletRepository,
) port.ReportService {
	return &reportService{
		log:              log,
		transactionMysql: transactionMysql,
		merchantMysql:    merchantMysql,
		outletMysql:      outletMysql,
	}
}

func (_instance *reportService) ReportMerchant(ctx context.Context, userID uint64, filter *domain.TransactionMerchantFilter) (*domain.ReportMerchant, *domain.TransactionMerchantFilter, error) {
	merchant, err := _instance.merchantMysql.FindOneByUserID(ctx, userID)
	if err != nil {
		_instance.log.Error("failed get data transaction merchant : %v", zap.Error(err))
		return nil, nil, response.NewError(http.StatusInternalServerError, response.WithMessageError(response.ErrDefault))
	}
	if merchant.IsEmpty() {
		return nil, nil, response.NewError(http.StatusNotFound, response.WithMessageError(ErrMerchantNotFound))
	}
	filter.MerchantID = *merchant.ID
	if filter.StartDate == nil && filter.EndDate == nil {
		firstOfMonth := time.Date(2021, time.November, 1, 0, 0, 0, 0, time.Local)
		lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
		filter.StartDate = &firstOfMonth
		filter.EndDate = &lastOfMonth
	}
	transaction, filter, err := _instance.transactionMysql.TransactionByMerchant(ctx, filter)
	if err != nil {
		_instance.log.Error("failed get data transaction merchant : %v", zap.Error(err))
		return nil, nil, response.NewError(http.StatusInternalServerError, response.WithMessageError(response.ErrDefault))
	}
	return &domain.ReportMerchant{
		Merchant: &domain.PublicMerchant{
			MerchantName: merchant.MerchantName,
		},
		Transactions: transaction,
	}, filter, nil
}

func (_instance *reportService) ReportOutlet(ctx context.Context, userID uint64, filter *domain.TransactionOutletFilter) (*domain.ReportOutlet, *domain.TransactionOutletFilter, error) {
	merchant, err := _instance.merchantMysql.FindOneByUserID(ctx, userID)
	if err != nil {
		_instance.log.Error("failed get data transaction merchant : %v", zap.Error(err))
		return nil, nil, response.NewError(http.StatusInternalServerError, response.WithMessageError(response.ErrDefault))
	}
	if merchant.IsEmpty() {
		return nil, nil, response.NewError(http.StatusNotFound, response.WithMessageError(ErrMerchantNotFound))
	}

	outlet, err := _instance.outletMysql.FindOneByOutletID(ctx, filter.OutletID)
	if err != nil {
		_instance.log.Error("failed get data transaction merchant : %v", zap.Error(err))
		return nil, nil, response.NewError(http.StatusInternalServerError, response.WithMessageError(response.ErrDefault))
	}
	if outlet.IsEmpty() {
		return nil, nil, response.NewError(http.StatusNotFound, response.WithMessageError(ErrOutletNotFound))
	}

	if !outlet.MyOutlet(*merchant.ID) {
		return nil, nil, response.NewError(http.StatusForbidden, response.WithMessageError(ErrOutletNotForYou))
	}
	filter.OutletID = *outlet.ID
	if filter.StartDate == nil && filter.EndDate == nil {
		firstOfMonth := time.Date(2021, time.November, 1, 0, 0, 0, 0, time.Local)
		lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
		filter.StartDate = &firstOfMonth
		filter.EndDate = &lastOfMonth
	}
	transaction, filter, err := _instance.transactionMysql.TransactionByOutlet(ctx, filter)
	if err != nil {
		_instance.log.Error("failed get data transaction merchant : %v", zap.Error(err))
		return nil, nil, response.NewError(http.StatusInternalServerError, response.WithMessageError(response.ErrDefault))
	}
	return &domain.ReportOutlet{
		Merchant: &domain.PublicMerchant{
			MerchantName: merchant.MerchantName,
		},
		Outlet: &domain.PublicOutlet{
			OutletName: outlet.OutletName,
		},
		Transactions: transaction,
	}, filter, nil
}
