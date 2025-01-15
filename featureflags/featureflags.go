package featureflags

import (
	"encoding/json"
	"errors"
	"math/rand"
	"os"
	"sync"
	"time"
)

// Flag represents a feature flag with various attributes.
type Flag struct {
	Name           string  `json:"name"`
	Enabled        bool    `json:"enabled"`
	PercentageRollout float64 `json:"percentage_rollout"`
}

// FeatureFlags manages a collection of feature flags.
type FeatureFlags struct {
	flags map[string]Flag
	mu    sync.RWMutex
}

// NewFeatureFlags initializes a new FeatureFlags instance.
func NewFeatureFlags() *FeatureFlags {
	return &FeatureFlags{
		flags: make(map[string]Flag),
	}
}

// LoadFromFile loads feature flags from a JSON file.
func (f *FeatureFlags) LoadFromFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var flags []Flag
	if err := json.NewDecoder(file).Decode(&flags); err != nil {
		return err
	}

	f.mu.Lock()
	defer f.mu.Unlock()

	for _, flag := range flags {
		f.flags[flag.Name] = flag
	}

	return nil
}

// IsEnabled checks if a feature flag is enabled for a given key (e.g., userID).
func (f *FeatureFlags) IsEnabled(flagName string, key string) bool {
	f.mu.RLock()
	defer f.mu.RUnlock()

	flag, exists := f.flags[flagName]
	if !exists {
		return false
	}

	// Check if the flag is globally enabled
	if flag.Enabled {
		return true
	}

	// Check percentage-based rollout
	if flag.PercentageRollout > 0 {
		rand.Seed(time.Now().UnixNano() + int64(hashString(key)))
		return rand.Float64()*100 < flag.PercentageRollout
	}

	return false
}

// SetFlag sets or updates a feature flag at runtime.
func (f *FeatureFlags) SetFlag(flag Flag) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.flags[flag.Name] = flag
}

// GetFlag retrieves a feature flag by name.
func (f *FeatureFlags) GetFlag(flagName string) (Flag, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()

	flag, exists := f.flags[flagName]
	if !exists {
		return Flag{}, errors.New("flag not found")
	}

	return flag, nil
}

// hashString generates a simple hash for consistent rollout percentage.
func hashString(s string) int {
	hash := 0
	for _, c := range s {
		hash += int(c)
	}
	return hash
}
