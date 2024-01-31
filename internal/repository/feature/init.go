package feature

func New(database databaseResource) *FeaturesRepo {
	return &FeaturesRepo{
		database: database,
	}
}
