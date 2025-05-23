package config

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	MaxBatchSize int
	MaxLatencyMs int
	ListenAddr   string
}

func Load() Config {
	pflag.Int("max-batch-size", 8, "maximum number of requests per batch")
	pflag.Int("max-latency-ms", 50, "maximum latency before flushing a batch")
	pflag.String("listen-addr", ":8080", "address to bind to")

	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	viper.SetEnvPrefix("livebatch")
	viper.AutomaticEnv()

	return Config{
		MaxBatchSize: viper.GetInt("max-batch-size"),
		MaxLatencyMs: viper.GetInt("max-latency-ms"),
		ListenAddr:   viper.GetString("listen-addr"),
	}
}

func (c Config) String() string {
	return fmt.Sprintf("LiveBatch Config: batch_size=%d, latency_ms=%d, listen=%s",
		c.MaxBatchSize, c.MaxLatencyMs, c.ListenAddr)
}
