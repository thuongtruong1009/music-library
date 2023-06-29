package helpers

import (
	"music-management-system/pkg/constants"
)

func (h *Helper) OutputColor(level int8) string {
	switch level {
	case constants.DEBUG:
		return constants.Gray
	case constants.INFO:
		return constants.Green
	case constants.WARNING:
		return constants.Orange
	case constants.ERROR:
		return constants.Red
	case constants.FATAL:
		return constants.Purple
	case constants.LABEL:
		return constants.Cyan
	case constants.DESC:
		return constants.Yellow
	case constants.INPUT:
		return constants.Cyan
	case constants.QUERY:
		return constants.Pink
	default:
		return constants.White
	}
}

func (h *Helper) OutputSuccess(statusMsg string) {
	println(constants.Green + "\n::: Status: " + statusMsg + constants.Reset)
}

func (h *Helper) OutputError(statusMsg string, err string) {
	println(constants.Red + "\n::: Status: " + statusMsg + "\n::: Message: " + err + constants.Reset)
}

func (h *Helper) OutputNomal(label int8, msg string) {
	println(h.OutputColor(label) + msg + constants.Reset)
}

func OutputTime(time string) {
	println(constants.Purple + "\n::: Query time: " + time + constants.Reset)
}