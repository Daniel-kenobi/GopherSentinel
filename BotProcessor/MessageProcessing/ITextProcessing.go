package MessageProcessing

type ITextProcessing interface {
	IsMessageAppropriated(message string) bool
}
