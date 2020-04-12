package command

type Participant struct {
	UUID                  string  `json:"uuid"`
	URI                   string  `json:"uri"`
	DisplayName           string  `json:"display_name"`
	HasMedia              bool    `json:"has_media"`
	StartTime             uint    `json:"start_time"`
	Vad                   uint    `json:"vad"`
	IsMuted               string  `json:"is_muted"`
	IsPresenting          string  `json:"is_presenting"`
	RxPresentationPolicy  string  `json:"rx_presentation_policy"`
	Role                  string  `json:"role"`
	ServiceType           string  `json:"service_type"`
	Type                  string  `json:"type"`
	PresentationSupported string  `json:"presentation_supported"`
	StageIndex            int     `json:"stage_index"`
	OverlayText           string  `json:"overlay_text"`
	Spotlight             float64 `json:"spotlight"`
}

type ParticipantsResult struct {
	Status  string        `json:"statusx"`
	Objects []Participant `json:"result"`
}

type ConferencesResult struct {
	Objects []struct {
		Id        string `json:"id"`
		Name      string `json:"name"`
		StartTime string `json:"start_time"`
		IsLocked  bool   `json:"is_locked"`
	} `json:"objects"`
}
