package transactionhttp

import (
	"net/http"
	"postapp/internal/domain"
	"postapp/internal/midleware"
	"postapp/internal/port"
	"postapp/pkg/paginator"
	"postapp/pkg/request"
	"postapp/pkg/response"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type transactionHandlerHttp struct {
	app           *chi.Mux
	reportService port.ReportService
}

func NewReportHttp(app *chi.Mux, reportService port.ReportService) {
	api := transactionHandlerHttp{
		app:           app,
		reportService: reportService,
	}
	api.app.Group(func(r chi.Router) {
		r.Use(midleware.Authenticate())
		r.Use(request.RequestMiddleware)
		r.Get("/report-merchant", api.GetReportMerchant)
		r.Get("/report-outlet/{outletID:[0-9]}", api.GetReportOutlet)
	})
}

// Report Merchant godoc
// @Summary      List Transaction Merchant
// @Description  get list transaction merchant
// @Security 	 Bearer
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param        page   query     int  false  "page filter"  	   minimum(1)
// @Param        limit  query     int  false  "limit data filter"  minimum(5)
// @Success      200  {object} response.AppSuccess
// @Failure      default {object}  response.AppError
// @Router       /report-merchant [get]
func (_instance *transactionHandlerHttp) GetReportMerchant(w http.ResponseWriter, r *http.Request) {
	paginate, err := paginator.Paginate(r)
	if err != nil {
		response.Error(w, response.NewError(http.StatusBadRequest, response.WithMessageError(response.ErrBadRequest)))
		return
	}
	query := &domain.TransactionMerchantFilter{
		Pagination: paginate,
	}
	user := midleware.User(r.Context())
	result, meta, err := _instance.reportService.ReportMerchant(r.Context(), *user.ID, query)
	if err != nil {
		response.Error(w, err)
		return
	}
	response.Success(w, http.StatusOK, response.SuccessData(result), response.SuccessMeta(meta.Pagination))
}

// Report Outlet godoc
// @Summary      List Transaction Outlet
// @Description  get list transaction Outlet
// @Tags         transactions
// @Security	 Bearer
// @Accept       json
// @Produce      json
// @Param        page   	query     int  false  "page filter"  	   minimum(1)
// @Param        limit  	query     int  false  "limit data filter"  minimum(5)
// @Param        outlet_id  path      int  true   "Outlet ID"          minimum(1)
// @Success      200 {object} response.AppSuccess
// @Failure      default {object}  response.AppError
// @Router       /report-outlet [get]
func (_instance *transactionHandlerHttp) GetReportOutlet(w http.ResponseWriter, r *http.Request) {
	paginate, err := paginator.Paginate(r)
	if err != nil {
		response.Error(w, response.NewError(http.StatusBadRequest, response.WithMessageError(response.ErrBadRequest)))
		return
	}
	outletID, err := strconv.ParseUint(chi.URLParam(r, "outletID"), 10, 64)
	if err != nil {
		response.Error(w, response.NewError(http.StatusBadRequest, response.WithMessageError(response.ErrBadRequest)))
	}
	query := &domain.TransactionOutletFilter{
		Pagination: paginate,
		OutletID:   outletID,
	}
	user := midleware.User(r.Context())
	result, meta, err := _instance.reportService.ReportOutlet(r.Context(), *user.ID, query)
	if err != nil {
		response.Error(w, err)
		return
	}
	response.Success(w, http.StatusOK, response.SuccessData(result), response.SuccessMeta(meta.Pagination))
}
