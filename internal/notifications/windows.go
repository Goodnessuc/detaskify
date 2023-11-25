package notifications

import (
	"fmt"
	"github.com/go-toast/toast"
)

// NotifyLinuxUsers TODO: Make sure you provide installation for go-toast

func NotifyWindowsUsers(message string) {
	notification := toast.Notification{
		AppID:   "Detaskify",
		Title:   "Detaskify Notification",
		Message: message, // Use the parameter 'message' here
		// Icon: "", // Optional: Path to an icon file
		// Actions: []toast.Action{}, // Optional: Define actions for the notification
	}

	err := notification.Push()
	if err != nil {
		fmt.Println("Failed to push notification:", err)
	}
}
