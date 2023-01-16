package restv2

import "time"

const (
	defaultWaitDuration = 500 * time.Millisecond
)

type RPCRetryConfig map[string]RetryConfig

type RetryConfig struct {
	Enabled                  bool          `json:"enabled" yaml:"enabled"`
	Retries                  uint8         `json:"retries" yaml:"retries"`
	RetryIntervalMilliSecond int64         `json:"retry_interval_millisecond" yaml:"retry_interval_millisecond"`
	RetryInterval            time.Duration `json:"-" yaml:"-"`
}

func (s *RetryConfig) init() {
	// 如果 retry_interval 没有设置, 则设置为500ms
	if s.RetryIntervalMilliSecond <= 0 {
		s.RetryInterval = defaultWaitDuration
	} else {
		s.RetryInterval = time.Duration(s.RetryIntervalMilliSecond) * time.Millisecond
	}
}

func (s *RetryConfig) getSendCount() int {
	if !s.Enabled {
		return 1
	}
	return int(s.Retries + 1)
}
