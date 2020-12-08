package waylay

type PlugResponse struct {
	ObservedState string `json:"observedState"`
	RawData interface{} `json:"rawData"`
	Message string `json:"message"`
}