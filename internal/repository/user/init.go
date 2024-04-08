package user

func New(
	database databaseResource,
	jwtSetupKey string,
) *UsersRepo {
	return &UsersRepo{
		database:    database,
		jwtSetupKey: jwtSetupKey,
	}
}
