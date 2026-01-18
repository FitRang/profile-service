package db

type Access struct {
	Username  string `bson:"username"`
	Requester string `bson:"requester"`
}

func ToBsonAccess(username, requester string) Access {
	return Access{
		Username:  username,
		Requester: requester,
	}
}
