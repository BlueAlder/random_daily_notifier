// FILEPATH: /home/sam/Documents/random_daily_notifier/pkg/notifier/notifier_test.go

package notifier

import (
	"fmt"
	"testing"
	"time"
)

func TestNewNotifier(t *testing.T) {
	config := NotifierConfig{
		Name:     "Test Notifier",
		Timezone: time.UTC,
		TimeRange: struct {
			Start time.Duration
			End   time.Duration
		}{
			Start: time.Duration(time.Hour * 8),
			End:   time.Duration(time.Hour * 20),
		},
	}

	notifier := NewNotifier(config)

	if notifier == nil {
		t.Error("Expected a non-nil notifier, but got nil")
	}

	if notifier.config.Name != config.Name {
		t.Errorf("Expected notifier name to be %s, but got %s", config.Name, notifier.config.Name)
	}

	// Add more assertions for other fields in the config struct
}

func TestVerifyConfig(t *testing.T) {
	config := NotifierConfig{
		Name:     "Test Notifier",
		Timezone: time.UTC,
		TimeRange: struct {
			Start time.Duration
			End   time.Duration
		}{
			Start: time.Duration(time.Hour * 8),
			End:   time.Duration(time.Hour * 20),
		},
	}

	err := verifyConfig(config)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Add more test cases to verify different scenarios for the config struct
}

func TestStart(t *testing.T) {
	config := NotifierConfig{
		Name:     "Test Notifier",
		Timezone: time.UTC,
		TimeRange: struct {
			Start time.Duration
			End   time.Duration
		}{
			Start: time.Duration(time.Hour * 8),
			End:   time.Duration(time.Hour * 20),
		},
	}

	_ = NewNotifier(config)

	// Call the Start method and assert the expected behavior
}

func TestCalculateNextNotificationTime(t *testing.T) {
	config := NotifierConfig{
		Name:     "Test Notifier",
		Timezone: time.Local,
		TimeRange: struct {
			Start time.Duration
			End   time.Duration
		}{
			Start: time.Duration(time.Hour * 8),
			End:   time.Duration(time.Hour * 20),
		},
	}

	notifier := NewNotifier(config)

	// Call the calculateNextNotificationTime method and assert the expected behavior
	notifier.calculateNextNotificationTime()
	fmt.Println(notifier.nextNotification)
}
