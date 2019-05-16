package getstream

import "time"

// Activity is a Stream activity entity.
type Activity struct {
	ID        string                 `json:"id,omitempty"`
	Actor     string                 `json:"actor,omitempty"`
	Verb      string                 `json:"verb,omitempty"`
	Object    string                 `json:"object,omitempty"`
	ForeignID string                 `json:"foreign_id,omitempty"`
	Target    string                 `json:"target,omitempty"`
	Time      time.Time              `json:"time,omitempty"`
	Origin    string                 `json:"origin,omitempty"`
	To        []string               `json:"to,omitempty"`
	Score     float64                `json:"score,omitempty"`
	Extra     map[string]interface{} `json:"-"`
}
