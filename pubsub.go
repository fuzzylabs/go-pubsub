package go_pubsub

import (
	"cloud.google.com/go/pubsub"
	pubsubapi "cloud.google.com/go/pubsub"
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"io"
)

// IPubSubPublishResult interface for results of publishing a message to Pub/Sub topic
type IPubSubPublishResult interface {
	Get(ctx context.Context) (serverID string, err error)
}

// IPubSubTopic interface for a Pub/Sub topic (either existing or non-existing) for publishing
type IPubSubTopic interface {
	Exists(ctx context.Context) (bool, error)
	Publish(ctx context.Context, msg *pubsubapi.Message) IPubSubPublishResult
}

// PubSubTopic implementation of a Pub/Sub topic for publishing
type PubSubTopic struct {
	topic *pubsubapi.Topic
}

// Exists checks if topic exists on PubSub
func (t PubSubTopic) Exists(ctx context.Context) (bool, error) {
	return t.topic.Exists(ctx)
}

// Publish publishes a message to the topic
func (t PubSubTopic) Publish(ctx context.Context, msg *pubsubapi.Message) IPubSubPublishResult {
	return t.topic.Publish(ctx, msg)
}

// IPubSubClient interface for a Pub/Sub clienbt
type IPubSubClient interface {
	Topic(id string) IPubSubTopic
}

// PubSubClient implementation of a Pub/Sub client for publishing
type PubSubClient struct {
	client *pubsubapi.Client
}

// Topic returns a Pub/Sub topic for a given ID
func (p PubSubClient) Topic(id string) IPubSubTopic {
	return PubSubTopic{p.client.Topic(id)}
}

type IPubSub interface {
	SubmitResults(topicID string, submission *proto.Message) error
	DecodeBody(body io.ReadCloser) ([]byte, error)
}

// PubSub a struct that holds a Pub/Sub client for publishing
type PubSub struct {
	ctx    context.Context
	client IPubSubClient
}

// NewPubSub returns a PubSub struct with a Pub/Sub client for a given Google Cloud project ID
func NewPubSub(projectID string) (*PubSub, error) {
	ctx := context.Background()

	client, err := pubsubapi.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}

	return &PubSub{ctx, PubSubClient{client}}, nil
}

// PublishMessage publishes the message to the Pub/Sub topic
func (p PubSub) PublishMessage(topicID string, message proto.Message) error {
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	t := p.client.Topic(topicID)

	exists, err := t.Exists(p.ctx)

	if !exists {
		return errors.New(fmt.Sprintf("Topic `%s` to does not exist", topicID))
	}

	if err != nil {
		return err
	}

	result := t.Publish(p.ctx, &pubsub.Message{
		Data: messageBytes,
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(p.ctx)
	if err != nil {
		return err
	}
	log.Infof("Published a message; msg ID: %v\n", id)
	return nil
}

type Message struct {
	Attributes map[string]interface{}
	Data       string
}

type PushMessage struct {
	Message Message
}

// DecodeBody decodes the body of a push subscription request
func (p PubSub) DecodeBody(body io.ReadCloser) ([]byte, error) {
	var pr PushMessage
	if err := json.NewDecoder(body).Decode(&pr); err != nil {
		return nil, err
	}

	log.Infof("Encoded push message received: %#v", pr)

	data, err := b64.StdEncoding.DecodeString(pr.Message.Data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
