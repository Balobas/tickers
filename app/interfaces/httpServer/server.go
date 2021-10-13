package httpServer

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tickers/app/data"
	"tickers/app/interfaces/database"
	"tickers/app/interfaces/httpServer/converter"
	"tickers/app/repo"
)

type Config struct {
	Port string
	Db   database.Database
}

func Run(config Config) {

	router := gin.New()

	router.GET("/", GetAllTickers(repo.NewTickerRepository(config.Db)))
	router.GET("/favicon.ico", func(context *gin.Context) {
		context.JSONP(http.StatusOK, "")
	})

	err := router.Run(":" + config.Port)
	if err != nil {
		log.Fatal(err)
	}
}

func GetAllTickers(tickersRepo data.TickerRepository) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		tickers, err := tickersRepo.GetAll()
		if err != nil {
			ctx.Writer.WriteHeader(http.StatusBadGateway)
			log.Println(err)
			return
		}

		ctx.Writer.WriteHeader(http.StatusOK)

		ctx.JSONP(http.StatusOK, converter.ConvertTickers(tickers))
	}
}
