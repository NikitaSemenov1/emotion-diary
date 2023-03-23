package requestEntities

type authEntity struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

type RegisterEntity struct {
	authEntity
}

type LoginEntity struct {
	authEntity
}

type LogoutEntity struct {
}
