package main

import (
	"fmt"
	"featureflags"
)

func main() {
	ff := featureflags.NewFeatureFlags()

	// Load flags from a file
	err := ff.LoadFromFile("flags.json")
	if err != nil {
		fmt.Println("Error loading flags:", err)
		return
	}

	// Check if a feature is enabled
	if ff.IsEnabled("new_feature", "user123") {
		fmt.Println("New feature is enabled for user123")
	} else {
		fmt.Println("New feature is disabled for user123")
	}

	// Check percentage-based rollout
	if ff.IsEnabled("beta_feature", "user456") {
		fmt.Println("Beta feature is enabled for user456")
	} else {
		fmt.Println("Beta feature is disabled for user456")
	}
}
