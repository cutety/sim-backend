package logger

import "github.com/sirupsen/logrus"

type GormLoggerI struct{}

func (*GormLoggerI) Print(v ...interface{}) {
	format := v[0].(string)
	caller := v[1].(string)
	v = v[3:]

	logger.WithFields(logrus.Fields{"caller": caller}).Infof(format+"%v", v)
}

func NewGormLogger() *GormLoggerI {
	return &GormLoggerI{}
}