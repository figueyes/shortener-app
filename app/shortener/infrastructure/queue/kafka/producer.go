package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/figueyes/shortener-app/app/shortener/infrastructure/queue/kafka/config"
	"reflect"

	"github.com/figueyes/shortener-app/app/shortener/infrastructure/utils"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type Queue interface {
	Publish(topic string, data interface{}) error
}

type producer struct {
	config config.WriterConfig
}

func newPublisher(config config.WriterConfig) *producer {
	return &producer{
		config: config,
	}
}

func GetPublisherFactory(config config.WriterConfig) *producer {
	return newPublisher(config)
}

func (p *producer) Publish(topic string, data interface{}) error {
	ctx := context.Background()
	kafkaMessages, err := p.createKafkaMessages(data)
	if err != nil {
		return errors.New("error on publishing kafka message")
	}

	if p.config.Compression == 0 { // 0 = none
		p.config.Compression = config.Snappy
	}

	kafkaWriter := &kafka.Writer{
		Addr:        kafka.TCP(p.config.Brokers...),
		Topic:       topic,
		Balancer:    &kafka.LeastBytes{},
		Compression: p.config.Compression,
		Logger:      p.config.Logger,
		ErrorLogger: p.config.ErrorLogger,
	}
	defer kafkaWriter.Close()
	err = kafkaWriter.WriteMessages(ctx, kafkaMessages...)
	if err != nil {
		fmt.Printf("error writing kafka message with: %s", err.Error())
		return err
	} else {
		fmt.Printf("message has been published successfully with\ntopic: %s,\nmessage: %s,\nsize: %d bytes\n",
			topic,
			utils.EntityToJson(data),
			uint(kafkaWriter.Stats().Bytes),
		)
	}
	return nil
}

func (p *producer) createKafkaMessages(data interface{}) ([]kafka.Message, error) {
	var kafkaMessages []kafka.Message

	if utils.IsNilFixed(data) {
		return nil, errors.New("kafka error the data is empty")
	}

	switch reflect.TypeOf(data).Kind() {
	case reflect.Array, reflect.Slice:
		value := reflect.ValueOf(data)
		for index := 0; index < value.Len(); index++ {
			kafkaMessages = append(kafkaMessages, p.messageBuilder(value.Index(index)))
		}

	default:
		kafkaMessages = append(kafkaMessages, p.messageBuilder(data))
	}
	return kafkaMessages, nil
}

func (p *producer) messageBuilder(data interface{}) kafka.Message {
	payload, err := json.Marshal(data)
	if err != nil {
		payload = []byte("{}")
	}
	key := uuid.New().String()
	kafkaMessage := kafka.Message{
		Key:   []byte(key),
		Value: payload,
	}
	return kafkaMessage
}
