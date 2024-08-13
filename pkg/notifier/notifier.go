package notifier

import (
	"fmt"
	"math/rand"
	"time"
)

type Notifier struct {
	config           NotifierConfig
	nextNotification time.Time
}

type NotifierConfig struct {
	Name      string
	Timezone  *time.Location
	TimeRange struct {
		// Using durations here because parsing is easy and cbf with time.Time
		Start time.Duration
		End   time.Duration
	}
	Callback func()
}

func NewNotifier(config NotifierConfig) *Notifier {
	if err := verifyConfig(config); err != nil {
		panic(err)
	}

	// Create a new Notifier instance
	return &Notifier{
		config: config,
	}
}

func verifyConfig(config NotifierConfig) error {
	// Verify the config
	if config.Timezone == nil {
		return fmt.Errorf("timezone cannot be nil")
	}

	if config.TimeRange.Start < 0 || config.TimeRange.Start > time.Hour*24 {
		return fmt.Errorf("invalid start time")
	}

	if config.TimeRange.End < 0 || config.TimeRange.End > time.Hour*24 {
		return fmt.Errorf("invalid end time")
	}

	if config.TimeRange.Start > config.TimeRange.End {
		return fmt.Errorf("start time must be before end time")
	}

	if config.Callback == nil {
		fmt.Println("⚠️ callback is nil")
	}

	return nil
}

func (n *Notifier) Start() {
	// Start the notifier loop
	for {
		// Calculate the next notification time
		n.calculateNextNotificationTime()
		fmt.Println("Next notification coming up at", n.nextNotification)

		// Wait until the next notification time
		time.Sleep(time.Until(n.nextNotification))

		// Send the notification
		n.sendNotification()
	}

}

func (n *Notifier) calculateNextNotificationTime() {
	now := time.Now().In(n.config.Timezone)

	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, n.config.Timezone).Add(n.config.TimeRange.Start)
	end := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, n.config.Timezone).Add(n.config.TimeRange.End)
	// Check if the current hour is within the time range
	if now.After(start) {
		// Go to next day
		start = start.Add(24 * time.Hour)
		end = end.Add(24 * time.Hour)
	}
	// Random number between two unix timestamps
	randomUnix := start.Unix() + rand.Int63n(end.Unix()-start.Unix())

	// Set the next notification time to the random time within the time range
	n.nextNotification = time.Unix(randomUnix, 0)
}

func (n *Notifier) sendNotification() {
	// Send the notification to the endpoint
	fmt.Printf("Sending notification\n")
	n.config.Callback()
}
