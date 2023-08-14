package TextProcessor

type ITextProcessor interface {
	IsMessageAppropriated(message string) bool
}
