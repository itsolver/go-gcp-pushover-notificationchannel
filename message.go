package pushover_notificationchannel

// Message is a type that represents the Pushover message
type Message struct {
	ProjectID    string
	State        string
	Summary      string
	Started      int
	Ended        int
	SystemLabels map[string]string
}

// NewMessage is a function that maps an Incident body into a Pushover message
func NewMessage(body *Body) (*Message, error) {
	return &Message{
		ProjectID:    body.Incident.ProjectID,
		State:        body.Incident.State,
		Summary:      body.Incident.Summary,
		Started:      body.Incident.Started,
		Ended:        body.Incident.Ended,
		SystemLabels: body.Incident.Metadata.SystemLabels,
	}, nil
}
