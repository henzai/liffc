package things

type TrialProducts []TrialProduct

type TrialProduct struct {
	ID                     string
	Name                   string
	Type                   string
	ChannelID              int64
	ActionURI              string
	ServiceUUID            string
	PsdiServiceUUID        string
	PsdiCharacteristicUUID string
}
