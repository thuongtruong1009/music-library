package helpers

import (
	"time"
	"fmt"
)

func (h *Helper) GenerateID() string {
	randomID := fmt.Sprint(time.Now().UnixNano())
	return randomID
}

func (h *Helper) ConvertArrayToString(arr []string) string {
	var output string
	for _, v := range arr {
		output += v + ","
	}
	return output
}