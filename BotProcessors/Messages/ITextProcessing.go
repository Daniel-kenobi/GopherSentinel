package Messages

type ITextProcessing interface {
	IsMessageAppropriated(message string) bool
}
