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

func (w *Writer) Error(format string, args ...any) {
	date := new(strings.Builder)
	log.New(date, "", log.LstdFlags).Print()

	w.color.Add(color.FgRed).Printf("%s %s", strings.TrimSpace(date.String()), fmt.Sprintf(format, args...))
}

func (w *Writer) Log(format string, args ...any) {
	date := new(strings.Builder)
	log.New(date, "", log.LstdFlags).Print()

	w.color.Add(color.FgGreen).Printf("%s %s", strings.TrimSpace(date.String()), fmt.Sprintf(format, args...))
}

func (w *Writer) Warn(format string, args ...any) {
	date := new(strings.Builder)
	log.New(date, "", log.LstdFlags).Print()

	w.color.Add(color.FgHiYellow).Printf("%s %s", strings.TrimSpace(date.String()), fmt.Sprintf(format, args...))
}
