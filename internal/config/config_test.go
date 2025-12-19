package config

import (
	"os"
	"testing"
)

// TestLoadConfig tests configuration loading
func TestLoadConfig(t *testing.T) {
	// Clear environment variables
	originalDatabaseURL := os.Getenv("DATABASE_URL")
	originalPort := os.Getenv("PORT")
	originalJWTSecret := os.Getenv("JWT_SECRET")

	defer func() {
		if originalDatabaseURL != "" {
			os.Setenv("DATABASE_URL", originalDatabaseURL)
		}
		if originalPort != "" {
			os.Setenv("PORT", originalPort)
		}
		if originalJWTSecret != "" {
			os.Setenv("JWT_SECRET", originalJWTSecret)
		}
	}()

	// Test with defaults
	os.Unsetenv("DATABASE_URL")
	os.Unsetenv("PORT")
	os.Unsetenv("JWT_SECRET")

	cfg := LoadConfig()

	if cfg == nil {
		t.Fatal("Expected non-nil config")
	}

	if cfg.DatabaseURL == "" {
		t.Error("Expected non-empty DatabaseURL")
	}

	if cfg.Port == "" {
		t.Error("Expected non-empty Port")
	}

	if cfg.JWTSecret == "" {
		t.Error("Expected non-empty JWTSecret")
	}
}

// TestLoadConfigCustomValues tests configuration with custom environment variables
func TestLoadConfigCustomValues(t *testing.T) {
	// Set custom values
	testDatabaseURL := "postgres://custom:password@localhost:5432/custom_db"
	testPort := "9000"
	testJWTSecret := "custom-secret-key"

	os.Setenv("DATABASE_URL", testDatabaseURL)
	os.Setenv("PORT", testPort)
	os.Setenv("JWT_SECRET", testJWTSecret)

	defer func() {
		os.Unsetenv("DATABASE_URL")
		os.Unsetenv("PORT")
		os.Unsetenv("JWT_SECRET")
	}()

	cfg := LoadConfig()

	if cfg.DatabaseURL != testDatabaseURL {
		t.Errorf("Expected DatabaseURL %s, got %s", testDatabaseURL, cfg.DatabaseURL)
	}

	if cfg.Port != testPort {
		t.Errorf("Expected Port %s, got %s", testPort, cfg.Port)
	}

	if cfg.JWTSecret != testJWTSecret {
		t.Errorf("Expected JWTSecret %s, got %s", testJWTSecret, cfg.JWTSecret)
	}
}

// TestGetEnv tests the getEnv helper function
func TestGetEnv(t *testing.T) {
	testKey := "TEST_KEY_12345"
	testValue := "test-value"
	defaultValue := "default-value"

	// Test with environment variable set
	os.Setenv(testKey, testValue)
	result := getEnv(testKey, defaultValue)

	if result != testValue {
		t.Errorf("Expected %s, got %s", testValue, result)
	}

	os.Unsetenv(testKey)

	// Test with environment variable not set
	result = getEnv(testKey, defaultValue)

	if result != defaultValue {
		t.Errorf("Expected %s, got %s", defaultValue, result)
	}
}
