package entities

type User struct {
	Id                int64
	FullName          string
	Email             string
	Username          string
	Password          string
	ConfirmedPassword string
}
