package exceptionprocessor

import (
	"regexp"

	"github.com/cloudcredo/graphite-nozzle/metrics"
	"github.com/cloudfoundry/noaa/events"
)

//ExceptionProcessor searches for the word "Exception" in the log stream. We
//add a counter for every exception.
type ExceptionProcessor struct{}

//NewExceptionProcessor creates a processor
func NewExceptionProcessor() *ExceptionProcessor {
	return &ExceptionProcessor{}
}

var reg = regexp.MustCompile("(?i)exception")

//Process does the work of processing the metric. Returns nil if message has
//no exception
func (processor *ExceptionProcessor) Process(e *events.Envelope) *metrics.CounterMetric {
	hasException := reg.Match(e.GetLogMessage().GetMessage())
	if !hasException {
		return nil
	}

	stat := e.GetLogMessage().GetAppId() + "-exceptions"
	metric := metrics.NewCounterMetric(stat, 1)

	return metric
}
