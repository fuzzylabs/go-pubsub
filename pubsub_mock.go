// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package go_pubsub

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"io"
	"sync"
)

// Ensure, that IPubSubMock does implement IPubSub.
// If this is not the case, regenerate this file with moq.
var _ IPubSub = &IPubSubMock{}

// IPubSubMock is a mock implementation of IPubSub.
//
// 	func TestSomethingThatUsesIPubSub(t *testing.T) {
//
// 		// make and configure a mocked IPubSub
// 		mockedIPubSub := &IPubSubMock{
// 			DecodeDataFunc: func(body io.ReadCloser) ([]byte, error) {
// 				panic("mock out the DecodeData method")
// 			},
// 			DecodePushMessageFunc: func(body io.ReadCloser) (*PushMessage, error) {
// 				panic("mock out the DecodePushMessage method")
// 			},
// 			PublishMessageFunc: func(topicID string, submission protoreflect.ProtoMessage) error {
// 				panic("mock out the PublishMessage method")
// 			},
// 		}
//
// 		// use mockedIPubSub in code that requires IPubSub
// 		// and then make assertions.
//
// 	}
type IPubSubMock struct {
	// DecodeDataFunc mocks the DecodeData method.
	DecodeDataFunc func(body io.ReadCloser) ([]byte, error)

	// DecodePushMessageFunc mocks the DecodePushMessage method.
	DecodePushMessageFunc func(body io.ReadCloser) (*PushMessage, error)

	// PublishMessageFunc mocks the PublishMessage method.
	PublishMessageFunc func(topicID string, submission protoreflect.ProtoMessage) error

	// calls tracks calls to the methods.
	calls struct {
		// DecodeData holds details about calls to the DecodeData method.
		DecodeData []struct {
			// Body is the body argument value.
			Body io.ReadCloser
		}
		// DecodePushMessage holds details about calls to the DecodePushMessage method.
		DecodePushMessage []struct {
			// Body is the body argument value.
			Body io.ReadCloser
		}
		// PublishMessage holds details about calls to the PublishMessage method.
		PublishMessage []struct {
			// TopicID is the topicID argument value.
			TopicID string
			// Submission is the submission argument value.
			Submission protoreflect.ProtoMessage
		}
	}
	lockDecodeData        sync.RWMutex
	lockDecodePushMessage sync.RWMutex
	lockPublishMessage    sync.RWMutex
}

// DecodeData calls DecodeDataFunc.
func (mock *IPubSubMock) DecodeData(body io.ReadCloser) ([]byte, error) {
	if mock.DecodeDataFunc == nil {
		panic("IPubSubMock.DecodeDataFunc: method is nil but IPubSub.DecodeData was just called")
	}
	callInfo := struct {
		Body io.ReadCloser
	}{
		Body: body,
	}
	mock.lockDecodeData.Lock()
	mock.calls.DecodeData = append(mock.calls.DecodeData, callInfo)
	mock.lockDecodeData.Unlock()
	return mock.DecodeDataFunc(body)
}

// DecodeDataCalls gets all the calls that were made to DecodeData.
// Check the length with:
//     len(mockedIPubSub.DecodeDataCalls())
func (mock *IPubSubMock) DecodeDataCalls() []struct {
	Body io.ReadCloser
} {
	var calls []struct {
		Body io.ReadCloser
	}
	mock.lockDecodeData.RLock()
	calls = mock.calls.DecodeData
	mock.lockDecodeData.RUnlock()
	return calls
}

// DecodePushMessage calls DecodePushMessageFunc.
func (mock *IPubSubMock) DecodePushMessage(body io.ReadCloser) (*PushMessage, error) {
	if mock.DecodePushMessageFunc == nil {
		panic("IPubSubMock.DecodePushMessageFunc: method is nil but IPubSub.DecodePushMessage was just called")
	}
	callInfo := struct {
		Body io.ReadCloser
	}{
		Body: body,
	}
	mock.lockDecodePushMessage.Lock()
	mock.calls.DecodePushMessage = append(mock.calls.DecodePushMessage, callInfo)
	mock.lockDecodePushMessage.Unlock()
	return mock.DecodePushMessageFunc(body)
}

// DecodePushMessageCalls gets all the calls that were made to DecodePushMessage.
// Check the length with:
//     len(mockedIPubSub.DecodePushMessageCalls())
func (mock *IPubSubMock) DecodePushMessageCalls() []struct {
	Body io.ReadCloser
} {
	var calls []struct {
		Body io.ReadCloser
	}
	mock.lockDecodePushMessage.RLock()
	calls = mock.calls.DecodePushMessage
	mock.lockDecodePushMessage.RUnlock()
	return calls
}

// PublishMessage calls PublishMessageFunc.
func (mock *IPubSubMock) PublishMessage(topicID string, submission protoreflect.ProtoMessage) error {
	if mock.PublishMessageFunc == nil {
		panic("IPubSubMock.PublishMessageFunc: method is nil but IPubSub.PublishMessage was just called")
	}
	callInfo := struct {
		TopicID    string
		Submission protoreflect.ProtoMessage
	}{
		TopicID:    topicID,
		Submission: submission,
	}
	mock.lockPublishMessage.Lock()
	mock.calls.PublishMessage = append(mock.calls.PublishMessage, callInfo)
	mock.lockPublishMessage.Unlock()
	return mock.PublishMessageFunc(topicID, submission)
}

// PublishMessageCalls gets all the calls that were made to PublishMessage.
// Check the length with:
//     len(mockedIPubSub.PublishMessageCalls())
func (mock *IPubSubMock) PublishMessageCalls() []struct {
	TopicID    string
	Submission protoreflect.ProtoMessage
} {
	var calls []struct {
		TopicID    string
		Submission protoreflect.ProtoMessage
	}
	mock.lockPublishMessage.RLock()
	calls = mock.calls.PublishMessage
	mock.lockPublishMessage.RUnlock()
	return calls
}
