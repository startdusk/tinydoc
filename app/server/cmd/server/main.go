package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/namsral/flag"

	docDelivery "server/internal/doc/delivery/http"
	docRepo "server/internal/doc/repository/mysql"
	docUsecase "server/internal/doc/usecase"
)

var dbhost = flag.String("MYSQL_HOST", "localhost", "mysql host")
var dbport = flag.Int("MYSQL_PORT", 3306, "mysql port")
var dbuser = flag.String("MYSQL_USER", "", "mysql user")
var dbpass = flag.String("MYSQL_PASSWORD", "", "mysql password")
var dbname = flag.String("MYSQL_DATABASE_NAME", "", "mysql database name")
var ctxTimeout = flag.Int64("CONTEXT_TIMEOUT", 5, "context timeout")
var webPort = flag.Int("WEB_PORT", 8080, "web port")

func main() {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", *dbuser, *dbpass, *dbhost, *dbport, *dbname)
	var val url.Values
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Shanhai")
	dsn := fmt.Sprintf("%s?%s", conn, val.Encode())
	db, err := sql.Open(`mysql`, dsn)
	if err != nil {
		log.Panic(err)
	}
	err = db.Ping()
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := gin.Default()

	contextTimeout := time.Duration(*ctxTimeout) * time.Second
	repo := docRepo.NewDocRepository(db)
	usecase := docUsecase.NewDocUsecase(repo, contextTimeout)
	docDelivery.NewDocHandler(router, usecase)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", *webPort),
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
