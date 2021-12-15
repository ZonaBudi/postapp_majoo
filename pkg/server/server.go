package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"postapp/internal/app"
	"postapp/pkg/config"
	"postapp/pkg/logger"
	"postapp/pkg/mysql"
	"postapp/pkg/request"
	"runtime"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/spf13/viper"
	l "github.com/treastech/logger"
)

/*Run
*Setup Run Server
 */
func Run() {

	//Load Config
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Join(filepath.Dir(b), "../..")
	config := &config.EnvConfig{
		FileName: "config",
		Path:     basepath,
	}
	if err := config.ReadConfig(); err != nil {
		panic(err)
	}

	//Connect Mysql
	mysql, err := mysql.Connect()
	if err != nil {
		panic(err)
	}
	sqlDB, err := mysql.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	//zap logger init
	zap, err := logger.Initialize()
	if err != nil {
		panic(err)
	}
	defer zap.Sync()

	//router framework
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(l.Logger(zap))
	r.Use(request.Recoverer)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	ah := &app.Handlers{
		Mysql:  mysql,
		R:      r,
		Logger: zap,
	}

	// The HTTP Server
	server := &http.Server{Addr: viper.GetString("server.host") + ":" + viper.GetString("server.port"), Handler: ah.Service()}

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				panic("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			panic(err)
		}
		serverStopCtx()
	}()

	// Run the server
	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(err)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
}
