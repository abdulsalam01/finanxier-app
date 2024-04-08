package product

func New(database databaseResource) *ProductsRepo {
	return &ProductsRepo{
		database: database,
	}
}
