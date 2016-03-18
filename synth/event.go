package synth

var (
	//ErrorType is a constant used to define error events
	ErrorType = "error"

	//OKType is a constant used to define non-error events
	OKType = "ok"
)

// Event describes an event sent from a synthetic transaction
type Event struct {
	ID          string   `json:"id"`
	Application string   `json:"applicationName"`
	Title       string   `json:"title"`
	Type        string   `json:"type"`
	Tags        []string `json:"tags"`
	Metrics     []Metric `json:"metrics"`
	Message     string   `json:"message"`
}

// Metric describes a metric recorded by a synthetic transaction
type Metric struct {
	Type        string      `json:"type"`
	Payload     interface{} `json:"payload"`
	Description string      `json:"description"`
	Tags        []string    `json:"tags"`
}
