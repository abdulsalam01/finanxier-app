package channel

func New(channelRepo channelResource) *channelUsecase {
	return &channelUsecase{
		channelRepo: channelRepo,
	}
}
