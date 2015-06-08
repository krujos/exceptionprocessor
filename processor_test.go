package exceptionprocessor_test

import (
	"time"

	"github.com/cloudfoundry/noaa/events"
	"github.com/gogo/protobuf/proto"
	. "github.com/krujos/exceptionprocessor"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//TODO These functions were copied from noaa/consumer_test.go it would be
//nice if there was  library / package that we could consume them so they are
//tied to noaa and not to my code.
func createMessage(message string, timestamp int64) *events.Envelope {
	if timestamp == 0 {
		timestamp = time.Now().UnixNano()
	}

	logMessage := createLogMessage(message, timestamp)

	return &events.Envelope{
		LogMessage: logMessage,
		EventType:  events.Envelope_LogMessage.Enum(),
		Origin:     proto.String("fake-origin-1"),
		Timestamp:  proto.Int64(timestamp),
	}
}

func createLogMessage(message string, timestamp int64) *events.LogMessage {
	return &events.LogMessage{
		Message:     []byte(message),
		MessageType: events.LogMessage_OUT.Enum(),
		AppId:       proto.String("my-app-guid"),
		SourceType:  proto.String("DEA"),
		Timestamp:   proto.Int64(timestamp),
	}
}

var _ = Describe("Processor", func() {
	var processor *ExceptionProcessor

	BeforeEach(func() {
		processor = NewExceptionProcessor()
	})

	Describe("benign messages", func() {

		It("Don't process messages without exceptions ", func() {
			metric := processor.Process(createMessage("this is an okay message", 0))
			Expect(metric).To(BeNil())
		})
	})

	Describe("exception messages", func() {
		It("should process messages with 'exception' ", func() {
			metric := processor.Process(createMessage("this is an exception message", 0))
			Expect(metric).ToNot(BeNil())
		})
	})
})
