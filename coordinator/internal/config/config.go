package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"scroll-tech/common/database"
)

// ProverManager loads sequencer configuration items.
type ProverManager struct {
	// The amount of provers to pick per proof generation session.
	ProversPerSession uint8 `json:"provers_per_session"`
	// Number of attempts that a session can be retried if previous attempts failed.
	// Currently we only consider proving timeout as failure here.
	SessionAttempts uint8 `json:"session_attempts"`
	// Zk verifier config.
	Verifier *VerifierConfig `json:"verifier"`
	// BatchCollectionTimeSec batch Proof collection time (in seconds).
	BatchCollectionTimeSec int `json:"batch_collection_time_sec"`
	// ChunkCollectionTimeSec chunk Proof collection time (in seconds).
	ChunkCollectionTimeSec int `json:"chunk_collection_time_sec"`
	// BundleCollectionTimeSec bundle Proof collection time (in seconds).
	BundleCollectionTimeSec int `json:"bundle_collection_time_sec"`
	// Max number of workers in verifier worker pool
	MaxVerifierWorkers int `json:"max_verifier_workers"`
	// MinProverVersion is the minimum version of the prover that is required.
	MinProverVersion string `json:"min_prover_version"`
}

// L2 loads l2geth configuration items.
type L2 struct {
	// l2geth chain_id.
	ChainID uint64 `json:"chain_id"`
}

// Auth provides the auth coordinator
type Auth struct {
	Secret                     string `json:"secret"`
	ChallengeExpireDurationSec int    `json:"challenge_expire_duration_sec"`
	LoginExpireDurationSec     int    `json:"login_expire_duration_sec"`
}

// Config load configuration items.
type Config struct {
	ProverManager *ProverManager   `json:"prover_manager"`
	DB            *database.Config `json:"db"`
	L2            *L2              `json:"l2"`
	Auth          *Auth            `json:"auth"`
}

// VerifierConfig load zk verifier config.
type VerifierConfig struct {
	ForkName     string `json:"fork_name"`
	MockMode     bool   `json:"mock_mode"`
	ParamsPath   string `json:"params_path"`
	AssetsPathLo string `json:"assets_path_lo"` // lower version Verifier
	AssetsPathHi string `json:"assets_path_hi"` // higher version Verifier
}

// NewConfig returns a new instance of Config.
func NewConfig(file string) (*Config, error) {
	buf, err := os.ReadFile(filepath.Clean(file))
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = json.Unmarshal(buf, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
