package utils

import (
	"fmt"
)

var ErrorHandler ErrorHandleLogger

type ErrorHandleLogger struct {
	Logger LogsHandler
}

func (h *ErrorHandleLogger) Handle(err error, initiatorCommand string) {
	formattedError := fmt.Errorf("ошибка при исполнении команды '%s', ошибка: %w", initiatorCommand, err)

	h.Logger.LogError(formattedError)
}

func (h *ErrorHandleLogger) NotifyUserAboutError(content string) {

}
