package pushover_notificationchannel

import "strings"

// Message is a type that represents the Pushover message
type Message struct {
	Title        string
	RevisionName string
	State        string
	ProjectID    string
}

// extractRevisionName gets the revision name from the full summary
func extractRevisionName(summary string) string {
	if strings.Contains(summary, "revision_name=") {
		parts := strings.Split(summary, "revision_name=")
		if len(parts) > 1 {
			revName := strings.Split(parts[1], ",")[0]
			revName = strings.TrimPrefix(revName, "auto-solve-thanks-tickets-")
			revName = strings.TrimSuffix(revName, "}")
			return revName
		}
	}
	return ""
}

// NewMessage is a function that maps an Incident body into a Pushover message
func NewMessage(body *Body) (*Message, error) {
	// Create a human-readable title
	title := "Cloud Run Alert"
	if body.Incident.ResourceName != "" {
		title = "Cloud Run: " + extractRevisionName(body.Incident.Summary)
	}

	return &Message{
		Title:        title,
		RevisionName: extractRevisionName(body.Incident.Summary),
		State:        body.Incident.State,
		ProjectID:    body.Incident.ProjectID,
	}, nil
}
