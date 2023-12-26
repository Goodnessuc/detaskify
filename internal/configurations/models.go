package configurations

type Project struct {
	Name          string   `yaml:"name"`
	Status        string   `yaml:"status"`
	URL           string   `yaml:"url"`
	Collaborators []string `yaml:"collaborators"`
	Description   string   `yaml:"description"`
}

type Reminder struct {
	Message string   `yaml:"message"`
	Time    string   `yaml:"time"`
	Repeat  []string `yaml:"repeat"`
}

type Task struct {
	ID          string     `yaml:"id"`
	Group       string     `yaml:"group"`
	Title       string     `yaml:"title"`
	Description string     `yaml:"description"`
	Priority    string     `yaml:"priority"`
	Type        string     `yaml:"type"`
	Deadline    string     `yaml:"deadline"`
	Status      string     `yaml:"status"`
	Assignees   []string   `yaml:"assignees"`
	Tags        []string   `yaml:"tags"`
	Reminders   []Reminder `yaml:"reminders"`
	Comments    []string   `yaml:"comments"`
}

type ScheduleItem struct {
	Title       string     `yaml:"title"`
	Description string     `yaml:"description"`
	Priority    string     `yaml:"priority"`
	Type        string     `yaml:"type"`
	Deadline    string     `yaml:"deadline"`
	Command     string     `yaml:"command"`
	Status      string     `yaml:"status"`
	Reminder    []Reminder `yaml:"reminder"`
	Repeat      []string   `yaml:"repeat"`
}

type Data struct {
	Project  Project        `yaml:"project"`
	Tasks    []Task         `yaml:"tasks"`
	Schedule []ScheduleItem `yaml:"schedule"`
}
