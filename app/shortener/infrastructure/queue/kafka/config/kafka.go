package config

import (
	"time"
)

type Connection struct {
	Username string
	Password string
}

// ReaderConfig is a configuration object
type ReaderConfig struct {
	Connection
	Brokers        []string
	GroupID        string
	MinBytes       int
	MaxBytes       int
	MaxAttempts    int
	CommitInterval time.Duration
	Logger         Logger
	ErrorLogger    Logger
}

// WriterConfig is a configuration object
type WriterConfig struct {
	Connection
	Brokers     []string
	Logger      Logger
	ErrorLogger Logger
	Compression Compression
}
