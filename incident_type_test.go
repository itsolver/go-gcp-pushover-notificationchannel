package pushover_notificationchannel

import (
	"encoding/json"
	"testing"
)

func TestExample(t *testing.T) {
	got := &Body{}
	if err := json.Unmarshal([]byte(example), got); err != nil {
		t.Fatal(err)
	}
}
