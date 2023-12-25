package configurations

import "time"

type Project struct {
	ProjectStatus     string            `yaml:"project_status"`
	ProjectTasks      Task              `yaml:"project_tasks"`
	Reminders         Reminders         `yaml:"reminders"`
	AutoTaskDetection AutoTaskDetection `yaml:"auto_task_detection"`
}

type Task struct {
	Title       string         `yaml:"title"`
	Description string         `yaml:"description"`
	Priority    string         `yaml:"priority"`
	Type        string         `yaml:"type"`
	Deadline    time.Time      `yaml:"deadline"`
	Status      string         `yaml:"status"`
	Assignees   []string       `yaml:"assignees"`
	Tags        []string       `yaml:"tags"`
	Reminder    []TaskReminder `yaml:"reminder"`
}

type TaskReminder struct {
	Message string   `yaml:"message"`
	Time    string   `yaml:"time"`
	Repeat  []string `yaml:"repeat"`
}

type Reminders struct {
	DesktopNotification       Notification               `yaml:"desktop_notification"`
	EmailNotification         EmailNotification          `yaml:"email_notification"`
	MessagingAppNotifications []MessagingAppNotification `yaml:"messaging_app_notifications"`
	CalendarNotification      CalendarNotification       `yaml:"calendar_notification"`
	GeneralReminder           GeneralReminder            `yaml:"general_reminder"`
}

type Notification struct {
	Enabled     bool `yaml:"enabled"`
	DisplayTime int  `yaml:"display_time"`
}

type EmailNotification struct {
	Enabled      bool   `yaml:"enabled"`
	EmailAddress string `yaml:"email_address"`
}

type MessagingAppNotification struct {
	Name    string `yaml:"name"`
	Enabled bool   `yaml:"enabled"`
	// Add additional fields for Slack, Discord, etc.
	Channel  string `yaml:"channel,omitempty"`
	ServerID string `yaml:"server_id,omitempty"`
}

type CalendarNotification struct {
	Enabled      bool   `yaml:"enabled"`
	CalendarID   string `yaml:"calendar_id"`
	RemindBefore int    `yaml:"remind_before"`
}

type GeneralReminder struct {
	Enabled  bool `yaml:"enabled"`
	Interval int  `yaml:"interval"`
}

type AutoTaskDetection struct {
	Enabled                bool                    `yaml:"enabled"`
	AnalysisInterval       int                     `yaml:"analysis_interval"`
	ProjectManagementTools []ProjectManagementTool `yaml:"project_management_tools"`
}

type ProjectManagementTool struct {
	Name         string `yaml:"name"`
	Enabled      bool   `yaml:"enabled"`
	SyncInterval int    `yaml:"sync_interval,omitempty"`
	SyncComments bool   `yaml:"sync_comments,omitempty"`
}
