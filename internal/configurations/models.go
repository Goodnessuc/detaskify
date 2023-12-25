package configurations

type Project struct {
	Name          string   `yaml:"name"`
	Status        string   `yaml:"status"`
	Collaborators []string `yaml:"collaborators"`
	Notes         string   `yaml:"notes"`
	URL           string   `yaml:"url"`
}

type Task struct {
	Group       string     `yaml:"group"`
	Title       string     `yaml:"title"`
	Description string     `yaml:"description"`
	Priority    string     `yaml:"priority"`
	Type        string     `yaml:"type"`
	Deadline    string     `yaml:"deadline"`
	Status      string     `yaml:"status"`
	Assignees   []string   `yaml:"assignees"`
	Tags        []string   `yaml:"tags"`
	Reminder    []Reminder `yaml:"reminder"`
	Comment     string     `yaml:"comment"`
}

type Reminder struct {
	Message string   `yaml:"message"`
	Time    string   `yaml:"time"`
	Repeat  []string `yaml:"repeat"`
}

type Schedule struct {
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

type Config struct {
	Project  Project  `yaml:"project"`
	Tasks    []Task   `yaml:"tasks"`
	Schedule Schedule `yaml:"schedule"`
}
