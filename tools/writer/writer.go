package writer

import (
	"fmt"
	"log"
	"strings"

	"github.com/fatih/color"
)

type Writer struct {
	color *color.Color
}

func NewWriter() *Writer {
	return &Writer{
		color: color.New(color.Bold),
	}
}

func (w *Writer) Errorf(format string, args ...any) {
	date := new(strings.Builder)
	log.New(date, "", log.LstdFlags).Print()

	w.color.Add(color.FgRed).Printf("%s %s", strings.TrimSpace(date.String()), fmt.Sprintf(format, args...))
}

func (w *Writer) Error(message string, args ...any) {
	date := new(strings.Builder)
	log.New(date, "", log.LstdFlags).Print()

	w.color.Add(color.FgRed).Printf("%s %s", strings.TrimSpace(date.String()), fmt.Sprintln(append([]any{message}, args...)...))
}

func (w *Writer) Logf(format string, args ...any) {
	date := new(strings.Builder)
	log.New(date, "", log.LstdFlags).Print()

	w.color.Add(color.FgGreen).Printf("%s %s", strings.TrimSpace(date.String()), fmt.Sprintf(format, args...))
}

func (w *Writer) Log(message string, args ...any) {
	date := new(strings.Builder)
	log.New(date, "", log.LstdFlags).Print()

	w.color.Add(color.FgGreen).Printf("%s %s", strings.TrimSpace(date.String()), fmt.Sprintln(append([]any{message}, args...)...))
}

func (w *Writer) Warn(format string, args ...any) {
	date := new(strings.Builder)
	log.New(date, "", log.LstdFlags).Print()

	w.color.Add(color.FgHiYellow).Printf("%s %s", strings.TrimSpace(date.String()), fmt.Sprintf(format, args...))
}
