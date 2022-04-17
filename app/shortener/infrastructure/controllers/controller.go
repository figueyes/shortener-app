package controllers

import (
	"fmt"
	useCases "github.com/figueyes/shortener-app/app/shortener/application/use-cases/interfaces"
	"github.com/figueyes/shortener-app/app/shortener/domain/entities"
	"github.com/figueyes/shortener-app/app/shortener/infrastructure/controllers/model"
	"github.com/figueyes/shortener-app/app/shortener/infrastructure/log"
	"github.com/figueyes/shortener-app/app/shortener/infrastructure/persistence/cache"
	"github.com/figueyes/shortener-app/app/shortener/infrastructure/queue/kafka"
	"github.com/figueyes/shortener-app/app/shortener/infrastructure/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"time"
)

var (
	_topic_removed  = os.Getenv("SHORTENER_REDIS_TOPIC_REMOVED")
	_topic_observed = os.Getenv("SHORTENER_KAFKA_TOPIC_OBSERVED")
)

type shortenerHandler struct {
	create  useCases.Short
	get     useCases.Get
	modify  useCases.Modify
	cache   cache.Cache
	queue   kafka.Queue
	channel chan *model.QueueShort
}

func NewShortenerHandler(e *echo.Echo,
	create useCases.Short,
	get useCases.Get,
	modify useCases.Modify,
	cache cache.Cache,
	queue kafka.Queue,
) *shortenerHandler {
	sh := &shortenerHandler{
		create:  create,
		get:     get,
		modify:  modify,
		cache:   cache,
		queue:   queue,
		channel: make(chan *model.QueueShort, 1),
	}
	base := e.Group(fmt.Sprintf("%s/shortener", os.Getenv("BASE_PATH")))
	{
		base.POST("/", sh.createUrl)
		base.GET("/:shortUrl", sh.getUrl)
		base.PATCH("/:shortUrl", sh.modifyShort)
	}
	return sh
}

func (sh *shortenerHandler) createUrl(c echo.Context) error {
	request := new(model.CreateInput)
	err := c.Bind(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			echo.Map{
				"message": "invalid payload",
			},
		)
	}
	err = request.ValidateUrl()
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	e := new(entities.Short)
	e.OriginalUrl = request.Url
	e.User = request.User
	shorted, err := sh.create.Short(e)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			echo.Map{
				"message": err.Error(),
			},
		)
	}
	shortUrl := &model.CreateOutput{
		ShortUrl: shorted.ShortUrl,
	}
	return c.JSON(200, shortUrl)
}

func (sh *shortenerHandler) getUrl(c echo.Context) error {
	var status int
	msg := new(model.QueueShort)
	msg.RequestCreatedAt = time.Now()
	msg.MethodRunner = "GET_URL"
	response := new(model.GetOutput)
	shortUrl := c.Param("shortUrl")
	if len(shortUrl) == 0 {
		status = http.StatusBadRequest
		return c.JSON(status, echo.Map{
			"message": "no param in request",
		})
	}

	cached, _ := sh.cache.Get(shortUrl)
	if cached != nil {
		utils.JsonToEntity(*cached, &response)
		status = http.StatusMovedPermanently
		msg.StatusHttp = status
		msg.ShortUrl = shortUrl
		sh.channel <- msg
		return c.JSON(status, response)
	}

	shorted, err := sh.get.Get(shortUrl)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	if shorted == nil {
		status = http.StatusNotFound
		msg.StatusHttp = status
		msg.ShortUrl = shortUrl
		sh.channel <- msg
		if err != nil {
			log.Warn("cannot publish message... continue")
		}
		return c.JSON(status, echo.Map{
			"message": "shorted url not found",
		})
	}
	response.Url = shorted.OriginalUrl
	err = sh.cache.Set(shortUrl, utils.EntityToJson(response))
	if err != nil {
		log.Warn("error trying to save into cache: %s", err.Error())
	}
	status = http.StatusMovedPermanently
	msg.StatusHttp = status
	msg.ShortUrl = shortUrl
	sh.channel <- msg
	return c.JSON(status, response)
}

func (sh *shortenerHandler) modifyShort(c echo.Context) error {
	shortUrl := c.Param("shortUrl")
	if len(shortUrl) == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "no param in request",
		})
	}
	modifyInput := new(model.ModifyInput)
	err := c.Bind(modifyInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			echo.Map{
				"message": "invalid payload",
			},
		)
	}
	toUpdate := new(entities.Short)
	toUpdate.OriginalUrl = modifyInput.OriginalUrl
	toUpdate.User = modifyInput.User
	toUpdate.IsEnable = modifyInput.IsEnable
	updated, err := sh.modify.Modify(shortUrl, toUpdate)
	if updated != nil {
		err = sh.cache.Publish(_topic_removed, updated)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
		}
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	if updated == nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "shorted url not found",
		})
	}
	return c.JSON(http.StatusOK, updated)
}

func (sh *shortenerHandler) Publish() {
	for {
		msg := <-sh.channel
		err := sh.queue.Publish(_topic_observed, msg)
		if err != nil {
			log.Warn("cannot publish message: %s\n... continue", err.Error())
		}
	}
}
