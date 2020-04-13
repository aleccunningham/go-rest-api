package client

type ConferenceInterface interface {
	DialOut
	SetConferenceLock
	SetMuteAllGuests
	SetParticipantMute
	SetParticipantSpotlight
	SetParticipantText
	SetRole
	UnlockParticipant
	SendDTMF
	SendFECC
	SetBuzz
	ClearBuzz
	ClearAllBuzz
	TransformLayout
	TransferParticipant
	StartConference
	DisconnectParticipant
	DisconnectAll
}
