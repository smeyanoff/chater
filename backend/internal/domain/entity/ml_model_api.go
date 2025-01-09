package entities

type MLModelApi interface {
	SendRequest(input string) (response string, err error)
}
