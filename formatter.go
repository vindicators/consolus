package consolus

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"os"
	"sync"
	"strings"
	"fmt"
)




type ConsoleFormatter struct {
	// Set to true to bypass checking for a TTY before outputting colors.
	ForceColors bool

	// Force disabling colors. For a TTY colors are enabled by default.
	DisableColors bool

	// Force formatted layout, even for non-TTY output.
	ForceFormatting bool

	// Whether the logger's out is to a terminal.
	isTerminal bool

	// Timestamp format to use for display when a full timestamp is printed.
	TimestampFormat string

	// The fields are sorted by default for a consistent output. For applications
	// that log extremely frequently and don't use the JSON formatter this may not
	// be desired.
	DisableSorting bool

	sync.Once
}


func (f *ConsoleFormatter) init(entry *logrus.Entry) {
	if entry.Logger != nil {
		f.isTerminal = f.checkIfTerminal(entry.Logger.Out)
	}
}

func (f *ConsoleFormatter) checkIfTerminal(w io.Writer) bool {
	switch v := w.(type) {
	case *os.File:
		return terminal.IsTerminal(int(v.Fd()))
	default:
		return false
	}
}

func (f *ConsoleFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer

	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	f.Do(func() { f.init(entry) })

	isFormatted := f.ForceFormatting || !f.isTerminal

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	if isFormatted {
		f.printFormatted(b, entry)
	} else {
		f.printConsole(b, entry)
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}

func (f *ConsoleFormatter) printFormatted(b *bytes.Buffer, entry * logrus.Entry)
{


	printConsole(b, formattedEntry)
}

func (f *ConsoleFormatter) printConsole(b *bytes.Buffer, entry * logrus.Entry)
{

}
