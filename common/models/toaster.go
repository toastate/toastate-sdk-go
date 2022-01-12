package models

type Toaster struct {
	ID      string `json:"id,omitempty"`
	OwnerID string `json:"owner_id,omitempty"`

	CryptoSecure bool `json:"cryptographically_secure,omitempty"`

	BuildCmd []string `json:"build_command,omitempty"`
	ExeCmd   []string `json:"execution_command,omitempty"`
	Env      []string `json:"environment_variables,omitempty"`

	JoinableForSec       int `json:"joinable_for_seconds,omitempty"`
	MaxConcurrentJoiners int `json:"max_concurrent_joiners,omitempty"`
	TimeoutSec           int `json:"timeout_seconds,omitempty"`

	Name     string   `json:"name,omitempty"`
	Readme   string   `json:"readme,omitempty"`
	Keywords []string `json:"keywords,omitempty"`

	Version int `json:"version,omitempty"`
}

type ToasterStats struct {
	// Milliseconds
	AggregatedDuration int64 `json:"durationms,omitempty"`

	CPUSeconds int64 `json:"cpus,omitempty"`

	// GigaBytes-seconds
	RAM float64 `json:"ramgbs,omitempty"`

	// Bytes
	NetIngress float64 `json:"ingress,omitempty"`

	// Bytes
	NetEgress float64 `json:"egress,omitempty"`
}
