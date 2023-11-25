package notifications

import (
	"fmt"
	"os/exec"
)

func NotifyAppleUsers(notification string) {
	// AppleScript command with dynamic message
	appleScript := fmt.Sprintf(`display notification "%s" with title "Go Notification"`, notification)

	// Execute the AppleScript command
	cmd := exec.Command("osascript", "-e", appleScript)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to execute AppleScript:", err)
	}
}
