package utils

import (
	"github.com/google/uuid"
	"sim-backend/utils/logger"
)

func UUID() string{
	newUUID, err := uuid.NewUUID()
	if err != nil {
		logger.Errorf("Error occurs when generate UUID", err)
		return ""
	}
	return newUUID.String()
}
