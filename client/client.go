package client

type ClientInterface interface {
	MakeCall
	Connect
	MuteAudio
	MuteVideo
	SendChatMessage
	Disconnect
	DisconnectCall
	AddCall
	Renegotiate
	GetPresentation
	StopPresentation
	Present
	GetMediaStatistics
}
