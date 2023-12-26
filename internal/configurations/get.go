package configurations

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

// ReadYAML reads a YAML file and returns the byte slice.
func ReadYAML(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UnmarshalYAMLData unmarshal the YAML data into the Data struct.
func UnmarshalYAMLData(yamlData []byte) (*Data, error) {
	var data Data
	err := yaml.Unmarshal(yamlData, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// GetProject returns the Project data.
func GetProject(data *Data) *Project {
	return &data.Project
}

// GetTasks returns the slice of Tasks.
func GetTasks(data *Data) []Task {
	return data.Tasks
}

// GetSchedule returns the slice of ScheduleItems.
func GetSchedule(data *Data) []ScheduleItem {
	return data.Schedule
}
