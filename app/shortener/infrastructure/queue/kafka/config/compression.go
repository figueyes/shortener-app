package config

import "github.com/segmentio/kafka-go/compress"

type Compression = compress.Compression

const (
	Gzip   = compress.Gzip
	Snappy = compress.Snappy
	Lz4    = compress.Lz4
	Zstd   = compress.Zstd
)
