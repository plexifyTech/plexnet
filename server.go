package plexnet

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func RunServer(ctx context.Context, engine *gin.Engine) {
	server := &http.Server{
		Addr:    ":" + config.Port,
		Handler: engine,
	}

	defer shutdown(server, ctx)

	exitSig := make(chan os.Signal, 1)
	signal.Notify(exitSig,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go bootServer(server)

	log.Info().Msgf("running app %s on port %s", config.AppName, config.Port)

	stop := <-exitSig
	log.Info().Msgf("received shutdown signal %v", stop)
}

func DefaultEngine() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(GinZeroLogger())
	engine.Use(gin.Recovery())
	engine.Use(cors.Default())

	return engine
}

func bootServer(server *http.Server) {
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal().Msg("application boot failed")
	}
}

func shutdown(server *http.Server, ctx context.Context) {
	if err := server.Shutdown(ctx); err != nil {
		log.Err(err).Msg("shutting down")
	}
}
