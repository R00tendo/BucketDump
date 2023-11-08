package Log

import (
	"fmt"
	"strings"

	"github.com/TwiN/go-color"
)

var templates map[string]string = map[string]string{
	"info":    color.With(color.Blue, "[INFO] MSG_HERE"),
	"error":   color.With(color.Red, "ERROR: MSG_HERE"),
	"success": color.With(color.Green, "<SUCCESS> MSG_HERE"),
}

func Msg(msg string, level string) {
	fmt.Println(strings.ReplaceAll(templates[level], "MSG_HERE", msg))
}
