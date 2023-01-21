package entities

type RouteError struct {
	Error string `json:"error"`
}

const (
	InvalidPayload string = "Invalid Payload"
)
