package table

import (
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
)

func Generator() string {
	s, err := uuid.NewUUID()
	if err != nil {
		log.Errorf("get new uuid error %v", err)
		return ""
	}
	return s.String()
}
