package logger

import (
	"context"
	"io"
)

type Option func(*Options)

type Options struct {
	// The logging level the logger should log at. default is `InfoLevel`
	Level Level

	Fields map[string]interface{}

	Out io.Writer

	CallerSkipCount int

	Context context.Context
}

func WithFields(fields map[string]interface{}) Option {
	return func(args *Options) {
		args.Fields = fields
	}
}

func WithLevel(level Level) Option {
	return func(args *Options) {
		args.Level = level
	}
}

func WithOutput(out io.Writer) Option {
	return func(args *Options) {
		args.Out = out
	}
}

// WithCallerSkipCount set frame count to skip
func WithCallerSkipCount(c int) Option {
	return func(args *Options) {
		args.CallerSkipCount = c
	}
}

func SetOption(k, v interface{}) Option {
	return func(o *Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, k, v)
	}
}
