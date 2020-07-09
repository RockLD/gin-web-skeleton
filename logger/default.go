package logger

import (
	"context"
	"fmt"
	dlog "github.com/micro/go-micro/v2/debug/log"
	"os"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
)

func init() {
	lvl, err := GetLevel(os.Getenv("MICRO_LOG_LEVEL"))
	if err != nil {
		lvl = InfoLevel
	}

	DefaultLogger = NewHelper(NewLogger(WithLevel(lvl)))
}

type defaultLogger struct {
	sync.RWMutex
	opts Options
}

func (l *defaultLogger) Init(opts ...Option) error {
	for _, o := range opts {
		o(&l.opts)
	}
	return nil
}

func (l *defaultLogger) Options() Options {
	l.RLock()
	defer l.RUnlock()
	opts := l.opts
	opts.Fields = copyFields(l.opts.Fields)
	return opts
}

func (l *defaultLogger) Fields(fields map[string]interface{}) Logger {
	l.Lock()
	defer l.Unlock()
	l.opts.Fields = copyFields(fields)
	//l.Unlock()
	return l
}

func copyFields(src map[string]interface{}) map[string]interface{} {
	dst := make(map[string]interface{}, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func logCallerfilePath(loggingFilePath string) string {
	idx := strings.LastIndexByte(loggingFilePath, '/')
	if idx == -1 {
		return loggingFilePath
	}

	idx = strings.LastIndexByte(loggingFilePath[:idx], '/')
	if idx == -1 {
		return loggingFilePath
	}
	return loggingFilePath[idx+1:]
}

func (l defaultLogger) Log(level Level, v ...interface{}) {
	if !l.opts.Level.Enabled(level) {
		return
	}

	l.RLock()
	fields := copyFields(l.opts.Fields)
	l.RUnlock()

	fields["level"] = level.String()

	if _, file, line, ok := runtime.Caller(l.opts.CallerSkipCount); ok {
		fields["file"] = fmt.Sprintf("%s:%d", logCallerfilePath(file), line)
	}

	rec := dlog.Record{
		Timestamp: time.Now(),
		Message:   fmt.Sprint(v...),
		Metadata:  make(map[string]string, len(fields)),
	}

	keys := make([]string, 0, len(fields))
	for k, v := range fields {
		keys = append(keys, k)
		rec.Metadata[k] = fmt.Sprintf("%v", v)
	}

	sort.Strings(keys)

	metadata := ""

	for _, k := range keys {
		metadata += fmt.Sprintf(" %s=%v", k, fields[k])
	}

	dlog.DefaultLog.Write(rec)

	t := rec.Timestamp.Format("2006-01-02 15:04:05")

	fmt.Printf("%s %s %v\n", t, metadata, rec.Message)
}

func (l *defaultLogger) Logf(level Level, format string, v ...interface{}) {
	//	 TODO decide does we need to write message if log level not used?
	if level < l.opts.Level {
		return
	}

	l.RLock()
	fields := copyFields(l.opts.Fields)
	l.RUnlock()

	fields["level"] = level.String()

	if _, file, line, ok := runtime.Caller(l.opts.CallerSkipCount); ok {
		fields["file"] = fmt.Sprintf("%s:%d", logCallerfilePath(file), line)
	}

	rec := dlog.Record{
		Timestamp: time.Now(),
		Message:   fmt.Sprintf(format, v...),
		Metadata:  make(map[string]string, len(fields)),
	}

	keys := make([]string, 0, len(fields))
	for k, v := range fields {
		keys = append(keys, k)
		rec.Metadata[k] = fmt.Sprintf("%v", v)
	}

	sort.Strings(keys)
	metadata := ""

	for _, k := range keys {
		metadata += fmt.Sprintf(" %s=%v", k, fields[k])
	}

	dlog.DefaultLog.Write(rec)

	t := rec.Timestamp.Format("2006-01-02 15:04:05")
	fmt.Printf("%s %s %v\n", t, metadata, rec.Message)
}

func (l *defaultLogger) String() string {
	return "default"
}

func NewLogger(opts ...Option) Logger {
	options := Options{
		Level:           InfoLevel,
		Fields:          make(map[string]interface{}),
		Out:             os.Stderr,
		CallerSkipCount: 2,
		Context:         context.Background(),
	}

	l := &defaultLogger{opts: options}

	if err := l.Init(opts...); err != nil {
		l.Log(FatalLevel, err)
	}

	return l
}
