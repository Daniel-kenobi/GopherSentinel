package ImageProcessor

type IImageProcessor interface {
	IsImageAppropriated(base64DecodedImage string) (bool, error)
	GetPlatformToken() string
}
