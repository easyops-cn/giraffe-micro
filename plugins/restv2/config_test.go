package restv2

import (
	"testing"
	"time"
)

func TestRetryConfig_getSendCount(t *testing.T) {
	type fields struct {
		Enabled                  bool
		Retries                  uint8
		RetryIntervalMilliSecond int64
		RetryInterval            time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "false",
			fields: fields{
				Enabled: false,
			},
			want: 1,
		},
		{
			name: "true",
			fields: fields{
				Enabled: true,
				Retries: 1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &RetryConfig{
				Enabled:                  tt.fields.Enabled,
				Retries:                  tt.fields.Retries,
				RetryIntervalMilliSecond: tt.fields.RetryIntervalMilliSecond,
				RetryInterval:            tt.fields.RetryInterval,
			}
			if got := s.getSendCount(); got != tt.want {
				t.Errorf("getSendCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRetryConfig_init(t *testing.T) {
	type fields struct {
		Enabled                  bool
		Retries                  uint8
		RetryIntervalMilliSecond int64
		RetryInterval            time.Duration
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "path1",
			fields: fields{
				RetryIntervalMilliSecond: 0,
			},
		},
		{
			name: "path2",
			fields: fields{
				RetryIntervalMilliSecond: 500,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &RetryConfig{
				Enabled:                  tt.fields.Enabled,
				Retries:                  tt.fields.Retries,
				RetryIntervalMilliSecond: tt.fields.RetryIntervalMilliSecond,
				RetryInterval:            tt.fields.RetryInterval,
			}
			s.init()
		})
	}
}
