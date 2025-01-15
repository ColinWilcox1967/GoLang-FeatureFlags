package featureflags

import (
	"testing"
)

func TestFeatureFlags(t *testing.T) {
	ff := NewFeatureFlags()

	// Set a feature flag
	ff.SetFlag(Flag{
		Name:           "new_feature",
		Enabled:        true,
		PercentageRollout: 0,
	})

	// Check if the feature flag is enabled
	if !ff.IsEnabled("new_feature", "user123") {
		t.Error("Expected new_feature to be enabled")
	}

	// Set a percentage rollout flag
	ff.SetFlag(Flag{
		Name:           "beta_feature",
		Enabled:        false,
		PercentageRollout: 50,
	})

	// Check if the percentage rollout works
	if !ff.IsEnabled("beta_feature", "user123") {
		t.Log("beta_feature is disabled for user123 (expected due to percentage rollout)")
	}
}
