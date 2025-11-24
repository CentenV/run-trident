package tridentutils

import (
	"log/slog"
	"os"
	"github.com/lmittmann/tint"
	"github.com/mattn/go-isatty"
)

type logger struct {
	error *slog.Logger
	info *slog.Logger
}

var instance *logger = nil

func GetInstance() *logger {
	// Initialize singleton instance if there is no instance present
	if instance == nil {
		instance = &logger {
			// Logging ERROR messages
			error: slog.New(
				tint.NewHandler(os.Stderr, &tint.Options{
					ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
						if a.Value.Kind() == slog.KindAny {
							if _, ok := a.Value.Any().(error); ok {
								return tint.Attr(9, a)
							}
						}
						return a
					},
				}),
			),
			// Logging INFO messages
			info: slog.New(
				tint.NewHandler(os.Stdout, &tint.Options{
					NoColor: !isatty.IsTerminal(os.Stdout.Fd()),
				}),
			),
		}
	}

	return instance
}
