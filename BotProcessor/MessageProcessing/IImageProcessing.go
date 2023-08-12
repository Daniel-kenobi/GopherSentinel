package MessageProcessing

type ImageProcessing interface {
	IsImageAppropriated(base64DecodedImage string) (bool, error)
	GetPlatformToken() (string, error)
}
