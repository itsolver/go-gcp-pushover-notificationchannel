package pushover_notificationchannel

import "encoding/json"

type Body struct {
	Incident Incident `json:"incident"`
}
type Incident struct {
	ID           string      `json:"incident_id"`
	ProjectID    string      `json:"scoping_project_id"`
	ProjectNum   json.Number `json:"scoping_project_number"`
	URL          string      `json:"url"`
	Started      int         `json:"started_at"`
	Ended        int         `json:"ended_at"`
	State        string      `json:"state"`
	Summary      string      `json:"summary"`
	ResourceID   string      `json:"resource_id"`
	ResourceName string      `json:"resource_name"`
	Resource     Resource    `json:"resource"`
	Metric       Metric      `json:"metric"`
	Metadata     Metadata    `json:"metadata"`
}
type Resource struct {
	Type   string            `json:"type"`
	Labels map[string]string `json:"labels"`
}
type Metric struct {
	Type   string            `json:"type"`
	Labels map[string]string `json:"labels"`
}
type Metadata struct {
	SystemLabels map[string]string `json:"system_labels"`
	UserLabels   map[string]string `json:"user_labels"`
}

const (
	// This differs from the JSON example shown on Google's documentation
	// scoping_project_id is a string **not** a number (int)
	example string = `{"version":"test","incident":{"incident_id":"12345","scoping_project_id":"12345","scoping_project_number":"12345","url":"http://www.example.com","started_at":0,"ended_at":0,"state":"OPEN","summary":"Test Incident","apigee_url":"http://www.example.com","observed_value":"1.0","resource":{"type":"example_resource","labels":{"example":"label"}},"resource_type_display_name":"Example Resource Type","resource_id":"12345","resource_display_name":"Example Resource","resource_name":"projects/12345/example_resources/12345","metric":{"type":"test.googleapis.com/metric","displayName":"Test Metric","labels":{"example":"label"}},"metadata":{"system_labels":{"example":"label"},"user_labels":{"example":"label"}},"policy_name":"projects/12345/alertPolicies/12345","policy_user_labels":{"example":"label"},"documentation":"Test documentation","condition":{"name":"projects/12345/alertPolicies/12345/conditions/12345","displayName":"Example condition","conditionThreshold":{"filter":"metric.type=\"test.googleapis.com/metric\" resource.type=\"example_resource\"","comparison":"COMPARISON_GT","thresholdValue":0.5,"duration":"0s","trigger":{"count":1}}},"condition_name":"Example condition","threshold_value":"0.5"}}`
)
