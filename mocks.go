package go_pubsub

//go:generate moq -out pubsub_mock_test.go . IPubSubClient IPubSubTopic IPubSubPublishResult
//go:generate moq -out pubsub_mock.go . IPubSub
