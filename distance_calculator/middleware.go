package main

import (
	"time"

	"github.com/gadisamenu/tolling/types"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next CalculatorServicer
}

func NewLogMiddleware(next CalculatorServicer) CalculatorServicer {
	return &LogMiddleware{
		next: next,
	}
}

func (l *LogMiddleware) CalculateDistance(data types.ObuData) (dist float64, err error) {

	defer func(start time.Time) {
		logrus.WithFields(
			logrus.Fields{
				"took":     time.Since(start),
				"err":      err,
				"distance": dist,
			},
		).Info("calculate distance")
	}(time.Now())
	dist, err = l.next.CalculateDistance(data)
	return dist, err
}
