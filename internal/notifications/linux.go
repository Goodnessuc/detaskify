package notifications

import (
	"fmt"
	"os/exec"
)

// TODO: Make sure you provide installation for lib-notify-gen sudo apt-get install libnotify-bin

func NotifyLinuxUsers(notification string) {
	// Command to send a notification
	cmd := exec.Command("notify-send", "Go Notification", notification)

	// Execute the command
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to execute notify-send:", err)
	}
}
