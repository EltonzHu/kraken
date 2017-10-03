package blobserver

import "time"

// Config defines the configuration used by Origin cluster for hashing blob digests.
type Config struct {
	NumReplica int                       `yaml:"num_replica"`
	HashNodes  map[string]HashNodeConfig `yaml:"hash_nodes"`
	Repair     RepairConfig              `yaml:"repair"`
}

// HashNodeConfig defines the config for a single origin node
type HashNodeConfig struct {
	Label  string `yaml:"label"`
	Weight int    `yaml:"weight"`
}

// RepairConfig defines the configuration used by Origin cluster for
// running concurrent repairs.
type RepairConfig struct {
	NumWorkers  int           `yaml:"num_workers"`
	NumRetries  int           `yaml:"num_retries"`
	RetryDelay  time.Duration `yaml:"retry_delay"`
	ConnTimeout time.Duration `yaml:"conn_timeout"`
}
