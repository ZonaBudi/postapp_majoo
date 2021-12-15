package port

import (
	"context"
	"postapp/internal/domain"
)

type (
	UserRepository interface {
		FindOneByUsername(ctx context.Context, username string) (*domain.User, error)
	}
	TransactionRepository interface {
		TransactionByMerchant(ctx context.Context, filter *domain.TransactionMerchantFilter) ([]*domain.TransactionMerchant, *domain.TransactionMerchantFilter, error)
		TransactionByOutlet(ctx context.Context, filter *domain.TransactionOutletFilter) ([]*domain.TransactionOutlet, *domain.TransactionOutletFilter, error)
	}
	MerchantRepository interface {
		FindOneByUserID(ctx context.Context, merchantID uint64) (*domain.Merchant, error)
	}

	OutletRepository interface {
		FindOneByOutletID(ctx context.Context, outletID uint64) (*domain.Outlet, error)
	}
)
