package healthevent

// EventDescription is AWS Health Event Descripion
type EventDescription struct {
	Language string `json:"language"`
	Latest   string `json:"latestDescription"`
}

// EventDetail is AWS Health Event Detail
type EventDetail struct {
	Service      string             `json:"service"`
	TypeCode     string             `json:"eventTypeCode"`
	TypeCategory string             `json:"eventTypeCategory"`
	StartTime    string             `json:"startTime"`
	EndTime      string             `json:"endTime`
	Description  []EventDescription `json:"eventDescription"`
}

// Event is AWS Health Event
type Event struct {
	Detail EventDetail `json:"detail"`
}
