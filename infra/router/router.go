package router

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/hidenari-yuda/direct-ai-go/domain/config"
	"github.com/hidenari-yuda/direct-ai-go/infra/database"
	"github.com/hidenari-yuda/direct-ai-go/infra/di"
	"github.com/hidenari-yuda/direct-ai-go/infra/driver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	// "github.com/improbable-eng/grpc-web/go/grpcweb"
)

func Start() {
	var (
		db       = database.NewDb()
		firebase = driver.NewFirebaseImpl()
	)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.App.Port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	ctx := context.Background()
	di.RegisterServiceServer(ctx, s, db, firebase)

	// for using grpcurl
	reflection.Register(s)

	go func() {
		log.Printf("start gRPC server, port: %d", config.App.Port)
		err = s.Serve(listener)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()

}

// http 
// // HTTP経由で接続できるようラップする
// wrappedServer := grpcweb.WrapServer(
// 	s,
// 	// CORSの設定
// 	grpcweb.WithOriginFunc(func(origin string) bool {
// 		return origin == "http://localhost:3000"
// 	}),
// )
// mux := http.NewServeMux()
// mux.Handle("/", http.HandlerFunc(wrappedServer.ServeHTTP))
// // ポート8080で起動
// hs := &http.Server{
// 	Addr:    ":8080",
// 	Handler: mux,
// }
// log.Fatal(hs.ListenAndServe())

// r.Engine.HidePort = true
// r.Engine.HideBanner = true
// r.Engine.Use(middleware.Recover())

// var origins []string

// if r.cfg.App.Env == "local" {
// 	origins = []string{
// 		"http://localhost:8080",
// 		"http://localhost:3000",
// 	}
// } else if r.cfg.App.Env == "dev" {
// 	origins = r.cfg.App.CorsDomains
// } else if r.cfg.App.Env == "prd" {
// 	origins = r.cfg.App.CorsDomains
// }

// r.Engine.Use(middleware.CORSWithConfig((middleware.CORSConfig{
// 	AllowOrigins: origins,
// 	AllowHeaders: []string{
// 		echo.HeaderAuthorization,
// 		echo.HeaderAccessControlAllowHeaders,
// 		echo.HeaderContentType,
// 		echo.HeaderOrigin,
// 		echo.HeaderAccessControlAllowOrigin,
// 		"FirebaseAuthorization",
// 	},
// 	AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
// })))

// r.Engine.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
// 	Skipper: func(c echo.Context) bool {
// 		if strings.Contains(c.Request().URL.Path, "healthz") {
// 			return true
// 		} else {
// 			return false
// 		}
// 	},
// }))

// func (r *Router) Start() {
// 	fmt.Printf("start server, port: %d", config.App.Port)
// 	r.Engine.Start(fmt.Sprintf(":%d", config.App.Port))
// }

// callInfo contains all related configuration and information about an RPC.
// type callInfo struct {
// 	compressorType        string
// 	failFast              bool
// 	maxReceiveMessageSize *int
// 	maxSendMessageSize    *int
// 	creds                 credentials.PerRPCCredentials
// 	contentSubtype        string
// 	codec                 baseCodec
// 	maxRetryRPCBufferSize int
// }

// CallOption configures a Call before it starts or extracts information from
// a Call after it completes.
// type CallOption interface {
// 	// before is called before the call is sent to any server.  If before
// 	// returns a non-nil error, the RPC fails with that error.
// 	before(*callInfo) error

// 	// after is called after the call has completed.  after cannot return an
// 	// error, so any failures should be reported via output parameters.
// 	after(*callInfo, *csAttempt)
// }

// var (
// db       = database.NewDB(r.cfg.Db, true)
// firebase = driver.NewFirebaseImpl(r.cfg.Firebase)
// basicAuth = utils.NewBasicAuth(r.cfg)
// )

// JSON REST API
// r.Engine.HidePort = true
// r.Engine.HideBanner = true
// r.Engine.Use(middleware.Recover())
// // TODO: Webクライアントのドメインが決まったら設定する 👆の`r.Engine.Use(middleware.CORS())`は消す
// // r.Engine.Use(middleware.CORSWithConfig((middleware.CORSConfig{
// // AllowOrigins: r.cfg.App.CorsDomains,
// // 	AllowHeaders: []string{echo.HeaderAuthorization, echo.HeaderContentType, echo.HeaderOrigin, echo.HeaderAccessControlAllowOrigin},
// // 	AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
// // })))
// r.Engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
// 	AllowOrigins: []string{"*"},
// 	AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
// }))
// r.Engine.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
// 	Skipper: func(c echo.Context) bool {
// 		if strings.Contains(c.Request().URL.Path, "healthz") {
// 			return true
// 		} else {
// 			return false
// 		}
// 	},
// }))
// r.Engine.HidePort = true
// r.Engine.HideBanner = true
// r.Engine.Use(middleware.Recover())

// var origins []string

// if r.cfg.App.Env == "local" {
// 	origins = []string{
// 		"http://localhost:9090",
// 		"http://localhost:3000",
// 		"http://localhost:3001",
// 		"http://localhost:3002",
// 	}
// } else if r.cfg.App.Env == "dev" {
// 	origins = r.cfg.App.CorsDomains
// } else if r.cfg.App.Env == "prd" {
// 	origins = r.cfg.App.CorsDomains
// }

// fmt.Println("------------")
// fmt.Println(r.cfg.App.Env)
// fmt.Println(origins)
// fmt.Println("------------")

// r.Engine.Use(middleware.CORSWithConfig((middleware.CORSConfig{
// 	AllowOrigins: origins,
// 	AllowHeaders: []string{
// 		echo.HeaderAuthorization,
// 		echo.HeaderAccessControlAllowHeaders,
// 		echo.HeaderContentType,
// 		echo.HeaderOrigin,
// 		echo.HeaderAccessControlAllowOrigin,
// 		"FirebaseAuthorization",
// 	},
// 	AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
// })))

// r.Engine.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
// 	Skipper: func(c echo.Context) bool {
// 		if strings.Contains(c.Request().URL.Path, "healthz") {
// 			return true
// 		} else {
// 			return false
// 		}
// 	},
// }))

// api := r.Engine.Group("")
// {
// 	api.GET("/healthz", func(c echo.Context) error {
// 		return c.NoContent(http.StatusOK)
// 	})

// 	api.GET("/*", func(c echo.Context) error {
// 		return c.NoContent(http.StatusNotFound)
// 	})

// 	api.POST("/*", func(c echo.Context) error {
// 		return c.NoContent(http.StatusNotFound)
// 	})

// 	api.PUT("/*", func(c echo.Context) error {
// 		return c.NoContent(http.StatusNotFound)
// 	})
// }

// /****************************************************************************************/
// /// No Auth API
// //

// noAuthAPI := api.Group("api")
// {
// 	noAuthAPI.GET("/healthz", func(c echo.Context) error {
// 		return c.NoContent(http.StatusOK)
// 	})

// 	// ユーザーの新規登録
// 	noAuthAPI.POST("/signup", userRoutes.SignUp(db, firebase))

// 	// ユーザーのログイン
// 	noAuthAPI.PUT("/signin", userRoutes.SignIn(db, firebase))

// }

// /****************************************************************************************/
// /// UserAPI
// //
// // userAPI := noAuthAPI.Group("/user")
// {
// 	// ユーザーのログイン

// }
