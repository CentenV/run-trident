package logger

import (
	"log/slog"
	"os"
	"github.com/lmittmann/tint"
	"github.com/mattn/go-isatty"
)

var instance *slog.Logger = nil

func GetInstance() *slog.Logger {
	// Initialize singleton instance if there is no instance present
	if instance == nil {
		instance = slog.New(
			tint.NewHandler(os.Stderr, &tint.Options{
				ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
					if a.Value.Kind() == slog.KindAny {
						if _, ok := a.Value.Any().(error); ok {
							return tint.Attr(9, a)
						}
					}
					return a
				},
				NoColor: !isatty.IsTerminal(os.Stdout.Fd()),
			}),
		)
	}

	return instance
}
