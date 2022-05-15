package pushover_notificationchannel

type Body struct {
	Incident Incident `json:"incident"`
}
type Incident struct {
	ID           string   `json:"incident_id"`
	ProjectID    string   `json:"scoping_project_id"`
	ProjectNum   int      `json:"scoping_project_number"`
	URL          string   `json:"url"`
	Started      int      `json:"started_at"`
	Ended        int      `json:"ended_at"`
	State        string   `json:"state"`
	Summary      string   `json:"summary"`
	ResourceID   string   `json:"resource_id"`
	ResourceName string   `json:"resource_name"`
	Resource     Resource `json:"resource"`
	Metric       Metric   `json:"metric"`
	Metadata     Metadata `json:"metadata"`
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
	example string = `
{
	"incident": {
	  "incident_id": "0.opqiw61fsv7p",
	  "scoping_project_id": "internal-project",
	  "scoping_project_number": 12345,
	  "url": "https://console.cloud.google.com/monitoring/alerting/incidents/0.lxfiw61fsv7p?project=internal-project",
	  "started_at": 1577840461,
	  "ended_at": 1577877071,
	  "state": "closed",
	  "resource_id": "11223344",
	  "resource_name": "internal-project gke-cluster-1-default-pool-e2df4cbd-dgp3",
	  "resource_display_name": "gke-cluster-1-default-pool-e2df4cbd-dgp3",
	  "resource_type_display_name": "VM Instance",
	  "resource": {
		"type": "gce_instance",
		"labels": {
		  "instance_id": "11223344",
		  "project_id": "internal-project",
		  "zone": "us-central1-c"
		}
	  },
	  "metric": {
		"type": "compute.googleapis.com/instance/cpu/utilization",
		"displayName": "CPU utilization",
		"labels": {
		  "instance_name": "the name of the VM instance"
		}
	  },
	  "metadata": {
		"system_labels": { "labelkey": "labelvalue" },
		"user_labels": { "labelkey": "labelvalue" }
	  },
	  "policy_name": "Monitor-Project-Cluster",
	  "policy_user_labels" : {
		  "user-label-1" : "important label",
		  "user-label-2" : "another label"
	  },
	  "condition_name": "VM Instance - CPU utilization [MAX]",
	  "threshold_value": "0.9",
	  "observed_value": "0.835",
	  "condition": {
		"name": "projects/internal-project/alertPolicies/1234567890123456789/conditions/1234567890123456789",
		"displayName": "VM Instance - CPU utilization [MAX]",
		"conditionThreshold": {
		  "filter": "metric.type=\"compute.googleapis.com/instance/cpu/utilization\" resource.type=\"gce_instance\" metadata.system_labels.\"state\"=\"ACTIVE\"",
		  "aggregations": [
			{
			  "alignmentPeriod": "120s",
			  "perSeriesAligner": "ALIGN_MEAN"
			}
		  ],
		  "comparison": "COMPARISON_GT",
		  "thresholdValue": 0.9,
		  "duration": "0s",
		  "trigger": {
			"count": 1
		  }
		}
	  },
	  "documentation": {
		"content": "TEST ALERT\n\npolicy.name=projects/internal-project/alertPolicies/1234567890123456789\n\npolicy.display_name=Monitored-Project-NO-GROUPBY\n\ncondition.name=projects/nternal-project/alertPolicies/1234567890123456789/conditions/1234567890123456789\n\ncondition.display_name=VM Instance - CPU utilization [MAX]\n\nproject=internal-project\n\nresrouce.project=internal-project \n\nDONE\n",
		"mime_type": "text/markdown"
	  },
	  "summary": "CPU utilization for internal-project gke-cluster-1-16-default-pool-e2df4cbd-dgp3 with metric labels {instance_name=gke-cluster-1-default-pool-e2df4cbd-dgp3} and system labels {state=ACTIVE} returned to normal with a value of 0.835."
	},
	"version": "1.2"
  }
`
)
