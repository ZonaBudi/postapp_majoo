package app

import (
	"net/http"
	"postapp/internal/adapter/indbound/authhttp"
	"postapp/internal/adapter/indbound/transactionhttp"
	"postapp/internal/adapter/outbound/merchanrepositorymysql"
	"postapp/internal/adapter/outbound/outletrepositorymysql"
	"postapp/internal/adapter/outbound/transactionrepositorymysql"
	"postapp/internal/adapter/outbound/userrepositorymysql"
	"postapp/internal/service/authservice"
	"postapp/internal/service/reportservice"

	"postapp/cmd/docs"
	_ "postapp/cmd/docs"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Handlers struct {
	Mysql  *gorm.DB
	R      *chi.Mux
	Logger *zap.Logger
}

/*Service
*Initialize the service with the router, the middleware and the dependencies
 */
func (h *Handlers) Service() http.Handler {

	//repository get data from other entities
	userRepoMysql := userrepositorymysql.NewUserRepMysql(h.Mysql)
	transactionRepMysql := transactionrepositorymysql.NewTransactionRepMysql(h.Mysql)
	merchantRepoMysql := merchanrepositorymysql.NewMerchantRepMysql(h.Mysql)
	outletRepoMysql := outletrepositorymysql.NewOutletRepMysql(h.Mysql)

	//service  usecase and bussiness logic
	authService := authservice.NewAuthService(h.Logger, userRepoMysql)
	reportService := reportservice.NewReportService(h.Logger, transactionRepMysql, merchantRepoMysql, outletRepoMysql)
	// reportService := reportservice.NewReportService(h.Logger, transactionRepoMysql)

	//router the endpoints http
	authhttp.NewAuthHttp(h.R, authService)
	transactionhttp.NewReportHttp(h.R, reportService)

	//health check
	h.R.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	//swagger
	docs.SwaggerInfo.Host = viper.GetString("server.host") + ":" + viper.GetString("server.port")
	h.R.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(viper.GetString("server.host")+":"+viper.GetString("server.port")+"/swagger/doc.json"), //The url pointing to API definition
	))
	return h.R
}
