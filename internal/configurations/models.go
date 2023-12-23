package main

type ProjectStatus string
type Priority string
type Status string

const (
	Development   ProjectStatus = "development"
	Testing       ProjectStatus = "testing"
	Documentation ProjectStatus = "documentation"
	Research      ProjectStatus = "research"
)

const (
	High   Priority = "high"
	Medium Priority = "medium"
	Low    Priority = "low"
)

const (
	Pending    Status = "pending"
	InProgress Status = "in progress"
	Completed  Status = "completed"
	OnHold     Status = "on hold"
)

type ProjectTasks struct {
	Title       string     `yaml:"title"`
	Description string     `yaml:"description"`
	Priority    Priority   `yaml:"priority"`
	Type        string     `yaml:"type"`
	Deadline    string     `yaml:"deadline"`
	Status      Status     `yaml:"status"`
	Assignees   []string   `yaml:"assignees"`
	Tags        []string   `yaml:"tags"`
	Reminders   []Reminder `yaml:"reminders"`
}

type Reminder struct {
	Message string   `yaml:"message"`
	Time    string   `yaml:"time"`
	Repeat  []string `yaml:"repeat"`
}

type Reminders struct {
	Types ReminderTypes `yaml:"types"`
}

type ReminderTypes struct {
	DesktopNotification       Notification               `yaml:"desktop_notification"`
	EmailNotification         EmailNotification          `yaml:"email_notification"`
	MessagingAppNotifications []MessagingAppNotification `yaml:"messaging_app_notifications"`
	CalendarNotification      CalendarNotification       `yaml:"calendar_notification"`
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
	Name     string `yaml:"name"`
	Enabled  bool   `yaml:"enabled"`
	Channel  string `yaml:"channel"`
	ServerID string `yaml:"server_id"`
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
	SyncInterval int    `yaml:"sync_interval"`
	SyncComments bool   `yaml:"sync_comments"`
}
