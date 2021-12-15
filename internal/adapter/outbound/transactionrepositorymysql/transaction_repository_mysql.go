package transactionrepositorymysql

import (
	"context"
	"fmt"
	"postapp/internal/domain"
	"postapp/internal/port"
	"postapp/pkg/model"
	"postapp/pkg/paginator"

	extraClausePlugin "github.com/WinterYukky/gorm-extra-clause-plugin"
	"github.com/WinterYukky/gorm-extra-clause-plugin/exclause"
	"gorm.io/gorm"
)

type transactionRepoMysql struct {
	mysql *gorm.DB
}

func NewTransactionRepMysql(mysql *gorm.DB) port.TransactionRepository {
	return &transactionRepoMysql{
		mysql: mysql,
	}
}

func (_instance *transactionRepoMysql) TransactionByMerchant(ctx context.Context, filter *domain.TransactionMerchantFilter) ([]*domain.TransactionMerchant, *domain.TransactionMerchantFilter, error) {
	var (
		result []*domain.TransactionMerchant
		ent    *domain.Transaction
		count  int64
	)
	StartDate := filter.StartDate.Format(model.DateFormatFilter)
	EndDate := filter.EndDate.Format(model.DateFormatFilter)
	queryRangeDate := fmt.Sprintf(`select '%s' dt union all select dt + interval 1 day from all_dates where dt + interval 1 day <= '%s'`, StartDate, EndDate)
	_instance.mysql.Use(extraClausePlugin.New())
	q := _instance.mysql.Model(&ent).
		Clauses(exclause.With{Recursive: true, CTEs: []exclause.CTE{{Name: "all_dates", Subquery: exclause.Subquery{DB: _instance.mysql.Raw(queryRangeDate)}}}}).
		Table("all_dates").
		Joins(`left join Transactions t on DATE_FORMAT(t.updated_at , "%Y-%m-%d")  = all_dates.dt`).
		Select(`coalesce(SUM(t.bill_total), 0) as summary, all_dates.dt as transaction_date`).
		Where("t.merchant_id = ?", filter.MerchantID).
		Group(`all_dates.dt`)
	counter := q.Count(&count)
	if counter.Error != nil {
		return nil, nil, counter.Error
	}
	filter.Pagination.Count = count
	filter.Pagination = paginator.Paging(filter.Pagination)
	err := q.Offset(filter.Pagination.Offset).Limit(filter.Pagination.Limit).Find(&result).Error
	if err != nil {
		return nil, nil, err
	}
	return result, filter, nil
}

func (_instance *transactionRepoMysql) TransactionByOutlet(ctx context.Context, filter *domain.TransactionOutletFilter) ([]*domain.TransactionOutlet, *domain.TransactionOutletFilter, error) {
	var (
		result []*domain.TransactionOutlet
		ent    *domain.Transaction
		count  int64
	)
	StartDate := filter.StartDate.Format(model.DateFormatFilter)
	EndDate := filter.EndDate.Format(model.DateFormatFilter)
	queryRangeDate := fmt.Sprintf(`select '%s' dt union all select dt + interval 1 day from all_dates where dt + interval 1 day <= '%s'`, StartDate, EndDate)
	_instance.mysql.Use(extraClausePlugin.New())
	q := _instance.mysql.Model(&ent).
		Clauses(exclause.With{Recursive: true, CTEs: []exclause.CTE{{Name: "all_dates", Subquery: exclause.Subquery{DB: _instance.mysql.Raw(queryRangeDate)}}}}).
		Table("all_dates").
		Joins(`left join Transactions t on DATE_FORMAT(t.updated_at , "%Y-%m-%d")  = all_dates.dt`).
		Select(`coalesce(SUM(t.bill_total), 0) as summary, all_dates.dt as transaction_date`).
		Where("t.outlet_id = ?", filter.OutletID).
		Group(`all_dates.dt`)
	counter := q.Count(&count)
	if counter.Error != nil {
		return nil, nil, counter.Error
	}
	filter.Pagination.Count = count
	filter.Pagination = paginator.Paging(filter.Pagination)
	err := q.Offset(filter.Pagination.Offset).Limit(filter.Pagination.Limit).Find(&result).Error
	if err != nil {
		return nil, nil, err
	}
	return result, filter, nil
}
