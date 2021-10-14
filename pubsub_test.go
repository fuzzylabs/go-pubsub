package go_pubsub

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/go-cmp/cmp"
	"io/ioutil"
	"strings"
	"testing"

	"cloud.google.com/go/pubsub"
)

func TestSubmitResults(t *testing.T) {
	ctx := context.Background()

	mockTopicExists := &IPubSubTopicMock{
		ExistsFunc: func(ctx context.Context) (bool, error) {
			return true, nil
		},
		PublishFunc: func(ctx context.Context, msg *pubsub.Message) IPubSubPublishResult {
			return &IPubSubPublishResultMock{
				GetFunc: func(ctx context.Context) (serverID string, err error) {
					return "1234", nil
				},
			}
		},
	}

	mockTopicNotExists := &IPubSubTopicMock{
		ExistsFunc: func(ctx context.Context) (bool, error) {
			return false, nil
		},
		PublishFunc: func(ctx context.Context, msg *pubsub.Message) IPubSubPublishResult {
			panic("mock out the Publish method")
		},
	}

	mockPubSubExists := PubSub{
		ctx,
		&IPubSubClientMock{
			TopicFunc: func(id string) IPubSubTopic {
				return mockTopicExists
			},
		},
	}

	mockPubSubNotExists := PubSub{
		ctx,
		&IPubSubClientMock{
			TopicFunc: func(id string) IPubSubTopic {
				return mockTopicNotExists
			},
		},
	}

	topicID := "mock-topic"

	err := mockPubSubNotExists.PublishMessage(topicID, &empty.Empty{})
	if err == nil || err.Error() != fmt.Sprintf("Topic `%s` to does not exist", topicID) {
		t.Errorf("Expected submit to result in 'Topic to publish to does not exist' error")
	}

	err = mockPubSubExists.PublishMessage(topicID, &empty.Empty{})
	if err != nil {
		t.Errorf("Expected submit to finish without errors")
	}
}

func TestDecodeBody(t *testing.T) {
	testMessage := `{
    "message": {
        "attributes": {
            "key": "value"
        },
        "data": "SGVsbG8gQ2xvdWQgUHViL1N1YiEgSGVyZSBpcyBteSBtZXNzYWdlIQ==",
        "messageId": "2070443601311540",
        "message_id": "2070443601311540",
        "publishTime": "2021-02-26T19:13:55.749Z",
        "publish_time": "2021-02-26T19:13:55.749Z"
    },
   "subscription": "projects/myproject/subscriptions/mysubscription"
}
`
	api, _ := NewPubSub("")

	reader := strings.NewReader(testMessage)
	body := ioutil.NopCloser(reader)
	got, err := api.DecodeBody(body)
	if err != nil {
		t.Fatalf("Unexpected error while parsing a request body: %s", err.Error())
	}

	expected := []byte("Hello Cloud Pub/Sub! Here is my message!")
	if !cmp.Equal(expected, got) {
		t.Fatalf("Expected message: %s, but got: %s", string(expected), string(got))
	}

}
