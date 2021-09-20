package rotation

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Config struct {
	filename   string
	maxSize    int
	maxBackups int
	maxAge     int
	compress   bool
}

type Option func(*Config)

func WithFilename(filename string) Option {
	return func(config *Config) {
		config.filename = filename
	}
}

func WithMaxSize(maxSize int) Option {
	return func(config *Config) {
		config.maxSize = maxSize
	}
}

func WithMaxBackups(maxBackups int) Option {
	return func(config *Config) {
		config.maxBackups = maxBackups
	}
}
func WithMaxAge(maxAge int) Option {
	return func(config *Config) {
		config.maxAge = maxAge
	}
}

func WithCompress(compress bool) Option {
	return func(config *Config) {
		config.compress = compress
	}
}

func defaultConfig() *Config {
	return &Config{
		filename:   "./app.log",
		maxSize:    100,
		maxBackups: 0,
		maxAge:     1,
		compress:   false,
	}
}

type rotationHook struct {
	logger *lumberjack.Logger
}

func NewHook(opts ...Option) (*rotationHook, error) {
	cfg, err := generateConfig(opts...)
	if err != nil {
		return nil, err
	}

	logger := &lumberjack.Logger{
		Filename:   cfg.filename,
		MaxSize:    cfg.maxSize,
		MaxBackups: cfg.maxBackups,
		MaxAge:     cfg.maxAge,
		Compress:   cfg.compress,
	}

	return &rotationHook{logger: logger}, nil
}

// Fire is called when a log event is fired.
func (hook *rotationHook) Fire(entry *logrus.Entry) error {
	line, err := entry.Bytes()
	if err != nil {
		return err
	}

	_, err = hook.logger.Write(line)
	if err != nil {
		return err
	}

	return nil
}

// Levels returns the available logging levels
func (hook *rotationHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.TraceLevel,
		logrus.DebugLevel,
		logrus.InfoLevel,
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
}
func generateConfig(opts ...Option) (*Config, error) {
	conf := defaultConfig()
	for _, opt := range opts {
		opt(conf)
	}

	if err := checkConfig(conf); err != nil {
		return nil, fmt.Errorf("create p2p: %w", err)
	}

	return conf, nil
}

func checkConfig(config *Config) error {
	if config.maxAge == 0 || config.maxSize == 0 {
		return fmt.Errorf("max age and size can not be 0")
	}

	return nil
}