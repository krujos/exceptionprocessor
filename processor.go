package exceptionprocessor

import (
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

//Process does the work of processing the metric.
func (processor *ExceptionProcessor) Process(e *events.Envelope) []metrics.Metric {
	return nil
}
