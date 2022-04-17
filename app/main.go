package main

import (
	useCases "github.com/figueyes/shortener-app/app/shortener/application/use-cases"
	"github.com/figueyes/shortener-app/app/shortener/infrastructure/controllers"
	"github.com/figueyes/shortener-app/app/shortener/infrastructure/log"
	"github.com/figueyes/shortener-app/app/shortener/infrastructure/persistence/mongo/repository"
	redisRepository "github.com/figueyes/shortener-app/app/shortener/infrastructure/persistence/redis/repository"
	"github.com/figueyes/shortener-app/app/shortener/infrastructure/queue/kafka"
	"github.com/figueyes/shortener-app/app/shortener/infrastructure/queue/kafka/config"
	"github.com/figueyes/shortener-app/app/version"
	"github.com/labstack/echo/v4"
	"strings"

	"fmt"
	"net/http"
	"os"
	"time"
)

var (
	_topic  = os.Getenv("SHORTENER_TOPIC")
	_broker = os.Getenv("SHORTENER_KAFKA_BROKER")
	_user   = os.Getenv("SHORTENER_KAFKA_USER")
	_pass   = os.Getenv("SHORTENER_KAFKA_PASS")
)

const (
	_versionApp = "1.0.0"
	_author     = "Alex Figueroa"
	_appName    = "api-shortener"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	// mongo config
	shortenerMongoRepo := repository.MongoRepositoryFactory("shorts")

	// redisConfig config
	shortenerRedisRepository := redisRepository.RedisRepositoryFactory()
	// subscribe to channel: _removed
	go shortenerRedisRepository.DeleteFromCache(_topic)

	// mapper config
	// mapperRepository := mapper.NewMapperRepository()

	// usecases
	shortUrlUseCase := useCases.NewShortUrlUseCase(shortenerMongoRepo)
	getUrlUseCase := useCases.NewGetUrlUseCase(shortenerMongoRepo)
	modifyShortUseCase := useCases.NewModifyShortUseCase(shortenerMongoRepo)

	// kafka
	producer := kafka.GetPublisherFactory(config.WriterConfig{
		Connection: config.Connection{
			Username: _user,
			Password: _pass,
		},
		Brokers:     strings.Split(_broker, ","),
		Compression: config.Snappy,
	})

	// controller
	c := controllers.NewShortenerHandler(e, shortUrlUseCase, getUrlUseCase, modifyShortUseCase, shortenerRedisRepository, producer)
	go c.Publish()
	version.NewHealthHandler(e, _versionApp, _appName, _author)

	log.Info("Starting server")
	portServer := os.Getenv("PORT_SERVER")
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", portServer),
		ReadTimeout:  3 * time.Minute,
		WriteTimeout: 3 * time.Minute,
	}
	e.Logger.Fatal(e.StartServer(server))

}
