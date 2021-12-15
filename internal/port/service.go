package port

import (
	"context"
	"postapp/internal/domain"
)

type (
	AuthService interface {
		Login(ctx context.Context, req *domain.LoginRequest) (*domain.Login, error)
	}
	ReportService interface {
		ReportMerchant(ctx context.Context, userID uint64, filter *domain.TransactionMerchantFilter) (*domain.ReportMerchant, *domain.TransactionMerchantFilter, error)
		ReportOutlet(ctx context.Context, userID uint64, filter *domain.TransactionOutletFilter) (*domain.ReportOutlet, *domain.TransactionOutletFilter, error)
	}
)
