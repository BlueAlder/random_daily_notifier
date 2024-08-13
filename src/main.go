package main

import (
	"fmt"
	"time"

	"github.com/BlueAlder/random-daily-notifier/pkg/notifier"
)

type Messenger struct {
	Notifier *notifier.Notifier
	URL      string
}

func (m *Messenger) SendMessage() {
	fmt.Println("Sending message")

}

func main() {

	// 1. Read config

	// 2. Parse config into NotifierConfig struct

	// 3. Create a new Notifier instance with the parsed config

	// 4. Calculate the next notification time

	// 5. Start the notifier loop

	m := &Messenger{}

	startDuration, _ := time.ParseDuration("18h31m")
	endDuration, _ := time.ParseDuration("18h32m")

	config := notifier.NotifierConfig{
		Name:     "Test Notifier",
		Timezone: time.Local,
		TimeRange: struct {
			Start time.Duration
			End   time.Duration
		}{
			Start: startDuration,
			End:   endDuration,
		},
		Callback: m.SendMessage,
	}

	m.Notifier = notifier.NewNotifier(config)
	m.Notifier.Start()
}
