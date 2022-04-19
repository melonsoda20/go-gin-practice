package services

import (
	"context"
	"gin-todo-app/models"
	"time"

	"cloud.google.com/go/pubsub"
)

func GetPubSubConfig() (models.PubSubConfig, error) {
	configPath := "../config/pubsub.json"

	configFile, file_err := OpenFile(configPath)
	if file_err != nil {
		return models.PubSubConfig{}, file_err
	}
	defer configFile.Close()

	configData, data_err := ReadFile(configFile)
	if data_err != nil {
		return models.PubSubConfig{}, data_err
	}

	redisConfig := models.PubSubConfig{}
	json_error := DeserializeJSON(configData, &redisConfig)
	return redisConfig, json_error
}

func InitPubSubClient(ctx context.Context, conf models.PubSubConfig) (*pubsub.Client, error) {
	client, err := pubsub.NewClient(ctx, conf.ProjectID)
	return client, err
}

func PublishMessage(client *pubsub.Client, ctx context.Context, topicID string, msg string) (serverID string, err error) {
	t := client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})
	id, err := result.Get(ctx)
	return id, err
}

func ReceiveMessage(client *pubsub.Client, ctx context.Context, subID string) (string, error) {
	sub := client.Subscription(subID)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var data string
	err := sub.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
		data = string(msg.Data)
		msg.Ack()
	})
	return data, err
}
