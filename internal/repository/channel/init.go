package channel

func New(database databaseResource) *ChannelsRepo {
	return &ChannelsRepo{
		database: database,
	}
}
