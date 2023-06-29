package helpers

import (
	"time"
	"fmt"
)

func (h *Helper) GenerateID() string {
	randomID := fmt.Sprint(time.Now().UnixNano())
	return randomID
}